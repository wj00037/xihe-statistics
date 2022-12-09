package domain

type UserWithBigModel struct {
	UserName string `json:"username"`
	BigModel BigModel
	CreateAt int64 `json:"create_at"`
}