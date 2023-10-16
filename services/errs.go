package services

import "errors"

var (
	ErrZeroAmount = errors.New("purchase amount could not zero")
	ErrRepository = errors.New("error unexpected repository")
)
