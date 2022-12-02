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

type Job struct {
	Name        string `json:"name"`
	Status      string `json:"status"`
	Depedencies []Job  `json:"dependencies"`
}

func printDependencies(dependencies []Job, counter int) {
	for _, job := range dependencies {
		fmt.Println(strings.Repeat("--", counter) + job.Name)
		printDependencies(job.Depedencies, counter+1)
}
}

func main() {
	response, err := http.Get("http://cornelius.deez.re/task/bigquery_table_load_data_analytics_aggregated_offer_scraping-20221017/graph?max=3")

	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(string(responseData))

	var responseObject Job
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.Name)

	printDependencies(responseObject.Depedencies, 1)
}
