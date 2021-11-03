package polices

import "errors"

// SpecialPolicy Minimum number of predefined special characters
type SpecialPolicy struct {
	MinSpecial        int
	SpecialCharacters map[rune]struct{} //The bool does not matter, it is just a cheap way of doing hash-based checking
}

func (policy *SpecialPolicy) Check(s string) error {
	numSpecial := 0
	//Loop over spring counting special characters
	for _, character := range s {
		if _, special := (*policy).SpecialCharacters[character]; special {
			numSpecial++
		}
		//Return once enough special characters have been found
		if numSpecial >= (*policy).MinSpecial {
			return nil
		}
	}
	//Generate error message
	return errors.New("not enough special letters")
}
