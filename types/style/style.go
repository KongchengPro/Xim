package style

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"reflect"
	"strings"
)

type CSSMap map[string]string

type Style interface{}

func GenerateCSS(style Style) string {
	defer func() {
		if err := recover(); err != nil {
			logrus.Error(err)
		}
	}()
	cssMap := make(CSSMap)
	st := reflect.TypeOf(style).Elem()
	sv := reflect.ValueOf(style).Elem()
	for i := 0; i < st.NumField(); i++ {
		if cssKey, ok := st.Field(i).Tag.Lookup("css"); ok {
			cssValue := sv.Field(i).String()
			cssMap[cssKey] = cssValue
		}
	}
	sb := strings.Builder{}
	for key, value := range cssMap {
		sb.WriteString(fmt.Sprintf("%s:%s;", key, value))
	}
	return sb.String()
}
