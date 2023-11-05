package user

const (
	ErrEmptyName     = "empty name"
	ErrTooLongName   = "name size cannot be greater than 80"
	ErrInvalidPasswd = "invalid password characters"
)

type Error struct {
	Message string
	Details []string
}

func NewError(message string, details []string) *Error {
	return &Error{
		Message: message,
		Details: details,
	}
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) Is(err error) bool {
	return e.Message == err.Error()
}
