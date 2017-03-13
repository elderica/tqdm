package tqdm

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"time"
)

func makeRenderer(w io.Writer) func(string) error {
	lastwrotelen := 0

	return func(s string) error {
		n := int(math.Max(0, float64(lastwrotelen-len(s))))
		spaces := strings.Repeat(" ", n)
		s2 := "\r" + s + spaces

		_, err := io.WriteString(w, s2)

		if f, ok := w.(*os.File); ok {
			f.Sync()
		}

		lastwrotelen = len(s)

		return err
	}
}

var (
	blockFinished = "#"
	blockLeft     = "-"
	saucerSize    = 10
	lParen        = "|"
	rParen        = "|"
)

func formatProgressBar(plan uint, finished uint, elapsed time.Duration) string {
	left := plan - finished

	saucer := make([]byte, 0, len(lParen)+saucerSize+len(rParen))
	saucer = append(saucer, lParen...)
	bflen := int(float64(finished) / float64(plan) * float64(saucerSize))
	saucer = append(saucer, strings.Repeat(blockFinished, bflen)...)
	bllen := saucerSize - bflen
	saucer = append(saucer, strings.Repeat(blockLeft, bllen)...)
	saucer = append(saucer, rParen...)

	percentage := fmt.Sprintf("%3d%%", int(float64(finished)/float64(plan)*100.0))
	remaining := time.Duration((elapsed.Seconds()/float64(finished))*float64(left)) * time.Second
	rate := fmt.Sprintf("%5.2f", float64(finished)/elapsed.Seconds())

	return fmt.Sprintf("%s %d/%d %s [elapsed: %s left: %s, %s iters/sec]",
		string(saucer), finished, plan, percentage, elapsed, remaining, rate)
}

func formatSpeedMeter(plan uint, finished uint, elapsed time.Duration) string {
	rate := fmt.Sprintf("%5.2f", float64(finished)/elapsed.Seconds())

	return fmt.Sprintf("%d [elapsed: %s, %s iters/sec]", finished, elapsed, rate)
}
