package iterators

import (
	"fmt"
)

func toArray(it Iterator) {
	a := make([]int, 0)
	for it.Remaining() {
		v, _ := it.Forward()
		a = append(a, v.(int))
	}
	fmt.Println(a)
}
