package types

type RawComponent struct {
	Id             string `json:"-"`
	LabelName      string
	Attributes     map[string]string
	EventListeners map[string]func() `json:"-"`
	Content        string
	Children       []*RawComponent
}
