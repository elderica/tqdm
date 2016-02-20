package iterators

import (
	"errors"
)

// An iterator-like interface.
// A type that satisfies tqdm.iterators.Iterator can be used by routines
// in this package.
type Iterator interface {
	// Plan tells how many iterations will repeat.
	// It returns negative number if inifinite iterations were planned.
	Plan() int
	// Remaining returns true if Iterator can produce item.
	// Otherwise, it returns false.
	Remaining() bool
	// Forward returns the current item, find out next element
	// and tells any error encountered.
	// When iterator.Remaining() == false, iterator.Forward() may
	// returns ErrStopIteration error.
	Forward() (item interface{}, err error)
}

var ErrStopIteration = errors.New("tqdm.iterators.Iterator: stopped iterations")
