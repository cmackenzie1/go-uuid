package uuid

import (
	"regexp"
	"testing"
)

var (
	stringUUIDRegex = regexp.MustCompile("[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}")
	urnUUIDRegex    = regexp.MustCompile("urn:uuid:[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}")
)

func TestParse(t *testing.T) {
	tests := map[string]struct {
		uuid    string
		want    string
		wantErr bool
	}{
		"with dashes":        {uuid: "53bfe550-4165-4f81-a8e7-c2609579ccc0", want: "53bfe550-4165-4f81-a8e7-c2609579ccc0"},
		"no dashes":          {uuid: "53bfe55041654f81a8e7c2609579ccc0", want: "53bfe550-4165-4f81-a8e7-c2609579ccc0"},
		"urn:uuid prefix":    {uuid: "urn:uuid:53bfe550-4165-4f81-a8e7-c2609579ccc0", want: "53bfe550-4165-4f81-a8e7-c2609579ccc0"},
		"uppercase":          {uuid: "53BFE550-4165-4F81-A8E7-C2609579CCC0", want: "53bfe550-4165-4f81-a8e7-c2609579ccc0"},
		"mixed case":         {uuid: "53bfe550-4165-4f81-A8E7-C2609579CCC0", want: "53bfe550-4165-4f81-a8e7-c2609579ccc0"},
		"invalid urn prefix": {uuid: "abc:1234:53bfe550-4165-4f81-a8e7-c2609579ccc0", want: "00000000-0000-0000-0000-000000000000", wantErr: true},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			uuid, err := Parse(tt.uuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() wantErr = %t, gotErr = %v", tt.wantErr, err)
				return
			}
			if uuid.String() != tt.want {
				t.Errorf("want = %s, got = %s", tt.want, uuid.String())
			}
		})
	}
}

func TestUUID_String(t *testing.T) {
	tests := map[string]struct {
		new func() (UUID, error)
	}{
		"nil": {new: func() (UUID, error) {
			return Nil, nil
		}},
		"version 4": {new: NewV4},
		"version 7": {new: NewV7},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			u, _ := tt.new()
			if !stringUUIDRegex.MatchString(u.String()) {
				t.Errorf("UUID.String(): did not match string regex")
			}
		})
	}
}

func TestUUID_URN(t *testing.T) {
	tests := map[string]struct {
		new func() (UUID, error)
	}{
		"nil": {new: func() (UUID, error) {
			return Nil, nil
		}},
		"version 4": {new: NewV4},
		"version 7": {new: NewV7},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			u, _ := tt.new()
			if !urnUUIDRegex.MatchString(u.URN()) {
				t.Errorf("UUID.URN(): did not match string regex")
			}
		})
	}
}

func TestUUID_Duplicates(t *testing.T) {
	var iterations int = 1e6 // 1 million
	set := make(map[UUID]struct{}, iterations)
	tests := map[string]struct {
		new func() (UUID, error)
	}{
		"version 4": {new: NewV4},
		"version 7": {new: NewV7},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			for i := 0; i < iterations; i++ {
				u, _ := tt.new()
				if _, ok := set[u]; ok {
					t.Errorf("iter %d: duplicate UUID detected!", i)
				} else {
					set[u] = struct{}{}
				}
			}
		})
	}
}

func TestPrint(t *testing.T) {
	u, _ := NewV4()
	t.Logf("v4: %s %v", u, u[:])

	u, _ = NewV7()
	t.Logf("v7: %s %v", u, u[:])
}
