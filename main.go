package main

import (
	"encoding/json"
	"log"
	"net/http"
)

type Product struct {
	ProductID      int    `json:"productId"`
	Manufacturer   string `json:"manufacturer"`
	Sku            string `json:"sku"`
	Upc            string `json:"upc"`
	PricePerUnit   string `json:"pricePerUnit"`
	QuantityOnHand string `json:"quantityOnHand"`
	ProductName    string `json:"productName"`
}

var productList []Product

func init() {
	productsJSON := `[
	{
		"productId":1,
		"manufacturer":"Johns-Jenkins",
		"sku":"p5z343vds",
		"upc":"939581000000",
		"pricePerUnit":"497.45",
		"quantityOnHand":9703,
		"productName":"sticky note"
	},
	{
		"productId":2,
		"manufacturer":"Hessel, Schimmel and Feeney",
		"sku":"i7v300kmx",
		"upc":"939581000000",
		"pricePerUnit":"282.29",
		"quantityOnHand":9217,
		"productName":"leg warmers"
	},
	{
		"productId":3,
		"manufacturer":"Swaniawski, Bartoletti and Bruen",
		"sku":"q0L657ys7",
		"upc":"111730000000",
		"pricePerUnit":"436.26",
		"quantityOnHand":5905,
		"productName":"leg shade"
	}
]`
	err := json.Unmarshal([]byte(productsJSON), &productList)
	if err != nil {
		log.Fatal(err)
	}
}

type fooHandler struct {
	Message string
}

func (f *fooHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(f.Message))
}

func barHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bar called"))
}

func main() {
	http.Handle("/foo", &fooHandler{Message: "foo called"})
	http.HandleFunc("/bar", barHandler)
	http.ListenAndServe(":5000", nil)
}
