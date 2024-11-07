package randomizer

type Integers interface {
	SignedIntegers | UnsignedIntegers
}

type SignedIntegers interface {
	~int8 | ~int16 | ~int | ~int32 | ~int64
}

type UnsignedIntegers interface {
	~uint8 | ~uint16 | ~uint | ~uint32 | ~uint64 | ~uintptr
}

// Int generates a random signed integer of type T.
func Int[T SignedIntegers]() T {
	return T(DefaultHashPool.Sum64())
}

// IntInterval generates a random signed integer of type T within a specified range.
func IntInterval[T SignedIntegers](min, max T) T {
	out := T(DefaultHashPool.Sum64())
	if out < 0 {
		out = -out
	}
	return min + out%(max-min)
}

// Uint generates a random unsigned integer of type T.
func Uint[T UnsignedIntegers]() T {
	return T(DefaultHashPool.Sum64())
}

// UintInterval generates a random unsigned integer of type T within a specified range.
func UintInterval[T UnsignedIntegers](min, max T) T {
	out := T(DefaultHashPool.Sum64())
	return min + out%(max-min)
}

// Float32 generates a random 32-bit float value in the range [0, 1) using the hashPool.
// It retrieves a random 32-bit number from the hash pool, masks the result to retain
// only the lower 24 bits, and converts it into a float32 by dividing by 2^24 to normalize
// the value into the range [0, 1).
func Float32() float32 {
	return float32(DefaultHashPool.Sum32()<<8>>8) / (1 << 24)
}

// Float64 generates a random 64-bit float value in the range [0, 1) using the hashPool.
// It retrieves a random 64-bit number from the hash pool, masks the result to retain
// only the lower 53 bits, and converts it into a float64 by dividing by 2^53 to normalize
// the value into the range [0, 1).
func Float64() float64 {
	return float64(DefaultHashPool.Sum64()<<11>>11) / (1 << 53)
}
