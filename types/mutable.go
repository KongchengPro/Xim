package types

type MutableVar struct {
	Data       Value
	Subscribes []Component
}
