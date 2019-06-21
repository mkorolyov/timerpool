package timerpool

import (
	"sync"
	"time"
)

// Pool is a timer pool that contains available time.Timer of same time.Duration
type Pool struct {
	p sync.Pool
	timeout time.Duration
}

// New builds new timer pool
func New(timeout time.Duration) *Pool {
	p := Pool{}
	p.p.New = func() interface{} {
		return time.NewTimer(timeout)
	}
	return &p
}

// Acquire loads free timer from pool or creates new one.
func (p *Pool) Acquire() *time.Timer {
	t := p.p.Get().(*time.Timer)
	t.Reset(p.timeout)
	return t
}

// Release frees timer and returns it to the pool if possible
func (p *Pool) Release(t *time.Timer) {
	if !t.Stop() {
		select {
		case <-t.C:
		default:
			return
		}
	}
	p.p.Put(t)
}