/*
Package tqdm allows you to add a progress indicator to your program with minimal effort.

It's golang reimplementation base on noamraph's tqdm library(https://github.com/noamraph/tqdm).

Make sure that import as shown below.

	import (
		"github.com/sbwhitecap/tqdm"
		. "github.com/sbwhitecap/tqdm/iterators"
	)

Here is a simple example.

	tqdm.With(Interval(0, 10), "", func(v interface{}) (brk bool) {
		time.Sleep(1000 * time.Millisecond)
		return
	})

This is equivalent to below.

	tqdm.R(0, 10, func(v interface{}) (brk bool) {
	       time.Sleep(1000 * time.Millisecond)
	       return
	})

As the result, you will see output like this.

	|#######---| 7/10  70% [elapsed: 00:07 left: 00:03,  1.00 iters/sec]
*/
package tqdm
