package callback

import "syscall/js"

type Value = js.Value
type Func func(this Value, args ...Value)
type Map map[string]Func
