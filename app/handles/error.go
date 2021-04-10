package handles

import "errors"

var (
	ErrExpDate = errors.New("the end time is greater than start time")
)
