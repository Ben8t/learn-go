package main

import (
"encoding/json"
"fmt"
"io/ioutil"
"log"
"net/http"
"os"
"strings"
)


func get_manufacturers() []byte {
	response, err := http.Get("https://private-anon-28c2bc22c0-carsapi1.apiary-mock.com/manufacturers")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	return responseData
}

type Manufacturer struct {
	Id 		 int `json:"id"`
	Name     string `json:"name"`
	AvgPrice float32 `json:"avg_price"`
}

func (manufacturer Manufacturer) get_average_price_and_name() string {
	return strings.Title(manufacturer.Name) + " - Average Price : " + fmt.Sprintf("%.2f", manufacturer.AvgPrice) + "€"
}

func main() {

	var responseData = get_manufacturers()

	var manufacturers []Manufacturer
	json.Unmarshal(responseData, &manufacturers)

	for _, manufacturer := range(manufacturers){
		fmt.Println(manufacturer.get_average_price_and_name())
	}
}
