package main

import (
	"github.com/sbwhitecap/tqdm"
	. "github.com/sbwhitecap/tqdm/iterators"
	"time"
)

func main() {
	tqdm.With(Ints([]int{1, 2, 3}), "ints", func(v interface{}) bool {
		time.Sleep(1 * time.Second)
		return false
	})
}
