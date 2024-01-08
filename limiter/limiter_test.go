package limiter

import (
	"fmt"
	"testing"
)

func TestTryAcquireFixedWindowLimiter(t *testing.T) {
	l := NewFixedWindowLimiter(2, 1000)
	success := 0
	for i := 0; i < 10; i++ {
		if l.TryAcquireFixedWindowLimiter() {
			success++
		}
	}
	fmt.Printf("请求10次，限流%d", 10-success)
}

func TestTryAcquireTokenBucketLimiter(t *testing.T) {
	l := NewFixedWindowLimiter(2, 1000)
	success := 0
	for i := 0; i < 10; i++ {
		if l.TryAcquireFixedWindowLimiter() {
			success++
		}
	}
	fmt.Printf("请求10次，限流%d", 10-success)
}
