package lib

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func MaHoa(data string) string {

	secret := "p@ssw0rdhmdl"

	// Create a new HMAC by defining the hash type and the key (as byte array)
	h := hmac.New(sha256.New, []byte(secret))

	// Write Data to it
	h.Write([]byte(data))

	// Get result and encode as hexadecimal string
	sha := hex.EncodeToString(h.Sum(nil))

	return sha
}
