package polices

import (
	"errors"
	"unicode"
)

// CasePolicy Minimum number of uppercase and lowercase letters
type CasePolicy struct {
	MinUpper int
	MinLower int
}

func (policy *CasePolicy) Check(s string) error {
	numUpper := 0
	numLower := 0
	//Loop over spring counting uppercase and lowercase letters
	for _, character := range s {
		if unicode.IsLetter(character) {
			if unicode.IsUpper(character) {
				numUpper++
			} else {
				numLower++
			}
		}
		//Return once enough letters of each case have been found
		if numLower >= (*policy).MinLower && numUpper >= (*policy).MinUpper {
			return nil
		}
	}
	//Generate error message
	if numLower < (*policy).MinLower && numUpper < (*policy).MinUpper {
		return errors.New("not enough lowercase and uppercase letters")
	} else if numLower < (*policy).MinLower {
		return errors.New("not enough lowercase letters")
	} else {
		return errors.New("not enough uppercase letters")
	}
}
