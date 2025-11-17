package main

import (
	"sync"
	"time"
)

type attemptLimiter struct {
	mu       sync.Mutex
	window   time.Duration
	max      int
	attempts []time.Time
}

func newAttemptLimiter(max int, window time.Duration) *attemptLimiter {
	return &attemptLimiter{
		window: window,
		max:    max,
	}
}

// Allow records an attempt when within limits.
// Returns false and remaining wait time when throttled.
func (l *attemptLimiter) Allow() (bool, time.Duration) {
	l.mu.Lock()
	defer l.mu.Unlock()

	now := time.Now()
	cutoff := now.Add(-l.window)

	// Drop stale attempts outside the window.
	valid := l.attempts[:0]
	for _, ts := range l.attempts {
		if ts.After(cutoff) {
			valid = append(valid, ts)
		}
	}
	l.attempts = valid

	if len(l.attempts) >= l.max {
		wait := l.attempts[0].Add(l.window).Sub(now)
		if wait < 0 {
			wait = 0
		}
		return false, wait
	}

	l.attempts = append(l.attempts, now)
	return true, 0
}
