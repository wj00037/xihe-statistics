package app

import (
	"project/xihe-statistics/domain"
	"project/xihe-statistics/domain/repository"
)

func (cmd *UserWithBigModelAddCmd) Validate() error {
	_, err := domain.NewAccount(cmd.UserName.Account())
	if err != nil {
		return err
	}

	_, err = domain.NewBigModel(cmd.BigModel.BigModel())

	if err != nil {
		return err
	}

	return nil
}

type BigModelRecordService interface {
	AddUserWithBigModel(*UserWithBigModelAddCmd) error
	GetBigModelRecordsByType(domain.BigModel) (BigModelDTO, error)
	GetCountsByTypeAndTimeDiff(BigModelCountIncreaseCmd) (BigModelCountIncreaseDTO, error)
	GetBigModelRecordAll() (BigModelAllDTO, error)
}

func NewBigModelRecordService(
	ub repository.UserWithBigModel,
) BigModelRecordService {
	return bigModelRecordService{
		ub: ub,
	}
}

type bigModelRecordService struct {
	ub repository.UserWithBigModel
}

func (b bigModelRecordService) AddUserWithBigModel(cmd *UserWithBigModelAddCmd) error {
	v := new(domain.UserWithBigModel)
	cmd.toBigModel(v)

	err := b.ub.Add(v)
	if err != nil {
		return err
	}
	return nil
}

func (b bigModelRecordService) GetBigModelRecordsByType(d domain.BigModel) (dto BigModelDTO, err error) {
	bm, err := b.ub.Get(d)
	if err != nil {
		return
	}

	calls := len(bm)

	users := make([]string, calls)
	for j := range bm {
		users[j] = bm[j].UserName.Account()
	}
	users = RemoveRepeatedElement(users)

	dto = BigModelDTO{
		BigModel: d.BigModel(),
		Users:    users,
		Calls:    int64(calls),
		Counts:   len(users),
		UpdateAt: getLocalTime(),
	}

	return
}

func (s bigModelRecordService) GetCountsByTypeAndTimeDiff(
	cmd BigModelCountIncreaseCmd,
) (
	dto BigModelCountIncreaseDTO,
	err error,
) {
	bigModel := cmd.BigModel
	startTime, err := toUnixTime(cmd.StartTime)
	if err != nil {
		return
	}
	endTime, err := toUnixTime(cmd.EndTime)
	if err != nil {
		return
	}

	startCount, err := s.ub.GetByTypeAndTime(bigModel, startTime)
	if err != nil {
		return
	}

	endCount, err := s.ub.GetByTypeAndTime(bigModel, endTime)
	if err != nil {
		return
	}

	dto = BigModelCountIncreaseDTO{
		BigModel: bigModel.BigModel(),
		Counts:   endCount - startCount,
	}

	return

}

func (b bigModelRecordService) GetBigModelRecordAll() (dto BigModelAllDTO, err error) {
	var (
		bigmodel domain.BigModel
		calls    int
		counts   = 0
		usersAll []string
	)

	for _, bigmodelType := range domain.GetBigModelTypeList() {

		bigmodel, err = domain.NewBigModel(bigmodelType)
		if err != nil {
			return
		}

		bm, err := b.ub.Get(bigmodel)
		if err != nil {
			return dto, err
		}

		calls += len(bm)

		users := make([]string, len(bm))
		for j := range bm {
			users[j] = bm[j].UserName.Account()
		}
		users = RemoveRepeatedElement(users) // TODO: maybe there is way to optimize

		counts += len(users)
		usersAll = append(usersAll, users...)
	}

	usersAll = RemoveRepeatedElement(usersAll)

	dto = BigModelAllDTO{
		Users:             usersAll,
		Calls:             int64(calls),
		Counts:            counts,
		DedupliacteCounts: len(usersAll),
		UpdateAt:          getLocalTime(),
	}

	return
}

func (cmd *UserWithBigModelAddCmd) toBigModel(r *domain.UserWithBigModel) {
	var createAt int64

	if createAt = cmd.CreateAt; cmd.CreateAt == 0 {
		createAt = getUnixLocalTime()
	}

	*r = domain.UserWithBigModel{
		UserName: cmd.UserName,
		BigModel: cmd.BigModel,
		CreateAt: createAt,
	}
}

func RemoveRepeatedElement(arr []string) (newArr []string) {
	newArr = make([]string, 0)
	for i := 0; i < len(arr); i++ {
		repeat := false
		for j := i + 1; j < len(arr); j++ {
			if arr[i] == arr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, arr[i])
		}
	}
	return
}

type BigModelDTO struct {
	BigModel string   `json:"bigmodel"`
	Users    []string `json:"users"`
	Calls    int64    `json:"counts"`      // calls
	Counts   int      `json:"user_counts"` // xihe user counts
	UpdateAt string   `json:"update_at"`
}

type BigModelAllDTO struct {
	Users             []string `json:"users"`
	Counts            int      `json:"user_counts"`        // xihe user counts
	Calls             int64    `json:"counts"`             // calls
	DedupliacteCounts int      `json:"deduplicate_counts"` // xihe user deduplicate_counts
	UpdateAt          string   `json:"update_at"`
}

type BigModelCountIncreaseDTO struct {
	BigModel string `json:"bigmodel"`
	Counts   int64  `json:"counts"`
}

type UserWithBigModelAddCmd struct {
	domain.UserWithBigModel
}

type BigModelCountIncreaseCmd struct {
	BigModel  domain.BigModel `json:"bigmodel"`
	StartTime string          `json:"start_time"`
	EndTime   string          `json:"end_time"`
}
