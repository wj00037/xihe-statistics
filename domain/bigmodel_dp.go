package domain

import (
	"errors"
)

func GetBigModelTypeList() []string {
	return []string{"vqa", "gen_picture",
		"desc_picture", "luojia", "pangu",
		"codegeex", "wukong", "ai_detector",
		"baichuan",
	}
}

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
	bigModelTypeList := GetBigModelTypeList()
	for j, bt := range bigModelTypeList {
		if t == bt {
			break
		}
		if j == len(bigModelTypeList)-1 {
			return false
		}
	}

	return true
}

type bigModel string

func (b bigModel) BigModel() string {
	return string(b)
}
