package polices

import (
	"bufio"
	"errors"
	"log"
	"os"
)

// ProhibitedPasswordsPolicy Predefined prohibited passwords
type ProhibitedPasswordsPolicy struct {
	ProhibitedPasswords map[string]struct{} //The bool does not matter, it is just a cheap way of doing hash-based checking
}

func (policy *ProhibitedPasswordsPolicy) Check(s string) error {
	//Check if the password is prohibited
	if _, prohibited := (*policy).ProhibitedPasswords[s]; prohibited {
		return errors.New("this password is prohibited")
	} else {
		return nil
	}
}

// GetProhibitedPasswords Load prohibited password file into a sub policy
func GetProhibitedPasswords(prohibitedPasswordsFile string) (*ProhibitedPasswordsPolicy, error) {
	prohibitedPolicy := ProhibitedPasswordsPolicy{ProhibitedPasswords: make(map[string]struct{})}
	//Open prohibited password file
	file, err := os.Open(prohibitedPasswordsFile)
	if err != nil {
		return nil, err
	}
	//Set file to close after loading is done
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Panic()
		}
	}(file)
	//Load prohibited password from lines in the files
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		prohibitedPolicy.ProhibitedPasswords[scanner.Text()] = struct{}{}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return &prohibitedPolicy, nil
}
