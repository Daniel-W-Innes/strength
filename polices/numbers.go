package polices

import (
	"errors"
	"unicode"
)

// NumberPolicy Minimum number of numbers
type NumberPolicy struct {
	MinNumbers int
}

func (policy *NumberPolicy) Check(s string) error {
	numNumbers := 0
	//Loop over spring counting numbers
	for _, character := range s {
		if unicode.IsNumber(character) {
			numNumbers++
		}
		//Return once enough numbers have been found
		if numNumbers >= (*policy).MinNumbers {
			return nil
		}
	}
	//Generate error message
	return errors.New("not enough numbers")
}
