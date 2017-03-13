package tqdm

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"time"
)

type renderer struct {
	w io.Writer
	lastwrotelen int
}

func (r *renderer) render(s string) error {
		n := int(math.Max(0, float64(r.lastwrotelen-len(s))))
		s2 := fmt.Sprintf("\r%s%s", s, strings.Repeat(" ", n))

		_, err := io.WriteString(r.w, s2)
		if err != nil {
			return err
		}
		if f, ok := r.w.(*os.File); ok {
			f.Sync()
		}

		r.lastwrotelen = len(s)

		return nil
}

func makeRenderer(w io.Writer) func(string) error {
	return (&renderer{w: w}).render
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
