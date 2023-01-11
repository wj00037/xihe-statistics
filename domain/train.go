package domain

type TrainRecord struct {
	UserName  Account `json:"username"`
	ProjectId string `json:"project_id"`
	TrainId   string `json:"train_id"`
	CreateAt  int64  `json:"create_at"`
}
