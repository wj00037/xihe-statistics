package repositories

import (
	"project/xihe-statistics/domain"
	"project/xihe-statistics/domain/repository"
)

type MediaMapper interface {
	Add(string, int64) error
	Get(string) (MediaDataDO, error)
}

func NewMediaRepository(mapper MediaMapper) repository.Media {
	return &media{mapper}
}

type media struct {
	mapper MediaMapper
}

func (impl *media) Add(m *domain.Media) error {
	return impl.mapper.Add(m.Name.MediaName(), m.CreateAt)
}

func (impl *media) GetAll() (data repository.AllMediaData, err error) {
	var total int64
	d := make([]repository.MediaData, len(domain.Medias))

	for i := range domain.Medias {
		do, err := impl.mapper.Get(domain.Medias[i])
		if err != nil {
			return data, err
		}

		total += do.Counts

		name, err := domain.NewMeidaName(domain.Medias[i])
		if err != nil {
			return data, err
		}

		d[i] = repository.MediaData{
			Name:   name,
			Counts: do.Counts,
		}
	}

	data.Data = d
	data.Total = total

	return
}

type MediaDataDO struct {
	Counts int64
}
