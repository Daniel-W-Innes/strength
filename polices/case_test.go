package polices

import "testing"

func TestCasePolicy_Check(t *testing.T) {
	casePolicy := CasePolicy{
		MinUpper: 1,
		MinLower: 1,
	}

	err := casePolicy.Check("Pp")
	if err != nil {
		t.Fail()
	}

	err = casePolicy.Check("PPpp")
	if err != nil {
		t.Fail()
	}

	err = casePolicy.Check("pp")
	if err == nil {
		t.Fail()
	}

	err = casePolicy.Check("PP")
	if err == nil {
		t.Fail()
	}
}

func BenchmarkCasePolicy_Check(b *testing.B) {
	casePolicy := CasePolicy{
		MinUpper: 1,
		MinLower: 1,
	}

	for i := 0; i < b.N; i++ {
		_ = casePolicy.Check("Pp")
	}
}
