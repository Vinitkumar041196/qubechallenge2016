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

func NewApp(countryStore store.CountryStorage, distStore store.DistributorStorage) (*App, error) {
	err := countryStore.LoadData()
	if err != nil {
		return nil, err
	}

	return &App{
		countryStore: countryStore,
		distStore:    distStore,
	}, nil
}

func (a *App) PutDistributor(dist *types.Distributor) error {
	err := a.validatePutDistributorReq(dist)
	if err != nil {
		return err
	}

	err = a.distStore.PutDistributorByCode(dist)
	if err != nil {
		return err
	}

	return nil
}

func (a *App) validatePutDistributorReq(dist *types.Distributor) error {
	if dist == nil {
		return fmt.Errorf("invalid input")
	}
	if dist.Code == "" {
		return fmt.Errorf("code cannot be empty")
	}

	for _, include := range dist.Permissions.Include {
		if checkRegionInSlice(include, dist.Permissions.Exclude) {
			return invalidInclusionError(include)
		}

		err := checkRegionValid(a.distStore, include, dist.ParentCode, false)
		if err != nil {
			return err
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

func (a *App) DeleteDistributor(code string) error {
	if code == "" {
		return fmt.Errorf("code cannot be empty")
	}

	err := a.distStore.DeleteDistributorByCode(code)
	if err != nil {
		return err
	}

	return nil
}
