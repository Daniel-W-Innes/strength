package polices

import "testing"

func TestSpecialPolicy_Check(t *testing.T) {
	specialPolicy := SpecialPolicy{
		MinSpecial:        1,
		SpecialCharacters: map[rune]struct{}{'$': {}, ':': {}},
	}

	err := specialPolicy.Check("pass$")
	if err != nil {
		t.Fail()
	}

	err = specialPolicy.Check("pass:")
	if err != nil {
		t.Fail()
	}

	err = specialPolicy.Check("pass$:")
	if err != nil {
		t.Fail()
	}

	err = specialPolicy.Check("pass")
	if err == nil {
		t.Fail()
	}

	err = specialPolicy.Check("pass!")
	if err == nil {
		t.Fail()
	}
}

func BenchmarkSpecialPolicy_Check(b *testing.B) {
	specialPolicy := SpecialPolicy{
		MinSpecial:        1,
		SpecialCharacters: map[rune]struct{}{'$': {}, ':': {}},
	}

	for i := 0; i < b.N; i++ {
		_ = specialPolicy.Check("pass$")
	}
}
