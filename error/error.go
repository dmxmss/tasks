package error

type Error int

const (
	ErrDbTransactionFailed Error = iota
	ErrDbInitError
	ErrDbUserForeignKeyViolation
	ErrDbTaskNotFound
	ErrInvalidRequestBody
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
	case ErrDbTaskNotFound:
		err = "Database: task not found"
	case ErrInvalidRequestBody:
		err = "Invalid request body"
	}

	return err
}
