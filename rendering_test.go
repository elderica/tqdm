package tqdm

import (
	"os"
	"testing"
	"time"
)

func ExampleMakeRenderer() {
	renderer := makeRenderer(os.Stdout)
	renderer("Hello, World!")
	// Output:
	// Hello, World!
}

type progressbar struct {
	plan     uint
	finished uint
	elapsed  time.Duration
}

var formatindicatortests = []struct {
	plan     uint
	finished uint
	elapsed  time.Duration
}{
	{10, 2, 2 * time.Second},
	{350, 221, 4 * time.Minute},
}

var formatprogressbartests = []string{
	"|##--------| 2/10  20% [elapsed: 2s left: 8s,  1.00 iters/sec]",
	"|######----| 221/350  63% [elapsed: 4m0s left: 2m20s,  0.92 iters/sec]",
}

func TestFormatProgressBar(t *testing.T) {
	for idx, in := range formatindicatortests {
		s := formatProgressBar(in.plan, in.finished, in.elapsed)
		if s != formatprogressbartests[idx] {
			t.Errorf("FormatProgressBar(%d, %d, %s)\ngot  %q,\nwant %q",
				in.plan, in.finished, in.elapsed, s, formatprogressbartests[idx])
		}
	}
}

func BenchmarkFormatProgressBar(b *testing.B) {
	in := formatindicatortests[0]
	for i := 0; i < b.N; i++ {
		_ = formatProgressBar(in.plan, in.finished, in.elapsed)
	}
}

var formatspeedmetertests = []string{
	"2 [elapsed: 2s,  1.00 iters/sec]",
	"221 [elapsed: 4m0s,  0.92 iters/sec]",
}

func TestFormatSpeedMeter(t *testing.T) {
	for idx, in := range formatindicatortests {
		s := formatSpeedMeter(in.plan, in.finished, in.elapsed)
		if s != formatspeedmetertests[idx] {
			t.Errorf("FormatSpeedmeter(%d, %s)\ngot  %q,\nwant %q",
				in.finished, in.elapsed, s, formatspeedmetertests[idx])
		}
	}
}
