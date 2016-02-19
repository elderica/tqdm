package iterators

func ExampleInterval() {
	toArray(Interval(0, 0))
	toArray(Interval(3, 10))
	toArray(Interval(10, 3))
	toArray(Interval(-3, 3))
	toArray(Interval(4, -2))
	toArray(Interval(-5, -10))
	// Output:
	// []
	// [3 4 5 6 7 8 9]
	// [10 9 8 7 6 5 4]
	// [-3 -2 -1 0 1 2]
	// [4 3 2 1 0 -1]
	// [-5 -6 -7 -8 -9]
}
