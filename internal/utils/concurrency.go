package utils

import (
	"context"

	"golang.org/x/sync/semaphore"
)

type Semaphore struct {
	sem *semaphore.Weighted
}

func NewSemaphore(maxConcurrent int64) *Semaphore {
	return &Semaphore{
		sem: semaphore.NewWeighted(maxConcurrent),
	}
}

func (s *Semaphore) Acquire(ctx context.Context) error {
	return s.sem.Acquire(ctx, 1)
}

func (s *Semaphore) Release() {
	s.sem.Release(1)
}
