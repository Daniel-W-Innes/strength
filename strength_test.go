package strength

import (
	"testing"
)

//getPasswordPolicy Generate password policy from information provided in assignment.
//This uses the strength module defined in the same repo.
func getPasswordPolicy() *PasswordPolicy {
	//Load prohibited passwords from prohibitive password file
	prohibitedPasswords, err := GetProhibitedPasswords("testing_files/prohibited_passwords")
	if err != nil {
		return nil
	}
	//Load prohibited regex from prohibited regex file
	prohibitedRegexes, err := GetProhibitedRegexes("testing_files/prohibited_regexes")
	if err != nil {
		return nil
	}
	return &PasswordPolicy{SubPolices: []SubPolicy{
		&LengthPolicy{MinLength: 8, MaxLength: 12},
		prohibitedPasswords,
		prohibitedRegexes,
	}, SubPolicesConcurrent: []SubPolicyConcurrent{
		&CasePolicy{MinLower: 1, MinUpper: 1},
		&NumberPolicy{MinNumbers: 1},
		&SpecialPolicy{MinSpecial: 1, SpecialCharacters: map[rune]bool{'!': true, '@': true, '#': true, '$': true, '%': true, '?': true, '∗': true}},
	}}
}

func Test1(t *testing.T) {
	passwordPolicy := getPasswordPolicy()
	if passwordPolicy == nil {
		t.Fail()
	}
	err := (*passwordPolicy).Check("Password1!")
	if err != nil {
		t.Fail()
	}
	err = (*passwordPolicy).Check("password")
	if err == nil {
		t.Fail()
	}
}
