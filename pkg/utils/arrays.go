package utils

// Contains finds e in array s
func Contains(s []interface{}, e interface{}) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
