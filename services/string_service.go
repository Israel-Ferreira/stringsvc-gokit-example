package services

import (
	"errors"
	"strings"
)

type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

type StrService struct {
}

func (str StrService) Uppercase(s string) (string, error) {
	if s == "" {
		return "", errors.New("empty string")
	}

	return strings.ToUpper(s), nil
}

func (str StrService) Count(s string) int {
	return len(s)
}
