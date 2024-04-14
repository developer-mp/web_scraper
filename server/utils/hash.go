package hash

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"time"
)

func GenerateHashId() string {
    timestamp := time.Now().UnixNano() / int64(time.Millisecond)

	randomBytes := make([]byte, 8)
    _, err := rand.Read(randomBytes)
    if err != nil {
        return ""
    }

    data := append([]byte(string(rune(timestamp))), randomBytes...)

    hasher := sha256.New()
    hasher.Write(data)
    hash := hex.EncodeToString(hasher.Sum(nil))

    return hash
}