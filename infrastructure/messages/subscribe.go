package messages

import "github.com/sirupsen/logrus"

func (d *SyncRepo) handle(event mq.Event) error {
	msg := event.Message()

	if err := d.validateMessage(msg); err != nil {
		return err
	}

	task, ok, err := d.generator.genTask(msg.Body, msg.Header)
	if err != nil || !ok {
		return err
	}

	d.messageChan <- message{
		msg:  msg,
		task: task,
	}

	return nil
}


func (d *SyncRepo) doTask(log *logrus.Entry)  {
	f := func(msg message) (err error) {
		task := &msg.task
		if err = d.syncservice.SyncRepo(task); err == nil {
			return nil
		}

		s := fmt.Sprintf(
			"%s/%s/%s", task.Owner.Account(), task.RepoName, task.RepoId,
		)
		log.Errorf("sync repo(%s) failed, err:%s", s, err.Error())

		if err = d.sendBack(msg.msg); err != nil {
			log.Errorf(
				"send back the message for repo(%s) failed, err:%s",
				s, err.Error(),
			)
		}

		return nil
	}

	for {
		msg, ok := <-d.messageChan
		if !ok {
			return
		}

		if err := f(msg); err != nil {
			log.Errorf("do task failed, err:%s", err.Error())
		}
	}
}