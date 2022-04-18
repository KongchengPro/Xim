package utils

import (
	"encoding/json"
	"fmt"
)

func PrintlnStruct(message string, v interface{}) {
	bs, _ := json.MarshalIndent(v, "", "  ")
	fmt.Println(message, string(bs))
}
