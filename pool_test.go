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
}

func TestPool_Acquire(t *testing.T) {
	p := &Pool{timeout: time.Second}
	assert.NotNil(t, p.Acquire(), "unexpected nil timer")

	activeTimer := time.NewTimer(time.Minute)
	p.Put(activeTimer)
	assert.NotEqual(t, activeTimer, p.Acquire(), "active timer should not be returned")
}
