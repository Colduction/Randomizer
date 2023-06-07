package randomizer

func RandInt(min, max int) int {
	out := int(generateMapHashRandNumber())
	if out < 0 {
		out = -out
	}
	return min + out%(max-min)
}

func RandInt8(min, max int8) int8 {
	out := int8(generateMapHashRandNumber())
	if out < 0 {
		out = -out
	}
	return min + out%(max-min)
}

func RandInt16(min, max int16) int16 {
	out := int16(generateMapHashRandNumber())
	if out < 0 {
		out = -out
	}
	return min + out%(max-min)
}

func RandInt32(min, max int32) int32 {
	out := int32(generateMapHashRandNumber())
	if out < 0 {
		out = -out
	}
	return min + out%(max-min)
}

func RandInt64(min, max int64) int64 {
	out := int64(generateMapHashRandNumber())
	if out < 0 {
		out = -out
	}
	return min + out%(max-min)
}
