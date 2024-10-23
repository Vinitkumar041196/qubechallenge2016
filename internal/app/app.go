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
	err := validatePutDistributorReq(dist)
	if err != nil {
		return err
	}

	err = a.distStore.PutDistributorByCode(dist)
	if err != nil {
		return err
	}

	return nil
}

func validatePutDistributorReq(dist *types.Distributor) error {
	if dist == nil {
		return fmt.Errorf("invalid input")
	}
	if dist.Code == "" {
		return fmt.Errorf("code cannot be empty")
	}

	// TODO: VALIDATION for include and exclude

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
