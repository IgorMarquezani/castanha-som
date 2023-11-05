package user

import "fmt"

type Validator struct {
	Name     string
	Password string
}

func NewValidator(name, passwd string) *Validator {
	return &Validator{
		Name:     name,
		Password: passwd,
	}
}

func (v *Validator) ValidateName() *Error {
	Err := NewError("", make([]string, 0))

	if v.Name == "" {
		Err.Message = ErrEmptyName
	} else if len(v.Name) > 80 {
		Err.Message = ErrTooLongName
		Err.Details = append(Err.Details, fmt.Sprintf("the given name is %d characters long", len(v.Name)))
	}

	if len(Err.Message) > 0 {
		return Err
	}

	return nil
}

func (v *Validator) ValidatePassword() *Error {
	var (
		Especial, Numeric, Upper bool
		Err                      *Error
		Details                  = make([]string, 0, 4)
	)

	for _, c := range v.Password {
		if c >= 'A' && c <= 'Z' {
			Upper = true
		} else if c >= '0' && c <= '9' {
			Numeric = true
		} else if c >= '!' && c <= '/' {
			Especial = true
		} else if c >= ':' && c <= '@' {
			Especial = true
		} else if c >= '[' && c <= '`' {
			Especial = true
		} else if c >= '{' && c <= '~' {
			Especial = true
		}
	}

	if !Especial {
		Details = append(Details, "missing a especial character")
	}
	if !Numeric {
		Details = append(Details, "missing a numeric character")
	}
	if !Upper {
		Details = append(Details, "missing a upper case character")
	}
	if len(v.Password) < 8 {
		Details = append(Details, "the password must have at least 8 characters")
	}
	if len(Details) > 0 {
		Err = NewError(ErrInvalidPasswd, Details)
	}

	return Err
}
