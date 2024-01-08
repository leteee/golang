package limiter

import (
	"lang-go/util"
	"sync"
	"time"
)

type TokenBucketLimiter struct {
	sync.Mutex              //避免并发问题
	capacity      int       //容量
	currentTokens int       //可用令牌
	rate          int       //发放令牌速率/秒
	lastTime      time.Time //上次发放令牌时间
}

func NewTokenBucketLimiter(capacity, rate int) *TokenBucketLimiter {
	return &TokenBucketLimiter{
		capacity: capacity,
		rate:     rate,
		lastTime: time.Now(),
	}
}
func (l *TokenBucketLimiter) TryAcquireTokenBucketLimiter() bool {
	l.Lock()
	defer l.Unlock()

	now := time.Now()
	interval := now.Sub(l.lastTime)
	if interval >= time.Second {
		l.currentTokens = util.MinInt(l.capacity, int(interval/time.Second)*l.rate)
		l.lastTime = now
	}

	if l.currentTokens <= 0 {
		return false
	}

	l.currentTokens--
	return true
}
