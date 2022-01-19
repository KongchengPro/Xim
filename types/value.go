package types

type StringValue struct {
	v string
}

func (s StringValue) Get() string {
	return s.v
}

func (s StringValue) Set(nv string) {
	s.v = nv
}
