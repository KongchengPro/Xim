package types

type Render interface {
	Render() (*RawComponent, bool)
}
