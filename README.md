# [tqdm](https://github.com/sbwhitecap/tqdm)

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

Default output is sent `os.Stderr`. Here is screenshot:

![|#######---| 7/10  70% [elapsed: 7.098091s left: 3s,  0.99 iters/sec]](http://i.imgur.com/lfXJ9uE.gifv)

## Usage
See godoc.
[github.com/sbwhitecap/tqdm](https://godoc.org/github.com/sbwhitecap/tqdm )
[github.com/sbwhitecap/tqdm/iterators](https://godoc.org/github.com/sbwhitecap/tqdm/iterators)

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
