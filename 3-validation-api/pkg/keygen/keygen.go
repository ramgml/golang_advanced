package keygen

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"time"
)

func GetUserKey(email string) string {
	s := fmt.Sprintf("%v:%v", time.Now().UnixNano(), email)
	h := sha1.New()
    h.Write([]byte(s))
    sha1_hash := hex.EncodeToString(h.Sum(nil))
	return sha1_hash[:6]
}
