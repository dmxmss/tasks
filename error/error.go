package error

type Error int

const (
	ErrDbTransactionFailed Error = iota
	ErrDbInitError
	ErrInvalidRequestBody
	ErrDbUserForeignKeyViolation
)

func (e Error) Error() string {
	var err string
	switch e {
	case ErrDbTransactionFailed:
		err = "Database: transaction failed"
	case ErrDbInitError:
		err = "Database: init error"
	case ErrDbUserForeignKeyViolation:
		err = "Database: user foreign key violation"
	case ErrInvalidRequestBody:
		err = "Invalid request body"
	}

	return err
}
