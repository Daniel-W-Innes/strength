package strength

import (
	"github.com/Daniel-W-Innes/strength/polices"
	"testing"
)

//getPasswordPolicy Generate password policy from information provided in assignment.
//This uses the strength module defined in the same repo.
func getPasswordPolicy() *PasswordPolicy {
	//Load prohibited passwords from prohibitive password file
	prohibitedPasswords, err := polices.GetProhibitedPasswords("testing_files/prohibited_passwords")
	if err != nil {
		return nil
	}
	//Load prohibited regex from prohibited regex file
	prohibitedRegexes, err := polices.GetProhibitedRegexes("testing_files/prohibited_regexes")
	if err != nil {
		return nil
	}
	return &PasswordPolicy{SubPolices: []polices.SubPolicy{
		&polices.CasePolicy{MinLower: 1, MinUpper: 1},
		&polices.NumberPolicy{MinNumbers: 1},
		&polices.SpecialPolicy{MinSpecial: 1, SpecialCharacters: map[rune]struct{}{'!': {}, '@': {}, '#': {}, '$': {}, '%': {}, '?': {}, '∗': {}}},
		&polices.LengthPolicy{MinLength: 8, MaxLength: 12},
		prohibitedPasswords,
		prohibitedRegexes},
	}
}

func TestPasswordPolicy_Check(t *testing.T) {
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

func BenchmarkPasswordPolicy_Check(b *testing.B) {
	passwordPolicy := getPasswordPolicy()
	if passwordPolicy == nil {
		b.Fail()
	}
	for i := 0; i < b.N; i++ {
		_ = (*passwordPolicy).Check("Password1!")
	}
}
