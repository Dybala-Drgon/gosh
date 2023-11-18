package object

import "fmt"

func show(args ...Object) {
	for _, arg := range args {
		if str, ok := arg.(*String); ok {
			fmt.Println(str.Value)
		} else {
			fmt.Println(arg.String())
		}
	}
}
