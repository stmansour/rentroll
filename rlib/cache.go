package rlib

import "time"

// SimpleCacheCtx describes a simple context for maintaining a
// data cache where multiple threads of execution will need
// write access to shared memory.
type SimpleCacheCtx struct {
	Expiry time.Duration // default time to live for elements in this cache
	Sem    chan int      // the channel to request write access
	SemAck chan int      // handshaking channel
}

// InitCaches is a single entry point to initialize all simple caches in this
// package.
//-----------------------------------------------------------------------------
func InitCaches() {
	go RARBalCacheController()
	go SecDepBalCacheController()
	go GLAcctCacheController()
	go ARCacheController()
}
