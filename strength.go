package strength

import (
	"github.com/Daniel-W-Innes/strength/polices"
)

// PasswordPolicy Class for checking passwords against predefined password policy
type PasswordPolicy struct {
	SubPolices []polices.SubPolicy
}

// Check All for the sub polices on a string
// This clears the temp polices after check
func (policy *PasswordPolicy) Check(s string) error {
	errs := make(chan error)
	numPolices := len((*policy).SubPolices)
	defer func() {
		for numPolices > 0 {
			<-errs
			numPolices--
		}
		close(errs)
	}()
	for _, subPolice := range (*policy).SubPolices {
		go func(errs chan<- error, subPolicy polices.SubPolicy) {
			errs <- subPolicy.Check(s)
		}(errs, subPolice)
	}
	for numPolices > 0 {
		err := <-errs
		numPolices--
		if err != nil {
			return err
		}
	}
	return nil
}
