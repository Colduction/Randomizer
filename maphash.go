package randomizer

import (
	"hash/maphash"
	"sync"
)

var (
	hashMapPool = sync.Pool{
		New: func() interface{} {
			return new(maphash.Hash)
		},
	}
)

// Inspired from: https://qqq.ninja/blog/post/fast-threadsafe-randomness-in-go/#using-hashmaphash
func generateMapHashRandNumber() uint64 {
	hash := hashMapPool.Get().(*maphash.Hash)
	defer hashMapPool.Put(hash)
	hash.SetSeed(maphash.MakeSeed())
	return hash.Sum64()
}
