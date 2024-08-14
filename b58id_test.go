package b58id

import (
	"testing"

	"github.com/google/uuid"
)

func TestBuild(t *testing.T) {
	got, err := Build()
	if err != nil {
		t.Fatal(err)
	}

	if len(got) == 0 {
		t.Fatalf("got len %v want len > 0", len(got))
	}
}

func TestNew(t *testing.T) {
	got := New()

	if len(got) == 0 {
		t.Fatalf("got len %v want len > 0", len(got))
	}
}

func TestFromUUID(t *testing.T) {
	tests := []struct {
		name string
		uuid string
		want string
	}{
		{
			name: "success",
			uuid: "4942b9df-9fec-419e-b382-9fd204079079",
			want: "A3hYyNE2ZXNFJZbTenaB36",
		},
	}

	for _, tt := range tests {
		// parse the UUID string
		uid, err := uuid.Parse(tt.uuid)
		if err != nil {
			t.Fatal(err)
		}

		got, err := FromUUID(uid)
		if err != nil {
			t.Fatal(err)
		}

		if got != tt.want {
			t.Fatalf("got %s want %s", got, tt.want)
		}
	}
}

func TestFromString(t *testing.T) {
	tests := []struct {
		name string
		uuid string
		want string
	}{
		{
			name: "success",
			uuid: "4942b9df-9fec-419e-b382-9fd204079079",
			want: "A3hYyNE2ZXNFJZbTenaB36",
		},
	}

	for _, tt := range tests {
		got, err := FromString(tt.uuid)
		if err != nil {
			t.Fatal(err)
		}

		if got != tt.want {
			t.Fatalf("got %s want %s", got, tt.want)
		}
	}
}
