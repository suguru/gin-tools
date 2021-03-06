package validator

// MaxLength ... check max length
type MaxLength struct {
	Max int
}

// Validate ... validate param of max length
func (m MaxLength) Validate(param string) bool {
	return !(len(param) > m.Max)
}
