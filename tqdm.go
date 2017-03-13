package tqdm

import (
	"github.com/sbwhitecap/tqdm/iterators"
	"io"
	"os"
	"time"
)

// Configuration variables. User can set these variables for customization.
var (
	// RedirectTo was set io.Writer to output the progress indicator.
	RedirectTo = os.Stderr

	// If LeaveProgressIndicator is false, tqdm deletes its traces
	// from screen after it has finished iterating over all elements.
	LeaveProgressIndicator = true

	// If less than RerenderingMinimumIntervalOfTime seconds or
	// RerenderingMinimumIntervalsOfIteration iterations have passed
	// since last progress indicator update, it is not updated again.
	RerenderingMinimumIntervalOfTime       = 500 * time.Millisecond
	RerenderingMinimumIntervalsOfIteration = 1
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

		if finished-lastfinished >= RerenderingMinimumIntervalsOfIteration {
			current := time.Now()
			if current.Sub(lastprint) >= RerenderingMinimumIntervalOfTime {
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
