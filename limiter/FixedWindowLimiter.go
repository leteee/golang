package limiter

import (
	"sync"
	"time"
)

type FixedWindowLimiter struct {
	sync.Mutex               //避免并发问题
	limit      int           //窗口请求上限
	window     time.Duration //窗口时间
	counter    int           //计数器
	lastTime   time.Time     //上一次请求时间
}

func NewFixedWindowLimiter(limit int, window time.Duration) *FixedWindowLimiter {
	return &FixedWindowLimiter{
		limit:    limit,
		window:   window,
		lastTime: time.Now(),
	}
}

func (l *FixedWindowLimiter) TryAcquireFixedWindowLimiter() bool {
	l.Lock()
	defer l.Unlock()

	//获取当前时间
	now := time.Now()
	//如果当前时间失效，计数器清零，开启新窗口
	if now.Sub(l.lastTime) > l.window {
		l.counter = 0
		l.lastTime = now
	}

	//若到达请求窗口请求上限，请求失败
	if l.counter >= l.limit {
		return false
	}

	l.counter++
	return true
}
