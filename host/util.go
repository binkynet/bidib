package host

// Assign value to *dst.
// Return if *dst was different from value.
func compareAndAssign[T comparable](dst *T, value T) bool {
	if *dst == value {
		return false
	}
	*dst = value
	return true
}
