package iterators

// Float64s wraps []float64 in Iterator
func Float64s(p []float64) Iterator {
	return &float64Slice{p, 0}
}

type float64Slice struct {
	contents []float64
	idx      int
}

func (is *float64Slice) Plan() int {
	return len(is.contents)
}

func (is *float64Slice) Remaining() bool {
	return is.idx < len(is.contents)
}

func (is *float64Slice) Forward() (interface{}, error) {
	if len(is.contents) == is.idx {
		return 0, ErrStopIteration
	}
	item := is.contents[is.idx]
	is.idx++
	return item, nil
}
