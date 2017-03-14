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
	w            io.Writer
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
	symbolFinished = "#"
	symbolLeft     = "-"
	saucerSize     = 10
	lParen         = "|"
	rParen         = "|"
)

func formatProgressBar(plan uint, finished uint, elapsed time.Duration) string {
	left := plan - finished
	ratio := float64(finished) / float64(plan)
	nsf := int(ratio * float64(saucerSize))
	saucer := fmt.Sprintf("%s%s%s%s",
		lParen,
		strings.Repeat(symbolFinished, nsf),
		strings.Repeat(symbolLeft, saucerSize-nsf),
		rParen)

	remaining := time.Duration((elapsed.Seconds()/float64(finished))*float64(left)) * time.Second
	rate := float64(finished)/elapsed.Seconds()

	return fmt.Sprintf("%s %d/%d %3d%% [elapsed: %s left: %s, %5.2f iters/sec]",
		string(saucer), finished, plan, int(ratio*100.0),
		formatTime(elapsed), formatTime(remaining), rate)
}

func formatSpeedMeter(plan uint, finished uint, elapsed time.Duration) string {
	rate := float64(finished)/elapsed.Seconds()

	return fmt.Sprintf("%d [elapsed: %s, %5.2f iters/sec]",
		finished,
		formatTime(elapsed), rate)
}

func formatTime(d time.Duration) string {
	s := (d % time.Minute) / time.Second
	m := (d % time.Hour) / time.Minute
	h := d / time.Hour
	if h == 0 {
		return fmt.Sprintf("%02d:%02d", m, s)
	}
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}
