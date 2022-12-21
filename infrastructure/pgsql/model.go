package pgsql

type UserWithRepo struct {
	UserName string `json:"username"`
	RepoName string `json:"repo_name"`
	CreateAt int64  `json:"create_at"`
}

type BigModelRecord struct {
	UserName string `json:"username"`
	BigModel string `json:"bigmodel"`
	CreateAt int64  `json:"create_at"`
}

func (BigModelRecord) TableName() string {
	return "bigmodel_recode"
}
