package domain

import "errors"

// BigModel
type BigModel interface {
	BigModel() string
}

func NewBigModel(t string) (BigModel, error) {
	if t == "" {
		return nil, errors.New("bigmodel can not be null")
	}

	return bigModel(t), nil
}

type bigModel string

func (t bigModel) BigModel() string {
	return string(t)
}
