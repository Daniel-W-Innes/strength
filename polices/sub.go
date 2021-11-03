package polices

// SubPolicy Generic sub policy for one specific role of a password policy
type SubPolicy interface {
	Check(s string) error
}
