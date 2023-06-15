package repository

import "project/xihe-statistics/domain"

type MediaData struct {
	Name   domain.MediaName
	Counts int64
}

type AllMediaData struct {
	Total int64
	Data  []MediaData
}

type Media interface {
	Add(*domain.Media) error
	GetAll() (AllMediaData, error)
}
