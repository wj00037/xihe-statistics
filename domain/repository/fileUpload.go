package repository

import "project/xihe-statistics/domain"

type FileUploadUsers struct {
	Users []string
}

type FileUploadRecord interface {
	Get() (FileUploadUsers, error)
	Add(*domain.FileUploadRecord) error
}
