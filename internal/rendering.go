package internal

import (
	"fmt"
	"io"
	"math"
	"os"
	"strings"
	"time"
)

// MakeRendererFunc returns rendering function binded to specific io.Writer.
func MakeRendererFunc(w io.Writer) func(string) error {
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
	BlockFinished = "#"
	BlockLeft     = "-"
	SaucerSize    = 10
	LParen        = "|"
	RParen        = "|"
)

/*
FormatProgressBar generates representation of progress indicator with progress bar shows
statistics of iterations.

Input

plan: planned number of iterations.

finished: number of finished iterations.

elapesd: time elapsed from start to now.

Output

representation of progress indicator with progress bar.
*/
func FormatProgressBar(plan uint, finished uint, elapsed time.Duration) string {
	left := plan - finished

	saucer := make([]byte, 0, len(LParen)+SaucerSize+len(RParen))
	saucer = append(saucer, LParen...)
	bflen := int(float64(finished) / float64(plan) * float64(SaucerSize))
	saucer = append(saucer, strings.Repeat(BlockFinished, bflen)...)
	bllen := SaucerSize - bflen
	saucer = append(saucer, strings.Repeat(BlockLeft, bllen)...)
	saucer = append(saucer, RParen...)

	percentage := fmt.Sprintf("%3d%%", int(float64(finished)/float64(plan)*100.0))
	remaining := time.Duration((elapsed.Seconds()/float64(finished))*float64(left)) * time.Second
	rate := fmt.Sprintf("%5.2f", float64(finished)/elapsed.Seconds())

	return fmt.Sprintf("%s %d/%d %s [elapsed: %s left: %s, %s iters/sec]",
		string(saucer), finished, plan, percentage, elapsed, remaining, rate)
}

/*
FormatSpeedMeter generates representation of progress indicator shows speed of iterations.

Input

plan: ignored

finished: number of finished iterations.

elapesd: time elapsed from start to now.

Output

representation of progress indicator shows speed of iterations.
*/
func FormatSpeedMeter(plan uint, finished uint, elapsed time.Duration) string {
	rate := fmt.Sprintf("%5.2f", float64(finished)/elapsed.Seconds())

	return fmt.Sprintf("%d [elapsed: %s, %s iters/sec]", finished, elapsed, rate)
}
