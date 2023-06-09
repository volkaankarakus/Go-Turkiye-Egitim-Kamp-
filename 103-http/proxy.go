package main

import (
	"time"

	"github.com/gofiber/fiber/v2"
)

// ** Proxy interfaceini implement eden structlarin 2 tane receiver'i olmak zorunda
type Proxy interface {
	Accept(key string) bool   // string bir deger alip bool dondurur.
	Proxy(c *fiber.Ctx) error // context alip error dondurur.
}

// ILK OLARAK RATE LIMITING PROXYSI YAPICAZ
var Proxies = []Proxy{
	// NewLimitProxy("user", 3, 3*time.Minute),
	NewCacheProxy("user", 3*time.Minute),  
	NewCacheProxy("event", 3*time.Second),
}

func ProxyHandler(c *fiber.Ctx) error {
	for _, i := range Proxies {
		if i.Accept(c.Params("key")) {
			return i.Proxy(c)
		}
	}
	c.Response().SetStatusCode(404)
	return nil
}
