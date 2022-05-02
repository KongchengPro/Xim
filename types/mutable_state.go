package types

var RefreshChannel = make(chan string, 10)

type MutableState[T any] struct {
	value       T
	subscribers []string
}

func checkInSubscribers(s []string, id string) bool {
	for _, v := range s {
		if id == v {
			return true
		}
	}
	return false
}

func (m *MutableState[T]) Get() T {
	return m.value
}

func (m *MutableState[T]) Value(id string) T {
	if !checkInSubscribers(m.subscribers, id) {
		m.subscribers = append(m.subscribers, id)
	}
	return m.value
}

func (m *MutableState[T]) SetValue(data T) {
	m.value = data
	for _, subscriber := range m.subscribers {
		RefreshChannel <- subscriber
	}
}

func MutableStateOf[T any](data T) MutableState[T] {
	return MutableState[T]{value: data}
}
