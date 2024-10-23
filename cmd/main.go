package main

import (
	"distributor-manager/internal/app"
	"distributor-manager/internal/store/localstore"
	"distributor-manager/internal/types"
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

	err = distApp.PutDistributor(&types.Distributor{
		Code: "1",
		Permissions: &types.DistributorPermissions{
			Include: []string{"IN", "US"},
			Exclude: []string{"UP-IN", "YEOLA-MH-IN"},
		},
		SubDistributors: []*types.Distributor{
			{
				Code: "1-a",
				Permissions: &types.DistributorPermissions{
					Include: []string{"MH-IN"},
					Exclude: []string{"US"},
				},
			},
		},
	})
	if err != nil {
		log.Fatal(" error:", err)
	}

	dist, err := distApp.GetDistributor("1")
	if err != nil {
		log.Fatal(" error:", err)
	}
	subDist, err := distApp.GetDistributor("1-a")
	if err != nil {
		log.Fatal(" error:", err)
	}
	log.Println("Dist", dist, dist.Permissions.Include, dist.Permissions.Exclude)
	log.Println("subDist", subDist, subDist.Permissions.ToDistributorPermissionsMap())
}
