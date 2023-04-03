package main

import (
	"github.com/gofiber/fiber/v2"
)

// ** asagidaki helloyu text.go'dan aldigi icin bunun paket olarak derlenmesi gerekiyor.
// ** Paket derlemek isterken eger GO PATH'inin dısındaysak bir go.mod dosyasına ihtiyacımız var.

// ** go mod init FiberTest gibi bir go.mod isimlendirmesi ve daha sonra go mod tidy yapip calismadigimiz dosyaları temizliyoruz.
// ** Asagidaki kullanim fiber documantasyonundan
// ** fiber ekledigimiz icin tekrar go mod tidy yapalim.

// ** go build . yapalim.
// ** bize bir 103-http dosyasi olusturdu. Bunu calistirmak icin de ./103-http yaparsak
// ┌───────────────────────────────────────────────────┐
// │                   Fiber v2.43.0                   │
// │               http://127.0.0.1:3000               │
// │       (bound on host 0.0.0.0 and port 3000)       │
// │                                                   │
// │ Handlers ............. 2  Processes ........... 1 │
// │ Prefork ....... Disabled  PID .............. 4443 │
// └───────────────────────────────────────────────────┘

// ** Yeni bir terminal acip curl http://localhost:3000 yapalim.
// OUTPUT :  Hello, World!%

// func main() {
// 	app := fiber.New()

// 	app.Get("/", func(c *fiber.Ctx) error {
// 		return c.SendString("Hello, World!")
// 	})

// 	app.Listen(":3000") // localhost:3000 dedikten sonraki her istek 17. satira düser.Burada router tanimi yapildi.
// }

// ** Aldigi istegi cevaba yazan bir http istegi yazalim
// ** asagidaki kodu yazdıktan sonra terminali kapatip tekrar build alalim.
// go build .
// ./103-http yaparsak uygulama ayaga kalkar.
// OUTPUT  :
//  ┌───────────────────────────────────────────────────┐
//  │                   Fiber v2.43.0                   │
//  │               http://127.0.0.1:3000               │
//  │       (bound on host 0.0.0.0 and port 3000)       │
//  │                                                   │
//  │ Handlers ............. 2  Processes ........... 1 │
//  │ Prefork ....... Disabled  PID .............. 5693 │

// ** yeni terminalde curl http://localhost:3000/GoTurkiye yapalim
// OUTPUT : Hello, GoTurkiye!%

// func main() {
// 	app := fiber.New()

// 	app.Get("/:name", func(c *fiber.Ctx) error {
// 		return c.SendString(fmt.Sprintf("Hello, %s!", c.Params("name")))
// 	})
// 	app.Listen(":3000")
// }

// ** proxy.go ile handler yazalim ve daha sonra asagidaki detaylar

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	// ** end point parametre detaylari icin :
	// ** https://docs.gofiber.ipo/guide/routing
	// ** Handler ile request yaratma
	app.Get("/:key/*", ProxyHandler)
	app.Listen(":3000")
	// ** burasi localhost:3000/user/sdfsdfsfsdf/sdfsdfsdf7sdfsdfsdfsdf gibidir. /'tan sonra ne olursa olsun (* ile)

	// ** Bizim denedigimiz sey ornek olarak : google.com/asdsdfdsf/sdfsdfsd'tan gelen response'u
	// **    localhost:3000/user/sdfsdfsfsdf/sdfsdfsdf7sdfsdfsdfsdf'tan dönen response olarak dönmeye calisicaz.

	
}
