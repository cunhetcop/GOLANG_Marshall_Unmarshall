package methods

import (
	"encoding/json"
	"fmt"
)

/*
	{ "name" : "Alice Brown", "sku" : "54321", "price" : 199.95, "shipTo" : { "name" : "Bob

Brown", "address" : "456 Oak Lane", "city" : "Pretendville", "state" : "HI", "zip" : "98999" },
"billTo" : { "name" : "Alice Brown", "address" : "456 Oak Lane", "city" : "Pretendville",
"state" : "HI", "zip" : "98999" } }
*/
type Item struct {
	Name   string  `json:"name"`
	SKU    string  `json:"sku"`
	Price  float64 `json:"price"`
	ShipTo struct {
		Name    string `json:"name"`
		Address string `json:"address"`
		City    string `json:"city"`
		State   string `json:"state"`
		Zip     string `json:"zip"`
	} `json:"shipTo"`
	BillTo struct {
		Name    string `json:"name"`
		Address string `json:"address"`
		City    string `json:"city"`
		State   string `json:"state"`
		Zip     string `json:"zip"`
	} `json:"billTo"`
}

func Marshall() {

	jsonData := `{"name":"Alice Brown","sku":"54321","price":199.95,"shipTo":{"name":"Bob Brown","address":"456 Oak Lane","city":"Pretendville","state":"HI","zip":"98999"},"billTo":{"name":"Alice Brown","address":"456 Oak Lane","city":"Pretendville","state":"HI","zip":"98999"}}`

	var item Item
	err := json.Unmarshal([]byte(jsonData), &item)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v\n",item)
fmt.Println(item.BillTo.Name)
}
