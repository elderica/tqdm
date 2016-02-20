package iterators

// rng represents interval.
type rng struct {
	first int
	last  int
	count int
}

// Interval represents an interval [first, last).
func Interval(first, last int) Iterator {
	return &rng{first, last, 0}
}

func (r *rng) Plan() int {
	if r.first == r.last {
		return 0
	}
	if r.first < r.last {
		return (r.last - r.first)
	} else {
		return (r.first - r.last)
	}
}

func (r *rng) Remaining() bool {
	if r.first == r.last {
		return false
	}
	if r.first < r.last {
		return (r.first + r.count) < r.last
	} else {
		return r.last < (r.first - r.count)
	}
}

func (r *rng) Forward() (interface{}, error) {
	if r.first == r.last {
		return 0, ErrStopIteration
	}
	c := r.count
	r.count++
	if r.first < r.last {
		return (r.first + c), nil
	} else {
		return (r.first - c), nil
	}
}
