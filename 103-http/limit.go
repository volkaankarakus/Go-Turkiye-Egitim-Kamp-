package main

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

// ** Burada kullanilmasi gereken 2 tane struct var
// ** /GoTurkiye/asdasdsd/adasdas/asdasds icin 3 dakikada 10 tane istek kabul etsin

type Limit struct {
	count int
	tll   time.Time
}

type LimitProxy struct {
	key   string        // buradaki key /GoTurkiye
	limit int           // /GoTurkiye'de 10 request alabilirsin
	ttl   time.Duration // /GoTurkiye'de 3 dakika sayicak.
}

// ** burasi ayni anda birden fazla erisilebilir hal alicagindan burası Thread Safe olmayacak.
// ** Shared variable(global variable)'a ayni anda birden fazla yerden erismeye calistigimizda hata alicaz. (asagidaki global)
var counter = map[string]*Limit{}

// ** Bizim burada 102 egitimdeki gibi mutex kullanimina ihtiyacimiz var.Bu sayede birisi key'e eristiginde locklicak, o birakinca digeri erisebilecek.

func (limitProxy LimitProxy) Accept(key string) bool {
	return limitProxy.key == key
}

func NewLimitProxy(key string, limit int, ttl time.Duration) LimitProxy {
	return LimitProxy{
		key:   key,
		limit: limit,
		ttl:   ttl,
	}
}

func (limitProxy LimitProxy) Proxy(context *fiber.Ctx) error {
	path := context.Path() // buradaki path /GoTurkiye/sdfsddsf/sdfsdfsdf'a denk geliyor.

	if response, ok := counter[path]; ok && response.count >= limitProxy.limit && response.tll.After(time.Now()) {
		context.Response().SetStatusCode(429) // TOO MANY REQUEST HTTP ERROR : 429
		fmt.Printf("request limit reached for %s \n", path)
		return nil
	} else if !ok {
		counter[path] = &Limit{
			count: 0, // counteri degistiricegimiz icin Limit'i pointerlı tuttuk.
			tll:   time.Now().Add(limitProxy.ttl),
		}
	}

	context.SendString("Go Turkiye - 103 Http Package")
	counter[path].count++
	return nil
}

// ** build alip 4 kere calistirmayi deneyelim
// go build .
// ./103-http

// bir diger terminalde : 4x curl -v http://localhost:3000/user/sdfsdfsddsf
// OUTPUT : < HTTP/1.1 429 Too Many Requests
