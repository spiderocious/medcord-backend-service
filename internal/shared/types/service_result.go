package types

import "github.com/spiderocious/medcord-backend/internal/shared/constants"

type ServiceResult[T any] struct {
	Success    bool
	Data       T
	Err        error
	MessageKey constants.MessageKey
}

func Success[T any](data T, key constants.MessageKey) ServiceResult[T] {
	return ServiceResult[T]{Success: true, Data: data, MessageKey: key}
}

func Failure[T any](err error, key constants.MessageKey) ServiceResult[T] {
	return ServiceResult[T]{Success: false, Err: err, MessageKey: key}
}
