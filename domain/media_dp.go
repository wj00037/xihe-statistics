package domain

import "errors"

const (
	MediaXinHua  = "xinhua"
	MediaSouHu   = "souhu"
	Media51CTO   = "51cto"
	MediaCSDN    = "csdn"
	MediaTianJi  = "tianji"
	MediaZhiDing = "zhiding"
	MediaZOL     = "zol"
)

var Medias = []string{MediaXinHua, MediaSouHu,
	Media51CTO, MediaCSDN,
	MediaTianJi, MediaZhiDing,
	MediaZOL}

// MediaName
type MediaName interface {
	MediaName() string
}

func NewMeidaName(v string) (MediaName, error) {
	b := v == MediaXinHua ||
		v == MediaSouHu ||
		v == Media51CTO ||
		v == MediaCSDN ||
		v == MediaTianJi ||
		v == MediaZhiDing ||
		v == MediaZOL

	if !b {
		return nil, errors.New("invalid media name")
	}

	return mediaName(v), nil
}

type mediaName string

func (r mediaName) MediaName() string {
	return string(r)
}
