package utils

import (
	"time"
)

type Timer struct {
	start time.Time
}

func StartTimer() *Timer {
	return &Timer{start: time.Now()}
}

func (t *Timer) Elapsed() time.Duration {
	return time.Since(t.start)
}

func (t *Timer) FormatElapsed() string {
	return t.Elapsed().String()
}
