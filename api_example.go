package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// **  API yazip daha sonra rabitMQ ile message publish edicez. Daha sonra veritabanı islemlerine bakicaz.
// **  Daha sonra bir consumer yazicaz.
func main() {
	// ** endpointler cogul olmali
	// ** API yazmak icin HandleFunc kullaniyoruz.
	// ** GO'nun kendi paketini kullandik ama cok efektif baska paketler de var.  (Go Fiber ya da fasthttp)
	http.HandleFunc("/shipments", func(writer http.ResponseWriter, request *http.Request) {
		if request.Method != http.MethodGet && request.Method != http.MethodPost {
			writer.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		// ** POST icin : curl -XPOST localhost:8080/shipments -d '{"ID":3,"Address":"Ankara","Status":"On the way"}' -v olsun
		if request.Method == http.MethodPost {
			// ** post gelmisse veriyi okuyalim
			readedByteArray, error := io.ReadAll(request.Body)
			if error != nil {
				http.Error(writer, error.Error(), http.StatusInternalServerError)
			}

			// ** hata yoksa datayi jsondan ceviricez
			var shipment Shipment
			error2 := json.Unmarshal(readedByteArray, &shipment) // ** neyi ceviriyoruz ve neye ceviriyoruz
			if error2 != nil {
				http.Error(writer, error.Error(), http.StatusInternalServerError)
			}

			fmt.Println(shipment)
			writer.WriteHeader(http.StatusCreated)
			return
		}

		// ** bunu denemek icin Termianalde : curl -XGET localhost:8080/shipments -v 'yi calistir.
		// ** digerleri icin de XPOST, XPUT falan gelicek.
		// writer.WriteHeader(http.StatusOK)

		shipments := []Shipment{
			{
				ID:      0,
				Address: "Mugla",
				Status:  "Ready",
			},
			{
				ID:      1,
				Address: "Izmır",
				Status:  "Prepearing",
			},
			{
				ID:      2,
				Address: "Istanbul",
				Status:  "Delivered",
			},
		}

		// ** yukarıda tanimli shipments icin islemler yapalim
		shipmentsByte, error := json.Marshal(shipments) // serialize eder
		if error != nil {
			http.Error(writer, error.Error(), http.StatusInternalServerError)
			return
		}
		// ** content type dafaultta : text/plain.  Bunu düzeltip yazdiralim
		writer.Header().Set("content-type", "application/json")
		writer.Write(shipmentsByte)

	})

	http.ListenAndServe(":8080", nil)

}

// ** Simdi geriye veri dondurelim
type Shipment struct {
	ID      int
	Address string
	Status  string
}
