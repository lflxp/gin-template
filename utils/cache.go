package utils

import (
	"time"

	cmap "github.com/orcaman/concurrent-map"
	"github.com/prometheus/client_golang/prometheus"
)

var (
	localCache cmap.ConcurrentMap
	// 统计缓存个数,分别是cache和redis
	cacheCount = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "Cache_Count",
			Help: "Count how many value in each cache",
		},
		[]string{
			"source",
		},
	)

	// TODO: 统计缓存查询次数
)

func init() {
	prometheus.MustRegister(cacheCount)
}

// 本地缓存Cli
func NewCacheCli() *cmap.ConcurrentMap {
	if localCache == nil {
		// 监控缓存个数
		go func() {
			for {
				cacheCount.WithLabelValues("cache").Set(float64(NewCacheCli().Count()))
				time.Sleep(3 * time.Second)
			}
		}()
		localCache = cmap.New()
	}
	return &localCache
}
