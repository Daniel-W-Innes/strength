package polices

import "testing"

func TestLengthPolicy_Check(t *testing.T) {
	lengthPolicy := LengthPolicy{
		MaxLength: 3,
		MinLength: 1,
	}

	err := lengthPolicy.Check("aa")
	if err != nil {
		t.Fail()
	}

	err = lengthPolicy.Check("")
	if err == nil {
		t.Fail()
	}

	err = lengthPolicy.Check("aaaa")
	if err == nil {
		t.Fail()
	}
}

func BenchmarkLengthPolicy_Check(b *testing.B) {
	lengthPolicy := LengthPolicy{
		MaxLength: 3,
		MinLength: 1,
	}
	for i := 0; i < b.N; i++ {
		_ = lengthPolicy.Check("aa")
	}
}
