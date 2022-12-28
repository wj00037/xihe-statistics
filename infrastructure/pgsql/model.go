package pgsql

// Repo
type UserWithRepo struct {
	UserName string `gorm:"column:username"`
	RepoName string `gorm:"column:repo_name"`
	CreateAt int64  `gorm:"column:create_at"`
}

func (UserWithRepo) TableName() string {
	return "repo_record"
}

// BigModel
type BigModelRecord struct {
	UserName string `gorm:"column:username"`
	BigModel string `gorm:"column:bigmodel"`
	CreateAt int64  `gorm:"column:create_at"`
}

func (BigModelRecord) TableName() string {
	return "bigmodel_record"
}

// Register
type RegisterRecord struct {
	UserName string `gorm:"column:username"`
	CreateAt int64  `gorm:"column:create_at"`
}

func (RegisterRecord) TableName() string {
	return "register_record"
}
