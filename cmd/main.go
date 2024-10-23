package main

import (
	"distributor-manager/internal/app"
	"distributor-manager/internal/server"
	"distributor-manager/internal/store/localstore"
	"log"
)

func main() {
	log.SetFlags(log.Lshortfile | log.Ldate)
	log.Println("Initializing server...")

	countryStore := localstore.NewLocalCountryStore("data/cities.csv")
	distStore := localstore.NewLocalDistributorStore()

	distApp, err := app.NewApp(countryStore, distStore)
	if err != nil {
		log.Fatal("could not initialize app error:", err)
	}

	server := server.NewServer(distApp)
	err = server.Start()
	if err != nil {
		log.Fatal("failed to start server error:", err)
	}
}
