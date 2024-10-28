package uuid_test

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/cmackenzie1/go-uuid"
)

type exampleJSON struct {
	ID uuid.UUID `json:"id"`
}

func TestNewV4_MarshalJSON(t *testing.T) {
	tests := map[string]struct {
		uuid string
		want string
	}{
		"valid": {
			uuid: "b5ae3fb7-9cf5-4220-b040-069badaa0092",
			want: "{\"id\":\"b5ae3fb7-9cf5-4220-b040-069badaa0092\"}",
		},
		"invalid": {
			uuid: "b5ae3fb7-9cf5-b040",
			want: "{\"id\":\"00000000-0000-0000-0000-000000000000\"}",
		},
		"nil": {
			uuid: "00000000-0000-0000-0000-000000000000",
			want: "{\"id\":\"00000000-0000-0000-0000-000000000000\"}",
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			u, _ := uuid.Parse(tt.uuid)
			ex := exampleJSON{ID: u}
			got, err := json.Marshal(ex)
			if err != nil {
				t.Errorf("json.Marshal() failed: %v", err)
				return
			}
			if !bytes.Equal(got, []byte(tt.want)) {
				t.Errorf("got = %s, wanted = %s", got, tt.want)
			}
		})
	}
}

func TestNewV4_UnmarshalJSON(t *testing.T) {
	tests := map[string]struct {
		uuid    string
		want    func() exampleJSON
		wantErr bool
	}{
		"valid": {
			uuid: "{\"id\":\"b5ae3fb7-9cf5-4220-b040-069badaa0092\"}",
			want: func() exampleJSON {
				u, _ := uuid.Parse("b5ae3fb7-9cf5-4220-b040-069badaa0092")
				return exampleJSON{ID: u}
			},
		},
		"invalid": {
			uuid: "{\"id\":\"b5ae3fb7-9cf5\"}",
			want: func() exampleJSON {
				return exampleJSON{ID: uuid.Nil}
			},
			wantErr: true,
		},
		"nil": {
			uuid: "{\"id\":\"\"}",
			want: func() exampleJSON {
				return exampleJSON{ID: uuid.Nil}
			},
		},
		"null": {
			uuid: "{\"id\":null}",
			want: func() exampleJSON {
				return exampleJSON{ID: uuid.Nil}
			},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := exampleJSON{}
			err := json.Unmarshal([]byte(tt.uuid), &got)
			if (err != nil) != tt.wantErr {
				t.Errorf("json.Unmarshal() failed: %v", err)
				return
			}
			if got.ID != tt.want().ID {
				t.Errorf("got = %s, wanted = %s", got, tt.want())
				return
			}
		})
	}
}
