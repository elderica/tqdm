package main

import (
	"github.com/sbwhitecap/tqdm"
	. "github.com/sbwhitecap/tqdm/iterators"
	"time"
)

func main() {
	m := map[string]int{
		"japan":   128000000,
		"china":   1375000000,
		"germany": 81198999,
	}

	keys := make([]string, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}

	tqdm.With(Strings(keys), "iterating map", func(v interface{}) (brk bool) {
		time.Sleep(500 * time.Millisecond)
		k := v.(string)
		_ = m[k]
		return
	})
}
