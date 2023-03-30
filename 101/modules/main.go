// ** Go'daki paket yonetimi icin : MODULE
// ** github uzerindeki paketleri eger ki module kullanimina uygunsa kullanabiliriz
// ** örnek olarak fiber kütüphanesiyle http server olusturmaya calisalim

package main

// ** module kullanmak icin önce initialize edicez : ******   go mod init httpTestWithFiber   *****
// ** go mod initi ana proje path'inde yapiyoruz.

// ** bunu dedikten sonra bize go.mod dosyasını olusturdu.
// ** güncel versiyonlari cekebilmek icin ****** go get ****** ( proje version gecislerini kontrollu yapmak icin module kullanimina ihtiyacimiz var)
// ** go get dedigimizde go.sum dosyası olusuyor. Bu dosya go.mod'un checkleme kontrolunu yaptıgı bir dosyadır.

// ** Daha sonra uygulama icerisinde kullanmadigimiz dosyalari temizlemek icin ****** go mod tidy *****

// ** Kullandigimiz paketin bir update'i varsa ve bunu kullanmak istiyorsak :
// ** go get -u github.com/gofiber/fiber/v2
import (
	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Listen(":3000")
}
