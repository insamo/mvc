package crypto

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"strings"
)

func GenerateMD5Hash(email string) string {
	email = strings.ToLower(strings.TrimSpace(email))
	hash := md5.New()
	hash.Write([]byte(email))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func GenerateSha256Hash(m interface{}) string {
	b, _ := json.Marshal(m)
	hash := sha256.New()
	hash.Write(b)
	return fmt.Sprintf("%x", hash.Sum(nil))
}
