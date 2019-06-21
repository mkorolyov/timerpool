package timerpool

import (
	"testing"
	"time"
)

func Benchmark_WithoutPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		t := time.NewTimer(1)
		select {
		case <- t.C:
		}
	}
}

func Benchmark_WithPool(b *testing.B) {
	p := New(1)
	for i := 0; i < b.N; i++ {
		t := p.Acquire()
		select {
		case <- t.C:
		}
		p.Release(t)
	}
}
