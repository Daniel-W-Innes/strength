package polices

import (
	"bufio"
	"errors"
	"log"
	"os"
	"regexp"
)

// ProhibitedRegexesPolicy Predefined prohibited regexes patterns.
// Useful for preventing phone numbers, postal codes, license plates, etc.
type ProhibitedRegexesPolicy struct {
	ProhibitedRegexes []*regexp.Regexp
}

func (policy *ProhibitedRegexesPolicy) Check(s string) error {
	for _, regexes := range (*policy).ProhibitedRegexes {
		if regexes.MatchString(s) {
			return errors.New("prohibited pattern use password")
		}
	}
	return nil
}

// GetProhibitedRegexes Load regex password file into a sub policy
func GetProhibitedRegexes(ProhibitedRegexesFile string) (*ProhibitedRegexesPolicy, error) {
	prohibitedRegexesPolicy := ProhibitedRegexesPolicy{[]*regexp.Regexp{}}
	file, err := os.Open(ProhibitedRegexesFile)
	if err != nil {
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Panic()
		}
	}(file)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		compile, err := regexp.Compile(scanner.Text())
		if err != nil {
			return nil, err
		}
		prohibitedRegexesPolicy.ProhibitedRegexes = append(prohibitedRegexesPolicy.ProhibitedRegexes, compile)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return &prohibitedRegexesPolicy, nil
}
