package polices

import (
	"regexp"
	"testing"
)

func TestProhibitedRegexesPolicy_Check(t *testing.T) {
	compile, err := regexp.Compile(`^[+]?[(]?[0-9]{3}[)]?[-\s.]?[0-9]{3}[-\s.]?[0-9]{4,6}$`)
	if err != nil {
		t.Fail()
	}
	prohibitedRegexesPolicy := ProhibitedRegexesPolicy{ProhibitedRegexes: []*regexp.Regexp{compile}}

	err = prohibitedRegexesPolicy.Check("password")
	if err != nil {
		t.Fail()
	}

	err = prohibitedRegexesPolicy.Check("123 654 7890")
	if err == nil {
		t.Fail()
	}

	err = prohibitedRegexesPolicy.Check("(123) 654 7890")
	if err == nil {
		t.Fail()
	}
}

func BenchmarkProhibitedRegexesPolicy_Check(b *testing.B) {
	compile, err := regexp.Compile(`^[+]?[(]?[0-9]{3}[)]?[-\s.]?[0-9]{3}[-\s.]?[0-9]{4,6}$`)
	if err != nil {
		b.Fail()
	}
	prohibitedRegexesPolicy := ProhibitedRegexesPolicy{ProhibitedRegexes: []*regexp.Regexp{compile}}

	for i := 0; i < b.N; i++ {
		_ = prohibitedRegexesPolicy.Check("password")
	}
}
