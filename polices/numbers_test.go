package polices

import "testing"

func TestNumberPolicy_Check(t *testing.T) {
	numberPolicy := NumberPolicy{MinNumbers: 1}

	err := numberPolicy.Check("a1")
	if err != nil {
		t.Fail()
	}

	err = numberPolicy.Check("11")
	if err != nil {
		t.Fail()
	}

	err = numberPolicy.Check("aa")
	if err == nil {
		t.Fail()
	}
}

func BenchmarkNumberPolicy_Check(b *testing.B) {
	numberPolicy := NumberPolicy{MinNumbers: 1}

	for i := 0; i < b.N; i++ {
		_ = numberPolicy.Check("a1")
	}
}
