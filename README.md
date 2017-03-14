# [tqdm](https://github.com/sbwhitecap/tqdm) [![GoDoc](https://godoc.org/github.com/sbwhitecap/tqdm?status.svg)](https://godoc.org/github.com/sbwhitecap/tqdm)

Attach a progress indicator quickly on your program.

This is golang reimplementation base on interesting library [tqdm](https://github.com/noamraph/tqdm).

## Example
```golang
package main

import (
	"github.com/sbwhitecap/tqdm"
	. "github.com/sbwhitecap/tqdm/iterators"
	"time"
)

func main() {
	tqdm.With(Interval(0, 10), "hello", func(v interface{}) (brk bool) {
		time.Sleep(1000 * time.Millisecond)
		return
	})
}
```

The default output is sent to `os.Stderr`. Here is screenshot:

![|#######---| 7/10  70% [elapsed: 00:07 left: 00:03,  1.00 iters/sec]](http://i.imgur.com/1a21isA.gif)

## Usage
See [tqdm's godoc](https://godoc.org/github.com/sbwhitecap/tqdm) and [tqdm/iterators's godoc](https://godoc.org/github.com/sbwhitecap/tqdm/iterators).

## Install
```bash
go get -u github.com/sbwhitecap/tqdm
```

## Contributing
Pull requests are welcome.

 * Revising documentation
 * Adding new feature
 * Sending patch for bug fix
 * Suggest for improvement

## License
* This library is licensed under [ISC license](https://opensource.org/licenses/ISC).
