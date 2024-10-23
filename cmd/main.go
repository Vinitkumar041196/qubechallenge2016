package main

import (
	"distributor-manager/internal/store/localstore"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate)
	log.Println("Initializing server...")

	store := localstore.NewLocalCountryStore("data/cities.csv")
	err := store.LoadData()
	if err != nil {
		log.Fatal("error while loading data to store: ", err)
	}

	log.Println(store.GetCityByCode("GB", "ENG", "SOMER"))
}
