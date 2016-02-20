package iterators

import (
	"fmt"
)

func toArray(it Iterator) {
	var a []int
	for it.Remaining() {
		v, _ := it.Forward()
		a = append(a, v.(int))
	}
	fmt.Println(a)
}
