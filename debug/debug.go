package debug

import "fmt"

type Debug bool

var Dbg Debug

func (d Debug) Printf(s string, a ...interface{}) {
	if d {
		fmt.Printf(s, a...)
	}
}
