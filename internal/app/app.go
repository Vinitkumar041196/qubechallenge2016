package app

import (
	"distributor-manager/internal/store"
	"distributor-manager/internal/types"
	"fmt"
)

type App struct {
	countryStore store.CountryStorage
	distStore    store.DistributorStorage
}

func NewApp(countryStore store.CountryStorage, distStore store.DistributorStorage) *App {
	return &App{
		countryStore: countryStore,
		distStore:    distStore,
	}
}

func (a *App) PutDistributor(dist *types.PutDistributorRequest) error {
	err := validatePutDistributorReq(dist)
	if err != nil {
		return err
	}

	distributor := dist.ToDistributor()

	err = a.distStore.PutDistributorByCode(distributor)
	if err != nil {
		return err
	}

	return nil
}

func validatePutDistributorReq(dist *types.PutDistributorRequest) error {
	if dist == nil {
		return fmt.Errorf("invalid input")
	}
	if dist.Code == "" {
		return fmt.Errorf("code cannot be empty")
	}

	if len(dist.SubDistributors) > 0 {
		for _, subD := range dist.SubDistributors {
			err := validatePutDistributorReq(subD)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (a *App) GetDistributor(code string) (*types.Distributor, error) {
	if code == "" {
		return nil, fmt.Errorf("code cannot be empty")
	}

	dist, err := a.distStore.GetDistributorByCode(code)
	if err != nil {
		return nil, err
	}

	return dist, nil
}
