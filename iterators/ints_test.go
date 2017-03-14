package iterators

import (
	"fmt"
)

func ExampleInts() {
	toArrayInts(Ints([]int{1, 2, 3}))
	toArrayInts(Ints([]int{2, 5, 1}))
	toArrayInts(Ints([]int{}))
	toArrayInts(Ints([]int{1}))
	_, err := Ints([]int{}).Forward()
	fmt.Println(err.Error() == ErrStopIteration.Error())
	// Output:
	// [1 2 3]
	// [2 5 1]
	// []
	// [1]
	// true
}

func toArrayInts(it Iterator) {
	var a []int
	for it.Remaining() {
		v, _ := it.Forward()
		a = append(a, v.(int))
	}
	fmt.Println(a)
}
