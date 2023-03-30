package repository

import "project/xihe-statistics/domain"

type CloudRecord interface {
	Add(*domain.Cloud) error
}
