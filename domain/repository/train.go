package repository

type TrainRecord interface {
	Get() (int64, error)
}