package main

import (
"encoding/json"
"fmt"
"io/ioutil"
"log"
"net/http"
"os"
)


type Manufacturer struct {
	Name string `json:"name"`
}

func main() {
	response, err := http.Get("https://private-anon-28c2bc22c0-carsapi1.apiary-mock.com/manufacturers")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(responseData))

	var manufacturers []Manufacturer
	json.Unmarshal(responseData, &manufacturers)

	fmt.Println(manufacturers)
	for _, manufacturer := range(manufacturers){
		fmt.Println(manufacturer.Name)
	}
}
