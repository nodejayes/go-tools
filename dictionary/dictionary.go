package dictionary

type (
	Dictionary[T comparable, K any] struct {
		data map[T]K
	}
)

// New create a new Dictionary
func New[T comparable, K any](data map[T]K) Dictionary[T, K] {
	return Dictionary[T, K]{
		data,
	}
}
