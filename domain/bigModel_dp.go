package domain

import (
	"errors"
)

var bigModelTypeList = []string{"taichu-VQA", "taichu-TextToImg",
	"taichu-ImgToText", "luojiaNet", "pangu-QA",
	"codegeex"}

// BigModel
type BigModel interface {
	BigModel() string
}

func NewBigModel(t string) (BigModel, error) {
	if t == "" {
		return nil, errors.New("bigmodel can not be null")
	}

	for j, bt := range bigModelTypeList {
		if t == bt {
			break
		}
		if j == len(bigModelTypeList)-1 {
			return nil, errors.New("can not support this bigmodel type")
		}
	}

	return bigModel(t), nil
}

type bigModel string

func (t bigModel) BigModel() string {
	return string(t)
}
