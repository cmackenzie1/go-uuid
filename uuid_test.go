package uuid

import (
	"fmt"
	"testing"
	"time"
)

func TestParse(t *testing.T) {
	tests := map[string]struct {
		uuid    string
		want    string
		wantErr bool
	}{
		"with dashes": {uuid: "53bfe550-4165-4f81-a8e7-c2609579ccc0", want: "53bfe550-4165-4f81-a8e7-c2609579ccc0"},
		"no dashes":   {uuid: "53bfe55041654f81a8e7c2609579ccc0", want: "53bfe550-4165-4f81-a8e7-c2609579ccc0"},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			uuid, err := Parse(tt.uuid)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() wantErr = %t, gotErr = %v", tt.wantErr, err)
			}
			if uuid.String() != tt.want {
				t.Errorf("want = %s, got = %s", tt.want, uuid.String())
			}
		})
	}
}

func TestNewV7(t *testing.T) {
	uuid1, _ := NewV7()
	time.Sleep(time.Millisecond)
	uuid2, _ := NewV7()
	time.Sleep(time.Millisecond)
	uuid3, _ := NewV7()

	fmt.Println(uuid1)
	fmt.Println(uuid2)
	fmt.Println(uuid3)
}
