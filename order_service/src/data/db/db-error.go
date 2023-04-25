package db

import "errors"

type DBConnectionError struct {
	message   string
	FieldName string
}

func (connectionError DBConnectionError) Error() string {
	return connectionError.message
}

func (connectionError DBConnectionError) UnWrap() error {
	return errors.New(connectionError.message)
}
