package types

type Storage[T any] struct {
	States    *T
	Mutations map[string]func(*T, ...any)
}

func (s *Storage[T]) Commit(name string, args ...any) {
	s.Mutations[name](s.States, args...)
}
