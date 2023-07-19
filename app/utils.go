package app

import (
	"time"

	"github.com/sirupsen/logrus"
)

var (
	timeFormat = "2006-01-02T15:04:05+08:00"
)

func getTimeLocation() *time.Location {
	cstSh, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		logrus.Debugf("get time location error: %s", err.Error())
	}

	return cstSh
}

func getLocalTime() (t string) {
	return time.Now().In(getTimeLocation()).Format(timeFormat)
}

func toStrTime(t int64) string {
	return time.Unix(t, 0).In(getTimeLocation()).Format(timeFormat)
}

func GetUnixLocalTime() (t int64) {
	return time.Now().Unix()
}

func toTimeStamp(t string) (stamp time.Time, err error) {
	stamp, err = time.ParseInLocation(timeFormat, t, getTimeLocation())
	if err != nil {
		return
	}

	return
}

func toUnixTime(t string) (tt int64, err error) {
	stamp, err := toTimeStamp(t)
	if err != nil {
		return
	}

	tt = stamp.Unix()

	return
}
