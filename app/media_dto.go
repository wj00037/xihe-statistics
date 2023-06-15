package app

import (
	"project/xihe-statistics/domain"
	"project/xihe-statistics/domain/repository"
)

type AddMediaCmd struct {
	Name     domain.MediaName
	CreateAt int64
}

type MediaDataDTO struct {
	Name   string `json:"name"`
	Counts int64  `json:"counts"`
}

type AllMediaDataDTO struct {
	Total    int64          `json:"total"`
	Data     []MediaDataDTO `json:"data"`
	UpdateAt string         `json:"update_at"`
}

func toMediaDTO(r *repository.MediaData) MediaDataDTO {
	return MediaDataDTO{
		Name:   r.Name.MediaName(),
		Counts: r.Counts,
	}
}

func toAllMediaDataDTO(r *repository.AllMediaData) (dto AllMediaDataDTO) {
	dto.Total = r.Total
	dto.UpdateAt = getLocalTime()

	data := make([]MediaDataDTO, len(r.Data))
	for i := range data {
		data[i] = toMediaDTO(&r.Data[i])
	}

	dto.Data = data

	return
}
