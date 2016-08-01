package cache

import (
	"strings"
	"time"
)

func InitCache(host string, idle, max int, toc, tor, tow, defaultExpiration time.Duration) {
	if strings.HasPrefix(host, "redis://") {
		Instance = NewRedisCacheURL(host, idle, max, toc, tor, tow, defaultExpiration)
	} else {
		Instance = NewRedisCache(host, idle, max, toc, tor, tow, defaultExpiration)
	}
}

func InitMemoryCache(defaultExpiration time.Duration) {
	Instance = NewInMemoryCache(defaultExpiration)
}
