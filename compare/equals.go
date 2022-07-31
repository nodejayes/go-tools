package compare

// Equals check if the given item1 == item2
func Equals[T any](item1, item2 T) bool {
	return interface{}(item1) == interface{}(item2)
}

// FilterEquals get a Filter that Equals the given Value
func FilterEquals[T any](item T) func(T, int, []T) bool {
	return func(t T, i int, ts []T) bool {
		return Equals(t, item)
	}
}

// FilterNotEquals get a Filter that not Equals the given Value
func FilterNotEquals[T any](item T) func(T, int, []T) bool {
	return func(t T, i int, ts []T) bool {
		return !Equals(t, item)
	}
}