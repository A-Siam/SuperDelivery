package db

import "errors"

type DBConnectionError struct {
	Message   string
	FieldName string
}

func (connectionError DBConnectionError) Error() string {
	return connectionError.Message
}

func (connectionError DBConnectionError) UnWrap() error {
	return errors.New(connectionError.Message)
}
