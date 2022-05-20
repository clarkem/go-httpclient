package go_httpclient

import (
	"github.com/clarkem/go-httpclient/gohttp"
)

func basicExample() {
	client := gohttp.New()

	client.Get()
}
