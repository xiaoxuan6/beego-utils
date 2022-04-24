package url

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseUrl(t *testing.T) {
	u := ParseUrl("www.baidu.com")
	u.AddQuery("name", "eto")
	ustr := u.BuildStr()
	assert.Equal(t, "www.baidu.com?name=eto", ustr)

	queryies := map[string]string{
		"name": "vinhson",
		"age":  "18",
	}

	u1 := ParseUrl("www.baidu.com")
	u1.AddQueries(queryies)
	ustr = u1.BuildStr()
	assert.Equal(t, "www.baidu.com?age=18&name=vinhson", ustr)

	u = ParseUrl(ustr)
	query := u.GetQuery().Encode()
	assert.Equal(t, "age=18&name=vinhson", query)

}
