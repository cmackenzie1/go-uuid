package uuid

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"io"
	"strings"
	"time"
)

var (
	errInvalidLength      = errors.New("invalid length")
	errInvalidFormat      = errors.New("invalid format")
	errUnsupportedVersion = errors.New("unsupported version")
	errUnsupportedVariant = errors.New("unsupported variant")
)

// UUID is a 128 bit (16 byte) value defined by RFC9562.
type UUID [16]byte

// Nil represents the zero-value UUID
var Nil UUID

// NewV4 returns a Version 4 UUID as defined in [RFC9562]. Random bits
// are generated using [crypto/rand].
//
//	 0                   1                   2                   3
//	 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
//	+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//	|                           random_a                            |
//	+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//	|          random_a             |  ver  |       random_b        |
//	+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//	|var|                       random_c                            |
//	+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//	|                           random_c                            |
//	+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//
// [RFC9562]: https://www.rfc-editor.org/rfc/rfc9562.html#name-uuid-version-7
func NewV4() (UUID, error) {
	var uuid UUID

	// fill entire uuid with secure, random bytes
	_, err := io.ReadFull(rand.Reader, uuid[:])
	if err != nil {
		return UUID{}, err
	}

	// Set version and variant bits
	uuid[6] = (uuid[6] & 0x0f) | 0x40 // Version 4 [0100]
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant is [10]
	return uuid, nil
}

// NewV7 returns a Version 7 UUID as defined in [RFC9562].
// Random bits are generated using [crypto/rand].
//
// This function employs method 3 (Replace Leftmost Random Bits with Increased Clock Precision)
// to increase the clock precision of the UUID. This helps support scenarios where
// several UUIDs are generated within the same millisecond.
//
//	 0                   1                   2                   3
//	 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
//	+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//	|                           unix_ts_ms                          |
//	+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//	|          unix_ts_ms           |  ver  |       rand_a          |
//	+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//	|var|                        rand_b                             |
//	+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//	|                            rand_b                             |
//	+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
//
// [RFC9562]: https://www.rfc-editor.org/rfc/rfc9562.html#name-uuid-version-7
func NewV7() (UUID, error) {
	var uuid UUID

	t := time.Now()
	ms := t.UnixMilli()

	// Extract each byte from the 48-bit timestamp
	uuid[0] = byte(ms >> 40) // Most significant byte
	uuid[1] = byte(ms >> 32)
	uuid[2] = byte(ms >> 24)
	uuid[3] = byte(ms >> 16)
	uuid[4] = byte(ms >> 8)
	uuid[5] = byte(ms) // Least significant byte

	// Calculate sub-millisecond precision for rand_a (12 bits)
	ns := t.UnixNano()

	// Calculate sub-millisecond precision by:
	// 1. Taking nanoseconds modulo 1 million to get just the sub-millisecond portion
	// 2. Multiply by 4096 (2^12) to scale to 12 bits of precision
	// 3. Divide by 1 million to normalize back to a 12-bit fraction
	// This provides monotonically increasing values within the same millisecond
	subMs := ((ns % 1_000_000) * (1 << 12)) / 1_000_000

	// Fill the increased clock precision into "rand_a" bits
	uuid[6] = byte(subMs >> 8)
	uuid[7] = byte(subMs)

	// Fill the rest with random data
	_, err := io.ReadFull(rand.Reader, uuid[8:])
	if err != nil {
		return UUID{}, err
	}

	// Set the version and variant bits
	uuid[6] = (uuid[6] & 0x0f) | 0x70 // Version 7 [0111]
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant [10]

	return uuid, nil
}

// String returns the UUID in "hex-and-dash" string format.
//
//	56c450b3-255d-4a2a-a761-cd1b1bf028e2
func (uuid UUID) String() string {
	var buf [36]byte
	encodeHex(buf[:], uuid)
	return string(buf[:])
}

// URN returns the UUID in "urn:uuid" string format.
//
//	urn:uuid:56c450b3-255d-4a2a-a761-cd1b1bf028e2
func (uuid UUID) URN() string {
	var buf [36]byte
	encodeHex(buf[:], uuid)
	return "urn:uuid:" + string(buf[:])
}

func encodeHex(dst []byte, uuid UUID) {
	hex.Encode(dst, uuid[:4])
	dst[8] = '-'
	hex.Encode(dst[9:13], uuid[4:6])
	dst[13] = '-'
	hex.Encode(dst[14:18], uuid[6:8])
	dst[18] = '-'
	hex.Encode(dst[19:23], uuid[8:10])
	dst[23] = '-'
	hex.Encode(dst[24:], uuid[10:])
}

// Parse a UUID string with or without the dashes.
//
//	u, err := Parse("56c450b3255d4a2aa761cd1b1bf028e2") // no dashes
//	u, err := Parse("56c450b3-255d-4a2a-a761-cd1b1bf028e2") // with dashes
//	u, err := Parse("urn:uuid:56c450b3-255d-4a2a-a761-cd1b1bf028e2") // urn uuid prefix
func Parse(s string) (UUID, error) {
	var x string
	switch len(s) {
	case 32: // uuid: "9178e496ba5c4c108b1513a1c70550d0", len: 32
		x = s[:8] + s[8:13] + s[13:18] + s[18:23] + s[23:]
	case 36: // uuid: "9178e496-ba5c-4c10-8b15-13a1c70550d0", len: 36
		if s[8] != '-' || s[13] != '-' || s[18] != '-' || s[23] != '-' {
			return UUID{}, errInvalidFormat
		}
		x = s[:8] + s[9:13] + s[14:18] + s[19:23] + s[24:]
	case 45: // uuid: "urn:uuid:9178e496-ba5c-4c10-8b15-13a1c70550d0"
		if !strings.HasPrefix(s, "urn:uuid:") || s[17] != '-' || s[22] != '-' || s[27] != '-' || s[32] != '-' {
			return UUID{}, errInvalidFormat
		}
		x = s[9:17] + s[18:22] + s[23:27] + s[28:32] + s[33:]
	default:
		return UUID{}, errInvalidLength
	}

	b, err := hex.DecodeString(x)
	if err != nil {
		return UUID{}, err
	}
	if len(b) != 16 {
		return UUID{}, errInvalidLength
	}

	switch { // assert version
	case (b[6] & 0xf0) == 0x40: // version 4
	case (b[6] & 0xf0) == 0x70: // version 7
	default:
		return UUID{}, fmt.Errorf("%v: %d", errUnsupportedVersion, b[6]&0xf0>>4)
	}

	if (b[8] & 0xc0) != 0x80 { // assert variant
		return UUID{}, errUnsupportedVariant
	}

	var uuid UUID
	copy(uuid[:], b)
	return uuid, nil
}
