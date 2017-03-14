package iterators

// Ints wraps []int in Iterator
func Ints(p []int) Iterator {
	return &intSlice{p, 0}
}

type intSlice struct {
	contents []int
	idx      int
}

func (is *intSlice) Plan() int {
	return len(is.contents)
}

func (is *intSlice) Remaining() bool {
	return is.idx < len(is.contents)
}

func (is *intSlice) Forward() (interface{}, error) {
	if len(is.contents) == is.idx {
		return 0, ErrStopIteration
	}
	item := is.contents[is.idx]
	is.idx++
	return item, nil
}
