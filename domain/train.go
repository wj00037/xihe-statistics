package domain

type TrainRecord struct {
	UserName string `json:"username"`
	TrainId  string `json:"train_id"`
	CreateAt int64  `json:"create_at"`
}
