package domain

// BigModel Type
type BigModelType interface {
	BigModelType() string
}

func NewBigModelType(t string) (BigModelType, error) {
	return bigModelType(t), nil
}

type bigModelType string

func (t bigModelType) BigModelType() string {
	return string(t)
}
