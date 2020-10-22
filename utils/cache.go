package utils

import (
	cmap "github.com/orcaman/concurrent-map"
)

var localCache cmap.ConcurrentMap

func init() {
	localCache = cmap.New()
}

func NewCache() *cmap.ConcurrentMap {
	return &localCache
}
