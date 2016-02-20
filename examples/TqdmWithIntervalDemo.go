package main

import (
	"github.com/sbwhitecap/tqdm"
	. "github.com/sbwhitecap/tqdm/iterators"
	"time"
)

func main() {
	//tqdm.LeaveProgressIndicator = false
	tqdm.With(Interval(0, 10), "hello", func(v interface{}) (brk bool) {
		time.Sleep(1000 * time.Millisecond)
		return
	})
}
