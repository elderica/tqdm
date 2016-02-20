package iterators

func ExampleInts() {
	toArray(Ints([]int{1, 2, 3}))
	toArray(Ints([]int{2, 5, 1}))
	toArray(Ints([]int{}))
	toArray(Ints([]int{1}))
	// Output:
	// [1 2 3]
	// [2 5 1]
	// []
	// [1]
}
