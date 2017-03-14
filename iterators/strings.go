package iterators

// Strings wraps []string in Iterator
func Strings(p []string) Iterator {
	return &stringSlice{p, 0}
}

type stringSlice struct {
	contents []string
	idx      int
}

func (is *stringSlice) Plan() int {
	return len(is.contents)
}

func (is *stringSlice) Remaining() bool {
	return is.idx < len(is.contents)
}

func (is *stringSlice) Forward() (interface{}, error) {
	if len(is.contents) == is.idx {
		return 0, ErrStopIteration
	}
	item := is.contents[is.idx]
	is.idx++
	return item, nil
}
