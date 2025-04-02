package error

type Error int

const (
	ErrDbTransactionFailed Error = iota
)

func (e Error) Error() string {
	var err string
	switch e {
	case ErrDbTransactionFailed:
		err = "Database: transaction failed"
	}

	return err
}
