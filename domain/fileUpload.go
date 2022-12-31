package domain

type FileUploadRecord struct {
	UserName   string `json:"username"`
	UploadPath string `json:"upload_path"`
	CreateAt   int64  `json:"create_at"`
}
