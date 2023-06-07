package randomizer

func RandUint(min, max uint) uint {
	out := uint(generateMapHashRandNumber())
	return min + out%(max-min)
}

func RandUint8(min, max uint8) uint8 {
	out := uint8(generateMapHashRandNumber())
	return min + out%(max-min)
}

func RandUint16(min, max uint16) uint16 {
	out := uint16(generateMapHashRandNumber())
	return min + out%(max-min)
}

func RandUint32(min, max uint32) uint32 {
	out := uint32(generateMapHashRandNumber())
	return min + out%(max-min)
}

func RandUint64(min, max uint64) uint64 {
	out := generateMapHashRandNumber()
	return min + out%(max-min)
}
