package timerpool

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	p := New(time.Second)
	i := p.Acquire()
	assert.NotNil(t, i, "unexpected nil instead of i")
	assert.NotNil(t, i.C, "timer channel is nil")
	p.Release(i)

	timer2 := p.p.Get().(*time.Timer)
	assert.Equal(t, i, timer2, "not the same timer")
	assert.NotEqual(t, i, p.p.Get().(*time.Timer), "the same timer")
}

func TestPool_Acquire(t *testing.T) {
	p := New(time.Second)
	assert.NotNil(t, p.Acquire(), "unexpected nil timer")

	activeTimer := p.Acquire()
	assert.NotEqual(t, activeTimer, p.Acquire(), "active timer should not be returned")
}

func TestPool_Release(t *testing.T) {
	p := New(time.Second)
	activeTimer := time.NewTimer(time.Minute)
	p.Release(activeTimer)
	assert.Equal(t, activeTimer, p.Acquire(), "timer should be stopped and reused")
}
