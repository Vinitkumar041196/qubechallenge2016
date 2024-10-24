package main

import (
	"distributor-manager/internal/app"
	"distributor-manager/internal/server"
	"distributor-manager/internal/store/localstore"
	"log"
)

// @title			Cinema Distributor API- Qube Challenge 2016
// @version		1.0
// @description	Submission for Qube Challenge 2016
// @BasePath		/
// @schemes		http
func main() {
	//set up logger
	log.SetFlags(log.Lshortfile | log.Ldate)
	log.Println("Initializing server...")

	//Store initialization
	countryStore := localstore.NewLocalCountryStore("data/cities.csv")
	distStore := localstore.NewLocalDistributorStore()

	//App initialization
	distApp, err := app.NewApp(countryStore, distStore)
	if err != nil {
		log.Fatal("could not initialize app error:", err)
	}

	//Server initialization
	server := server.NewServer(distApp)

	//Start HTTP server
	err = server.Start()
	if err != nil {
		log.Fatal("failed to start server error:", err)
	}
}
