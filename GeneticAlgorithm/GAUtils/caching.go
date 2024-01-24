package GAUtils

import (
	"fmt"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Analysis"
	"github.com/bcdannyboy/SecuritiesAnalysisTrader/Optimization"
	"hash/fnv"
	"strconv"
	"sync"
)

const MaxCacheEntries = 200

var ScoreCache = make(map[string]float64)
var CacheMutex sync.RWMutex

// AddToCache safely adds a new entry to the cache and checks if the cache needs to be cleared
func AddToCache(key string, value float64) {
	CacheMutex.Lock() // Lock for writing
	fmt.Printf("current cache size pre-store: %d\n", len(ScoreCache))
	if len(ScoreCache) > MaxCacheEntries {
		clearCache()
	}

	ScoreCache[key] = value
	CacheMutex.Unlock() // Unlock after writing
}

// clearCache safely clears the cache
func clearCache() {
	// Assuming clearCache is only called from within a locked section in AddToCache
	ScoreCache = make(map[string]float64) // Reset the cache
}

// GetFromCache safely retrieves an entry from the cache
func GetFromCache(key string) (float64, bool) {
	CacheMutex.RLock() // Lock for reading
	cacheLen := len(ScoreCache)
	value, exists := ScoreCache[key]
	CacheMutex.RUnlock() // Unlock after reading

	if !exists {

	}

	if cacheLen > MaxCacheEntries {
		CacheMutex.Lock() // Lock for writing
		clearCache()
		CacheMutex.Unlock() // Unlock after writing

	}
	return value, exists
}

func GenerateCacheKey(weights *Optimization.SecurityAnalysisWeights, companies []Analysis.CompanyData) string {
	hash := fnv.New64a()

	fallbackKey := fmt.Sprintf("%v%v", weights, companies)
	hash.Write([]byte(fallbackKey))

	cacheKey := strconv.FormatUint(hash.Sum64(), 10)

	return cacheKey
}
