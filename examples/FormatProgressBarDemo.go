package main

import . "github.com/sbwhitecap/tqdm/internal"
import "time"
import "os"

func main() {
	planned := uint(420)
	finished := uint(0)
	start := time.Now()
	render := MakeRendererFunc(os.Stdout)

	for ; finished <= planned; finished++ {
		time.Sleep(1 * time.Second)
		elapsed := time.Since(start)
		render(FormatProgressBar(planned, finished, elapsed))
	}
}
