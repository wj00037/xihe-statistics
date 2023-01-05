package repository

import "project/xihe-statistics/domain"

type DownloadRecord interface {
	Add(*domain.DownloadRecord) error
	Get() (int64, error)
}
