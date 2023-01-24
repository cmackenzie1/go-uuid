package uuid

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"errors"
	"io"
	"time"
)

var (
	errInvalidLength = errors.New("invalid length")
	errInvalidFormat = errors.New("invalid format")
)

type UUID [16]byte

var Nil UUID

// NewV4 returns a UUID Version 4 as defined in RFC4122. Random bits
// are generated using crypto/rand.
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
func NewV4() (UUID, error) {
	var uuid UUID

	// fill entire uuid with secure, random bytes
	_, err := io.ReadFull(rand.Reader, uuid[:])
	if err != nil {
		return [16]byte{}, err
	}

	// Set version and variant bits
	uuid[6] = (uuid[6] & 0x0f) | 0x40 // Version 4 [0100]
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant is [10]
	return uuid, nil
}

// NewV7 returns a UUID Version 7 as defined in the drafted revision for RFC4122.
// Random bits are generated using crypto/rand.
// Due to millisecond resolution of the timestamp, UUIDs generated during the
// same millisecond will sort arbitrarily.
// https://www.ietf.org/archive/id/draft-ietf-uuidrev-rfc4122bis-01.html#name-uuid-version-7
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
func NewV7() (UUID, error) {
	var uuid UUID

	t := time.Now()
	ms := t.UnixMilli()
	binary.BigEndian.PutUint64(uuid[:], uint64(ms<<16)) // lower 48 bits. Right 0 padded

	// Fill the rest with random data
	_, err := io.ReadFull(rand.Reader, uuid[7:])
	if err != nil {
		return Nil, err
	}

	// Set the version and variant bits
	uuid[6] = (uuid[6] & 0x0f) | 0x70 // Version 7 [0111]
	uuid[8] = (uuid[8] & 0x3f) | 0x80 // Variant [10]

	return uuid, nil
}

func (uuid UUID) String() string {
	var buf [36]byte
	encodeHex(buf[:], uuid)
	return string(buf[:])
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

func Parse(s string) (uuid UUID, err error) {
	var x string
	switch len(s) {
	case 32: // uuid: 9178e496ba5c4c108b1513a1c70550d0, len: 32
		x = s[:8] + s[8:13] + s[13:18] + s[18:23] + s[23:]
	case 36: // uuid: 9178e496-ba5c-4c10-8b15-13a1c70550d0, len: 36
		if s[8] != '-' || s[13] != '-' || s[18] != '-' || s[23] != '-' {
			return uuid, errInvalidFormat
		}
		x = s[:8] + s[9:13] + s[14:18] + s[19:23] + s[24:]
	default:
		return uuid, errInvalidLength
	}

	b, err := hex.DecodeString(x)
	if err != nil {
		return uuid, err
	}
	if len(b) != 16 {
		return uuid, errInvalidLength
	}

	copy(uuid[:], b)
	return uuid, nil
}
