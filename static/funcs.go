package static

// Check if an element is in a slice of any comparable type
func IsIn[T comparable](val T, all []T) bool {
	for _, v := range all {
		if val == v {
			return true
		}
	}
	return false
}

// Check if paths are diffrent
// NOTE : it doesn't check the first and last elements
func Diffrent[T comparable](a1, a2 *[]T) bool {
	for _, v := range (*a1)[1 : len(*a1)-1] {
		if IsIn(v, *a2) {
			return false
		}
	}
	*a2 = append(*a2, (*a1)[1:len(*a1)-1]...)
	return true
}
