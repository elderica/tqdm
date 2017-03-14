package iterators

import (
	"fmt"
)

func ExampleInterval() {
	toArrayInterval(Interval(0, 0))
	toArrayInterval(Interval(3, 10))
	toArrayInterval(Interval(10, 3))
	toArrayInterval(Interval(-3, 3))
	toArrayInterval(Interval(4, -2))
	toArrayInterval(Interval(-5, -10))
	// Output:
	// []
	// [3 4 5 6 7 8 9]
	// [10 9 8 7 6 5 4]
	// [-3 -2 -1 0 1 2]
	// [4 3 2 1 0 -1]
	// [-5 -6 -7 -8 -9]
}

func toArrayInterval(it Iterator) {
	var a []int
	for it.Remaining() {
		v, _ := it.Forward()
		a = append(a, v.(int))
	}
	fmt.Println(a)
}
