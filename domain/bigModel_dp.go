package domain

import "errors"

const (
	taichuVQA       = "taichu-VQA"
	taichuTextToImg = "taichu-TextToImg"
	taichuImgToText = "taichu-ImgToText"
	luojiaNet       = "luojiaNet"
	pangu           = "pangu-QA"
	codegeex        = "codegeex"
)

// BigModel
type BigModel interface {
	BigModel() string
}

func NewBigModel(t string) (BigModel, error) {
	if t == "" {
		return nil, errors.New("bigmodel can not be null")
	}

	if !(t == taichuVQA ||
		t == taichuTextToImg ||
		t == taichuImgToText ||
		t == luojiaNet ||
		t == pangu ||
		t == codegeex) {
		return nil, errors.New("can not support this bigmodel type")
	}

	return bigModel(t), nil
}

type bigModel string

func (t bigModel) BigModel() string {
	return string(t)
}
