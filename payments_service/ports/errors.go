package ports

type DBConnectionError struct {
	Message string
}

func (e *DBConnectionError) Error() string {
	return e.Message
}
