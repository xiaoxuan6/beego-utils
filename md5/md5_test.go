package md5

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMd5(t *testing.T) {
	str := Parse([]byte("md5"))
	assert.NotEmpty(t, str, "加密失败")

	pstr := ParseString("md5")
	assert.NotEmpty(t, pstr, "加密失败")

	assert.Equal(t, true, str == pstr, "加密生成的结果错误")
}
