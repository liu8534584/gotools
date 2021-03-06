package myerr

import (
	"time"
)

type MyError struct {
	Code       int
	ErrMessage string
	Time       time.Time
}

func NewError(code int, errMessage string) *MyError {
	return &MyError{
		Code:       code,
		ErrMessage: errMessage,
		Time:       time.Now(),
	}
}

func (err *MyError) Error() string {
	return err.ErrMessage
}

func Ok() *MyError {
	return &MyError{
		Code:       0,
		ErrMessage: "ok",
		Time:       time.Now(),
	}
}
