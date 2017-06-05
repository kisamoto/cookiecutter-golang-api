package cache

const (
	cacheHit  = "hit"
	cacheMiss = "miss"
)

// LogHit is a simple wrapper function to log when a cache
// hit for a key occurs.
func LogHit(key interface{}) {
	// Call in a Goroutine to avoid slowing application
	go logAction(cacheHit, key)
}

// LogMiss is a simple wrapper function to log when a cache
// miss for a key occurs
func LogMiss(key interface{}) {
	// Call in a Goroutine to avoid slowing application
	go logAction(cacheMiss, key)
}

func logAction(cacheType string, key interface{}) {

}
