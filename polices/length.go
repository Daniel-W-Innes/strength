package polices

import "errors"

// LengthPolicy Minimum and maximum password length
type LengthPolicy struct {
	MaxLength int //This is not recommended but required by the assignment
	MinLength int
}

func (policy *LengthPolicy) Check(s string) error {
	//Check if password is too short
	if len(s) < (*policy).MinLength {
		return errors.New("password is too short")
		//Check if password is too long
	} else if len(s) > (*policy).MaxLength {
		return errors.New("password is too long")
	} else {
		return nil
	}
}
