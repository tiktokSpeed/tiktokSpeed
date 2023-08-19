package md5

import (
	"crypto/md5"
	"encoding/hex"
)

func MD5(src string) string {
	d := []byte(src)
	m := md5.New()
	m.Write(d)
	return hex.EncodeToString(m.Sum(nil))
}
