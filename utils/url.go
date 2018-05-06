package utils

import "net/url"

func SetUrlQueryString(urlString string, key string, value string) string {
	u, _ := url.Parse(urlString)
	v := u.Query()
	v.Set(key, value)

	return u.Scheme + "://" + u.Host + u.Path + "?" + v.Encode()
}
