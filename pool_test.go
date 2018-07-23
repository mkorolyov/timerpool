package timerpool

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	p := New(time.Second)
	i := p.Get()
	assert.NotNil(t, i, "unexpected nil instead of i")
	timer, ok := i.(*time.Timer)
	assert.True(t, ok, "unexpected type %T", i)
	assert.NotNil(t, timer.C, "timer channel is nil")
	p.Release(timer)

	timer2 := p.Get().(*time.Timer)
	assert.Equal(t, timer, timer2, "not the same timer")
	assert.NotEqual(t, timer, p.Get().(*time.Timer), "the same timer")
}

func TestPool_Acquire(t *testing.T) {
	p := &Pool{timeout: time.Second}
	assert.NotNil(t, p.Acquire(), "unexpected nil timer")

	activeTimer := time.NewTimer(time.Minute)
	p.Put(activeTimer)
	assert.NotEqual(t, activeTimer, p.Acquire(), "active timer should not be returned")
}

func TestPool_Release(t *testing.T) {
	p := Pool{timeout: time.Second}
	activeTimer := time.NewTimer(time.Minute)
	p.Release(activeTimer)
	assert.Equal(t, activeTimer, p.Acquire(), "timer should be stopped and reused")
}
