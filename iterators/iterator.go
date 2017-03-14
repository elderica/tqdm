package iterators

import (
	"errors"
)

// Iterator is an iterator-like interface.
// A type that satisfies Iterator can be used with tqdm package.
type Iterator interface {
	// Plan tells how many iterations will repeat.
	// It returns negative number if infinite iterations will plan.
	Plan() int
	// Remaining returns true if Iterator can produce item.
	// Otherwise, it returns false.
	Remaining() bool
	// Forward returns the current item, find out next element and
	// tells any error encountered. When Remaining() returns false,
	// Forward may returns ErrStopIteration error.
	Forward() (item interface{}, err error)
}

// ErrStopIteration represents iterator no longer produce items.
var ErrStopIteration = errors.New("iterators.Iterator: stopped iterations")
