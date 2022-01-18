package component

import (
	. "gitee.com/kogic/xim/style"
)

type Component interface {
	// Generate 返回对象类型、内容、央视
	Generate() (string, string, Style)
}

type View struct {
	Components []Component
}

type Text struct {
	Style Style
	Text  string
}

func (t *Text) Generate() (string, string, Style) {
	return "p", t.Text, t.Style
}
