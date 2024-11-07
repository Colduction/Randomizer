package randomizer

import "hash/maphash"

// hashPool is a custom pool that limits the number of maphash.Hash objects.
type hashPool struct {
	pool chan *maphash.Hash
}

// NewHashPool creates a new hashPool with the specified size.
// If size is 0, it returns nil. The pool will preallocate size maphash.Hash
// objects for reuse.
func NewHashPool(size int) *hashPool {
	if size == 0 {
		return nil
	}
	p := &hashPool{
		pool: make(chan *maphash.Hash, size),
	}
	for i := 0; i < size; i++ {
		p.pool <- new(maphash.Hash)
	}
	return p
}

// Get retrieves a maphash.Hash object from the pool.
// If the pool is empty, it creates and returns a new maphash.Hash instance.
func (p *hashPool) Get() *maphash.Hash {
	select {
	case h := <-p.pool:
		return h
	default:
		return new(maphash.Hash)
	}
}

// Put returns a maphash.Hash object to the pool.
// If the pool is full, the hash object is discarded.
func (p *hashPool) Put(h *maphash.Hash) {
	select {
	case p.pool <- h:
	default:
	}
}

// Sum computes the hash of the provided byte slice b and appends the resulting
// hash bytes to b. It uses a maphash.Hash from the pool for the computation.
func (p *hashPool) Sum(b []byte) []byte {
	hash := p.Get()
	x := maphash.Bytes(maphash.MakeSeed(), nil)
	hash.Reset()
	p.Put(hash)
	return append(b,
		byte(x>>0),
		byte(x>>8),
		byte(x>>16),
		byte(x>>24),
		byte(x>>32),
		byte(x>>40),
		byte(x>>48),
		byte(x>>56))
}

// Sum32 generates a random 32-bit number using the hashPool.
// It retrieves a maphash.Hash from the pool, computes the hash, and returns the
// output as a uint32 by masking the lower 32 bits of the result.
// Inspired by: https://qqq.ninja/blog/post/fast-threadsafe-randomness-in-go/#using-hashmaphash
func (p *hashPool) Sum32() uint32 {
	hash := p.Get()
	output := maphash.Bytes(maphash.MakeSeed(), nil)
	hash.Reset()
	p.Put(hash)
	return uint32(output & 0xFFFFFFFF)
}

// Sum64 generates a random 64-bit number using the hashPool.
// It retrieves a maphash.Hash from the pool, computes the hash, and returns the
// output as a uint64.
//Inspired by: https://qqq.ninja/blog/post/fast-threadsafe-randomness-in-go/#using-hashmaphash
func (p *hashPool) Sum64() uint64 {
	hash := p.Get()
	output := maphash.Bytes(maphash.MakeSeed(), nil)
	hash.Reset()
	p.Put(hash)
	return output
}

// DefaultHashPool is a globally accessible hashPool with a preallocated size.
var DefaultHashPool = NewHashPool(50)
