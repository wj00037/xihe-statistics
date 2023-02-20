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

// FileUpload
type FileUploadRecord struct {
	UserName   string `gorm:"column:username"`
	UploadPath string `gorm:"column:upload_path"`
	CreateAt   int64  `gorm:"column:create_at"`
}

func (FileUploadRecord) TableName() string {
	return "fileupload_record"
}

// Download
type DownloadRecord struct {
	UserName     string `gorm:"column:username"`
	DownloadPath string `gorm:"column:download_path"`
	CreateAt     int64  `gorm:"column:create_at"`
}

func (DownloadRecord) TableName() string {
	return "download_record"
}

// GitLab
type GitLabRecord struct {
	Counts   int64 `gorm:"column:counts"`
	CreateAt int64 `gorm:"column:create_at"`
}

func (GitLabRecord) TableName() string {
	return "gitlab_record"
}

// Train
type TrainRecord struct {
	UserName  string `gorm:"column:username"`
	ProjectId string `gorm:"column:project_id"`
	TrainId   string `gorm:"column:train_id"`
	CreateAt  int64  `gorm:"column:create_at"`
}

func (TrainRecord) TableName() string {
	return "train_record"
}
