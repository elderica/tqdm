package iterators

// channel represents channel interval.
type channel struct {
	count     int
	last      int
	ch        chan interface{}
	nextValue interface{}
}

// Channel represents an channel interval [].
func Channel(ch chan interface{}, last int) Iterator {
	return &channel{0, last, ch, nil}
}

func (c *channel) Plan() int {
	if c.last > 0 {
		return c.last
	}

	return -1
}

func (c *channel) Remaining() bool {
	if c.count >= c.last {
		return false
	}

	v, ok := <-c.ch
	if !ok {
		return false
	}

	c.nextValue = v

	return true
}

func (c *channel) Forward() (interface{}, error) {
	var v interface{}
	if c.nextValue != nil {
		v, c.nextValue = c.nextValue, nil

		return v, nil
	}

	if c.count >= c.last {
		return nil, ErrStopIteration
	}

	v, ok := <-c.ch
	if !ok {
		return nil, ErrStopIteration
	}

	c.count++

	return v, nil
}
