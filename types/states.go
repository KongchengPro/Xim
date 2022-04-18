package types

type Storage[T any] struct {
	States    *T
	Mutations map[string]func(*T, ...Value)
}

func (s *Storage[T]) Commit(name string, args ...Value) {
	s.Mutations[name](s.States, args...)
}
