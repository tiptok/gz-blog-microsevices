package cache

import (
	"fmt"
	"github.com/tiptok/gocomm/pkg/cache"
	"github.com/tiptok/gocomm/pkg/cache/gzcache"
	"github.com/tiptok/gocomm/pkg/log"
)

func NewMultiLevelCache(hosts []string, password string) *cache.MultiLevelCache {
	fmt.Println("starting multi level cache...")
	mlCache := cache.NewMultiLevelCacheNew(cache.WithDebugLog(true, func() log.Log {
		return log.DefaultLog
	}))
	mlCache.RegisterCache(gzcache.NewClusterCache(hosts, password))
	return mlCache
}

func NewCachedRepository(c *cache.MultiLevelCache, options ...cache.QueryOption) *cache.CachedRepository {
	return cache.NewCachedRepository(c, options...)
}
