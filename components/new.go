package components

func New[T any](constructors ...func(*T)) *T {
	t := new(T)
	for _, constructor := range constructors {
		constructor(t)
	}
	return t
}
