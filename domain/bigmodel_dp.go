package domain

import (
	"errors"
)

var (
	BigModelTypes = [...]string{
		"vqa", "gen_picture",
		"desc_picture", "luojia", "pangu",
		"codegeex", "wukong", "ai_detector",
		"baichuan", "glm2",
	}
)

// BigModel
type BigModel interface {
	BigModel() string
}

func NewBigModel(t string) (BigModel, error) {
	if t == "" {
		return nil, errors.New("bigmodel can not be null")
	}

	if !isBigModelType(t) {
		return nil, errors.New("can not support this bigmodel type")
	}

	return bigModel(t), nil
}

func isBigModelType(t string) bool {
	for i := range BigModelTypes {
		if t == BigModelTypes[i] {
			return true
		}
	}

	return false
}

type bigModel string

func (b bigModel) BigModel() string {
	return string(b)
}
