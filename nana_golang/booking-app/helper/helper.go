package helper

func ValidateName(name string) bool {
	if len(name) < 2 || len(name) > 32 {
		return false
	}
	return true
}