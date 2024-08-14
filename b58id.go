package b58id

import (
	"github.com/btcsuite/btcutil/base58"
	"github.com/google/uuid"
)

// New returns a new ID.
//
// Will panic if UUID generation fails.
func New() string {
	id, err := Build()
	if err != nil {
		// this could only happen if the generated UUID was invalid
		panic(err)
	}

	return id
}

// Build returns an ID, and any error that could have occurred when generating a new UUID.
func Build() (string, error) {
	return FromUUID(uuid.New())
}

// FromString returns a base58 encoded ID from a string. The input string must be parsable as a UUID.
func FromString(id string) (string, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return "", err
	}

	return FromUUID(uid)
}

// FromUUID returns a base58 encoded ID from a UUID.
func FromUUID(u uuid.UUID) (string, error) {
	b, err := u.MarshalBinary()
	if err != nil {
		return "", err
	}

	return base58.Encode(b), nil
}
