package types

type Data[T any] interface {
	Calculate() T
}

type StaticData[T any] struct {
	value T
}

func NewStaticData[T any](value T) *StaticData[T] {
	return &StaticData[T]{value: value}
}

func (c *StaticData[T]) Calculate() T {
	return c.value
}

type DynamicData[T any] struct {
	block func() T
}

func NewDynamicData[T any](block func() T) *DynamicData[T] {
	return &DynamicData[T]{block: block}
}

func (d DynamicData[T]) Calculate() T {
	return d.block()
}
