package polices

import "testing"

func TestProhibitedPasswordsPolicy_Check(t *testing.T) {
	prohibitedPasswordsPolicy := ProhibitedPasswordsPolicy{ProhibitedPasswords: map[string]struct{}{"Password": {}, "pass": {}}}

	err := prohibitedPasswordsPolicy.Check("password")
	if err != nil {
		t.Fail()
	}

	err = prohibitedPasswordsPolicy.Check("Name")
	if err != nil {
		t.Fail()
	}

	err = prohibitedPasswordsPolicy.Check("Password")
	if err == nil {
		t.Fail()
	}

	err = prohibitedPasswordsPolicy.Check("pass")
	if err == nil {
		t.Fail()
	}
}

func BenchmarkProhibitedPasswordsPolicy_Check(b *testing.B) {
	prohibitedPasswordsPolicy := ProhibitedPasswordsPolicy{ProhibitedPasswords: map[string]struct{}{"Password": {}, "pass": {}}}

	for i := 0; i < b.N; i++ {
		_ = prohibitedPasswordsPolicy.Check("password")
	}
}
