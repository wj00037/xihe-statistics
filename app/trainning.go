package app

type TrainingAddCmd struct {
	UserName string
	DateTime string
}

type TrainingService interface {
	Add(cmd *TranningAddCmd) error
	Get(cmd *DateTimeGetCmd) error
}