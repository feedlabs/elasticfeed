package template

import (
	net "net/url"
)

func HasQueryParams(url *net.URL) bool {
	return len(url.Query()) != 0
}
