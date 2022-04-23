package url

import "net/url"

type Url struct {
	url   *url.URL
	query url.Values
}

func ParseUrl(uri string) *Url {
	u := &Url{}
	u.url, _ = url.Parse(uri)
	u.query = u.url.Query()

	return u
}

func (u *Url) AddQuery(key, val string) *Url {
	u.query.Add(key, val)
	return u
}

func (u *Url) AddQueries(queries map[string]string) *Url {
	for key, val := range queries {
		u.AddQuery(key, val)
	}

	return u
}

func (u *Url) GetQuery() url.Values {
	return u.query
}

func (u *Url) GetURL() *url.URL {
	return u.url
}

func (u *Url) Build() *url.URL {
	u.url.RawQuery = u.query.Encode()
	return u.url
}

func (u *Url) BuildStr() string {
	return u.Build().String()
}
