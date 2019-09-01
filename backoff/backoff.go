package backoff

import (
	"time"
)

// backoff implements various backoff mechanisms //
type BackOff interface {
	// increment backoff counter
	Increment()
	// set the backoff timer back to 0
	Reset()
	// blocks until duration of backoff timer
	Do()
}

type expoBackOff struct {
	t           time.Duration
	maxDuration time.Duration
}

// NewExpoBackoff returns a BackOff interface where Increment() raises
// the backoff timer by a power of 2 until maxDuration is hit.
func NewExpoBackoff(maxDuration time.Duration) BackOff {
	return &expoBackOff{
		t:           1 * time.Second,
		maxDuration: maxDuration,
	}
}

func (b *expoBackOff) Increment() {
	if b.t < b.maxDuration {
		b.t = b.t * 2
	}
}
func (b *expoBackOff) Reset() {
	b.t = 1 * time.Second
}
func (b *expoBackOff) Do() {
	if b.t != 1*time.Second {
		time.Sleep(b.t)
	}
}

type staticBackOff struct {
	t time.Duration
}

// NewExpoBackoff returns a BackOff interface where Increment() raises
// the backoff timer by a power of 2 until maxDuration is hit.
func NewStaticBackoff(duration time.Duration) BackOff {
	return &staticBackOff{
		t: duration,
	}
}

func (b *staticBackOff) Increment() {
	return
}
func (b *staticBackOff) Reset() {
	return
}
func (b *staticBackOff) Do() {
	time.Sleep(b.t)
}
