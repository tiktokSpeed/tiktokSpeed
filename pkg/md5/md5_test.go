package md5

import (
	"fmt"
	"testing"
)

func TestMD5(t *testing.T) {
	src := "123456"
	md5 := MD5(src)
	fmt.Println(md5)
}
