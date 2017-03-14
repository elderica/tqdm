package tqdm

import (
	"github.com/sbwhitecap/tqdm/iterators"
	"io"
	"os"
	"time"
)

// Configuration variables. Users can set these variables for customization.
var (
	// RedirectTo is where to output the progress indicator.
	RedirectTo io.Writer = os.Stderr

	// If LeaveProgressIndicator is false, tqdm deletes its traces
	// from RedirectTo after it has finished iterating over all elements.
	LeaveProgressIndicator bool = true

	// If less than MinimumIntervalOfTime seconds or
	// less than MinimumTimesOfIteration iterations have passed
	// since last progress indicator update, it is not updated again.
	MinimumIntervalOfTime   time.Duration = 500 * time.Millisecond
	MinimumTimesOfIteration uint          = 1
)

// With does iterations, render a progress indicator
// and rerender it every time a element is requested.
//
// 'description' can contain a short string, describing the progress,
// that it added in the beginning of the line.
//
// With calls 'block' callback every iteration. Callback should return false
// except you want to "break" loop.
func With(it iterators.Iterator, description string, block func(v interface{}) (brk bool)) error {
	render := makeRenderer(RedirectTo)

	prefix := ""
	if description != "" {
		prefix = description + ": "
	}

	plan := it.Plan()
	format := formatProgressBar
	if plan < 0 {
		format = formatSpeedMeter
	}

	start := time.Now()
	lastprint := start
	finished := 0
	lastfinished := finished

	b := false
	for it.Remaining() && !b {
		v, err := it.Forward()
		if err != nil {
			return err
		}

		render(prefix +
			format(uint(plan), uint(finished), time.Since(start)))

		b = block(v)
		finished++

		if uint(finished-lastfinished) >= MinimumTimesOfIteration {
			current := time.Now()
			if current.Sub(lastprint) >= MinimumIntervalOfTime {
				render(prefix +
					format(uint(plan), uint(finished), current.Sub(start)))
				lastfinished = finished
				lastprint = current
			}
		}
	}

	if LeaveProgressIndicator {
		if lastfinished < finished {
			render(prefix +
				format(uint(plan), uint(finished), time.Since(start)))
		}
		io.WriteString(RedirectTo, "\n")
	} else {
		render("")
		// Jump over whitespaces.
		io.WriteString(RedirectTo, "\r")
	}

	return nil
}

// R is a shortcut for writing tqdm.With(Interval(first, last), ...)
func R(first, last int, block func(v interface{}) (brk bool)) {
	With(iterators.Interval(first, last), "", block)
}
