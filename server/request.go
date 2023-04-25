package server

import (
	"net/url"
)

func GetQuery(qr url.Values) map[string]interface{} {
	queryString := make(map[string]interface{}, 0)
	for key, value := range qr {
		queryString[key] = value[0]
	}
	return queryString
}
