package domain

type UserWithBigModel struct {
	UserName Account `json:"username"`
	BigModel BigModel
	CreateAt int64 `json:"create_at"`
}
