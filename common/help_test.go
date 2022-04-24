package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestEmail(t *testing.T) {
	re := IsEmail([]byte("123@qq.com"))
	assert.Equal(t, true, re, "邮箱格式错误")

	r := IsEmail([]byte("123"))
	assert.Equal(t, false, r, "邮箱格式错误")
}

func TestIFTHEN(t *testing.T) {
	a := 1
	b := 2

	res := IFTHEN(a > b, a, b)
	assert.Equal(t, 2, res)

	res = IFTHEN(a < b, a, b)
	assert.Equal(t, 1, res)
}

func TestMapOperatorIf(t *testing.T) {
	m := map[string]interface{}{
		"name": "eto",
		"age":  18,
	}

	data := MapOperatorIf(m, "name", "default")
	assert.Equal(t, "eto", data)

	data = MapOperatorIf(m, "names", "default")
	assert.Equal(t, "default", data)
}

func TestMapExists(t *testing.T) {
	data := make(map[string]interface{})
	data["name"] = "eto"
	data["age"] = "6"

	e1 := MapExists(data, "name")
	assert.Equal(t, true, e1, "data 中存在 key=name")

	e2 := MapExists(data, "eto")
	assert.Equal(t, false, e2, "data 中不存在 key=eto")
}

func TestMapVal(t *testing.T) {
	data := make(map[string]interface{})
	data["name"] = "eto"
	data["age"] = "6"

	val := MapVal(data, "name")
	assert.Equal(t, "eto", val, "val 取值错误")
	assert.NotEqual(t, "name", val, "val 取值错误")
}
