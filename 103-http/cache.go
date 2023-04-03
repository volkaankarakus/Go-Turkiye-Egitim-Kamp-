package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

type Cache struct {
	body []byte
	ttl  time.Time
}

type CacheProxy struct {
	key string
	ttl time.Duration
}

var cacheModelList = map[string]Cache{}

func NewCacheProxy(key string, ttl time.Duration) CacheProxy {
	return CacheProxy{
		key: key,
		ttl: ttl,
	}
}

func (cacheProxy CacheProxy) Accept(key string) bool {
	return cacheProxy.key == key
}

func (cacheProxy CacheProxy) Proxy(context *fiber.Ctx) error {
	path := context.Path()
	key := context.Params("key")

	// ** Cache'de varsa durumu
	if response, ok := cacheModelList[path]; ok && response.ttl.After(time.Now()) {
		context.Response().SetBody(response.body)
		context.Response().Header.Add(fiber.HeaderCacheControl, fmt.Sprintf("max-age:%d", cacheProxy.ttl/time.Second))
		return nil
	}

	// https://mocki.io/v1/d4d63bce-1809-4250-91ac-0470aa392ca5 -> mock.io'dan bir mock olusturduk. ama suanda yasamıyor.
	// v1/d4d63bce-1809-4250-91ac-0470aa392ca5
	url := "https://mocki.io/" + strings.TrimPrefix(path, "/"+key+"/")
	fmt.Printf("http request redirecting to %s \n", url)

	// ** buradaki kodlar paketten. REVERSE PROXY icin proxy.Do()
	if err := proxy.Do(context, url); err != nil {
		return err
	}

	cacheModel := Cache{
		ttl:  time.Now().Add(cacheProxy.ttl),
		body: context.Response().Body(), // byteArray olarak body'yi verir.
	}

	cacheModelList[path] = cacheModel
	context.Response().Header.Del(fiber.HeaderServer)
	return nil
}
