package error

type Error int

const (
	ErrDbTransactionFailed Error = iota
	ErrDbInitError
	ErrDbUserForeignKeyViolation
	ErrDbTaskNotFound
	ErrAuthSignatureInvalid
	ErrAuthTokenExpired
	ErrAuthInvalidCredentials
	ErrAuthTokenInvalid
	ErrAuthFailed
	ErrHashingFailed
	ErrInvalidRequestBody
	ErrUserAlreadyExists
	ErrUserNotFound
	ErrCityNotFound
	ErrUserTasksNotFound
	ErrTokenSigningError
	ErrGetWeatherFailed
	ErrRedisKeyNotFound
	ErrRedisSettingValue
	ErrJSONError
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
	case ErrAuthSignatureInvalid:
		err = "Auth: token signature invalid"
	case ErrAuthTokenExpired:
		err = "Auth: token expired"
	case ErrAuthTokenInvalid:
		err = "Auth: token invalid"
	case ErrAuthFailed:
		err = "Auth: failed"
	case ErrAuthInvalidCredentials:
		err = "Auth: invalid credentials"
	case ErrHashingFailed:
		err = "Hashing failed"
	case ErrInvalidRequestBody:
		err = "Invalid request body"
	case ErrUserAlreadyExists:
		err = "User already exists"
	case ErrUserNotFound:
		err = "User not found"
	case ErrUserTasksNotFound:
		err = "User tasks not found"
	case ErrTokenSigningError:
		err = "Token signing error"
	case ErrGetWeatherFailed:
		err = "Failed getting weather"
	case ErrCityNotFound:
		err = "City not found"
	case ErrRedisKeyNotFound:
		err = "Redis key not found"
	case ErrRedisSettingValue:
		err = "Redis failed setting value"
	case ErrJSONError:
		err = "Error serializing to or deserializing from json value"
	}

	return err
}
