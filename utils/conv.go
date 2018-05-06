package utils

import (
	"encoding/hex"

	"github.com/google/uuid"
)

func HexStringToUUID(s string) uuid.UUID {
	var b []byte
	b, _ = hex.DecodeString(s[2:])
	u, _ := uuid.ParseBytes(b)
	return u
}
