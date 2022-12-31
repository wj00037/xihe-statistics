package repository

type FileUploadUsers struct {
	Users []string
}

type FileUploadRecord interface {
	Get() (FileUploadUsers, error)
}
