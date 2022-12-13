package domain

type UserWithBigModel struct { // TODO: important! UseWithBigModel can not be domain model, domain model must essentail entity, and combine interface in BigModelRcordService in app layer
	UserName string `json:"username"`
	BigModel BigModel
	CreateAt int64 `json:"create_at"`
}
