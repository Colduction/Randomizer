package randomizer

import "net"

type network struct{}

var Network network

// ref: https://datatracker.ietf.org/doc/html/rfc3513#section-2.5
type UnicastType uint8

const (
	GlobalType UnicastType = iota + 1
	LinkLocalType
	SiteLocalType
	UniqueLocalType
	PrivateType UnicastType = UniqueLocalType
)

// ref: https://datatracker.ietf.org/doc/html/rfc3513#section-2.7
type MulticastScope uint8

const (
	InterfaceLocalScope MulticastScope = 0x1
	LinkLocalScope      MulticastScope = 0x2
	AdminLocalScope     MulticastScope = 0x4
	SiteLocalScope      MulticastScope = 0x5
	OrgLocalScope       MulticastScope = 0x8
	GlobalScope         MulticastScope = 0xE
)

// IPv4Addr generates a random IPv4 address by creating a 4-byte IP
// using a hash-based approach for randomness, ensuring a unique address.
func (network) IPv4Addr() net.IP {
	var (
		b = make(net.IP, 0, net.IPv4len)
		x = DefaultHashPool.Sum64()
	)
	b = append(b,
		byte(x>>0),
		byte(x>>8),
		byte(x>>16),
		byte(x>>24))
	return net.IP(b)
}

// IPv6Addr generates a random IPv6 address by creating a 16-byte IP
// through a hash-based approach, ensuring a unique 128-bit address.
func (network) IPv6Addr() net.IP {
	b := make(net.IP, 0, net.IPv6len)
	for i := 0; i < net.IPv6len; i += 8 {
		b = DefaultHashPool.Sum(b)
	}
	return net.IP(b)
}

// MACAddr generates a random MAC address with configurable local and multicast
// bits. The U/L bit controls whether the address is locally administered, and the
// I/G bit controls whether the address is intended for multicast traffic.
func (network) MACAddr(local, multicast bool) net.HardwareAddr {
	var (
		b = make(net.HardwareAddr, 0, 6)
		x = DefaultHashPool.Sum64()
	)
	b = append(b,
		byte(x>>0),
		byte(x>>8),
		byte(x>>16),
		byte(x>>24),
		byte(x>>32),
		byte(x>>40))
	// Set the U/L bit in the first byte
	if local {
		b[0] = b[0] | 0x02
	} else {
		b[0] = b[0] &^ 0x02
	}
	// Set the I/G bit in the first byte
	if multicast {
		b[0] = b[0] | 0x01
	} else {
		b[0] = b[0] &^ 0x01
	}
	return net.HardwareAddr(b)
}

// IPv6UnicastAddr generates a random IPv6 unicast address of a specified
// unicast type by configuring address prefixes.
func (network) IPv6UnicastAddr(unicastType UnicastType) net.IP {
	b := net.IP{
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
	}
	hash := DefaultHashPool.Sum(nil)
	for i := 0; i < net.IPv6len/2; i++ {
		b[i+net.IPv6len/2] = hash[i]
	}
	switch unicastType {
	case GlobalType:
		b[0] = (b[0] & 0x1F) | 0x20
		b[1] = (hash[1] + hash[2] + hash[3] + hash[4]) & 0xFF
	case LinkLocalType:
		b[0] = 0xFE
		b[1] = 0x80
	case SiteLocalType:
		b[0] = 0xFE
		b[1] = 0xC0
	case UniqueLocalType:
		b[0] = 0xFD
		b[1] = 0x00
	}
	return b
}

// IPv6MulticastAddr generates a random IPv6 multicast address with a
// specified multicast scope, setting the appropriate prefix and scope bits.
func (network) IPv6MulticastAddr(scope MulticastScope) net.IP {
	b := net.IP{
		0xFF, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
	}
	b[1] = (b[1] & 0xF0) | (uint8(scope) & 0x0F)
	for i, hash := 0, DefaultHashPool.Sum(nil); i < net.IPv6len/2; i++ {
		b[i+net.IPv6len/2] = hash[i]
	}
	return b
}
