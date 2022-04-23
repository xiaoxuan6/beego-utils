package md5

import (
	"crypto/md5"
	"fmt"
)

func Parse(str []byte) string {
	hash := md5.New()
	hash.Write(str)

	return fmt.Sprintf("%x", hash.Sum(nil))
}
