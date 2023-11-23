package bloom

import (
	"shortUrl/internal/conf"
	"sync"

	"github.com/bits-and-blooms/bloom/v3"
)

var (
	numbers   uint
	errorRate float64
	g_bloom   *bloom.BloomFilter
	once      sync.Once
)

func MustInit(c *conf.Filter) {
	numbers = uint(c.Bloom.ElementNumer)
	errorRate = float64(c.Bloom.ErrorRate)
}

func LocalInit() {
	numbers = 100000
	errorRate = 0.001
}
func GetBloom() *bloom.BloomFilter {
	once.Do(func() {
		g_bloom = bloom.NewWithEstimates(numbers, errorRate)

	})
	return g_bloom
}
