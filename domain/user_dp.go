package domain

import (
	"errors"
	"regexp"
	"strings"
)

var (
	reName = regexp.MustCompile("^[a-zA-Z0-9_-]+$")
)

type Account interface {
	Account() string
}

type dpAccount string

func (r dpAccount) Account() string {
	return string(r)
}

func NewAccount(v string) (Account, error) {
	if v == "" || strings.ToLower(v) == "root" || !reName.MatchString(v) {
		return nil, errors.New("invalid user name")
	}

	return dpAccount(v), nil
}
