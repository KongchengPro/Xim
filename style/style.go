package style

import (
	"fmt"
	"strings"
)

type Style struct {
	FontSize string
}

func GenerateCSS(style Style) string {
	cssMap := make(map[string]string)
	if style.FontSize != "" {
		cssMap["font-size"] = style.FontSize
	}
	sb := strings.Builder{}
	for key, value := range cssMap {
		sb.WriteString(fmt.Sprintf("%s:%s;", key, value))
	}
	return sb.String()
}
