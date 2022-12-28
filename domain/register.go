package domain

type RegisterRecord struct {
	UserName string `json:"username"`
	CreateAt int64  `json:"create_at"`
}
