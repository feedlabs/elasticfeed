package template

import (
	net "net/url"
)

func QueryParamsCount(url *net.URL) int {
	return len(url.Query())
}
