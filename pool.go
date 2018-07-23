package timerpool

import (
	"sync"
	"time"
)

// Pool is a timer pool that contains available time.Timer of same time.Duration
type Pool struct {
	sync.Pool
	timeout time.Duration
}

// New builds new timer pool
func New(timeout time.Duration) *Pool {
	return &Pool{timeout: timeout}
}

// Acquire loads free timer from pool or creates new one.
func (p *Pool) Acquire() *time.Timer {
	tv := p.Get()
	if tv == nil {
		return time.NewTimer(p.timeout)
	}

	t := tv.(*time.Timer)
	if t.Reset(p.timeout) {
		// active timer in sync.Pool, remove it.
		return time.NewTimer(p.timeout)
	}
	return t
}

// Release frees timer and returns it to the pool
func (p *Pool) Release(t *time.Timer) {
	if !t.Stop() {
		// Collect possibly added time from the channel
		// if timer has been stopped and nobody collected its' value.
		select {
		case <-t.C:
		default:
		}
	}

	p.Put(t)
}
