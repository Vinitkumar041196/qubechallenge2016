package main

import (
	"distributor-manager/internal/parser"
	"fmt"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate)
	log.Println("Initializing server...")
	data, err := parser.ParseCSVToCountries("data/cities.csv")
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Println(data)
}
