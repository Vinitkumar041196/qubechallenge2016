package app

import (
	"distributor-manager/internal/store"
	"distributor-manager/internal/types"
	"fmt"
	"strings"
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
		err := a.checkRegionStringValid(include)
		if err != nil {
			return invalidRegionStringError(include)
		}

		err = a.checkRegionValidity(include, dist)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) checkRegionValidity(region string, dist *types.Distributor) error {
	if checkRegionInSlice(region, dist.Permissions.Exclude) {
		return invalidInclusionError(region)
	}

	err := checkRegionValid(a.distStore, region, dist.ParentCode, false)
	if err != nil {
		return err
	}
	return nil
}

func (a *App) checkRegionStringValid(region string) error {
	parts := strings.Split(region, "-")
	lenParts := len(parts)
	if lenParts == 1 {
		_, err := a.countryStore.GetCountryByCode(parts[0])
		if err != nil {
			return err
		}
	} else if lenParts == 2 {
		_, err := a.countryStore.GetProvinceByCode(parts[1], parts[0])
		if err != nil {
			return err
		}
	} else if lenParts == 3 {
		_, err := a.countryStore.GetCityByCode(parts[2], parts[1], parts[0])
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

func (a *App) CheckIsServiceable(req *types.IsServiceableRequest) (bool, error) {
	if req.Code == "" {
		return false, fmt.Errorf("code cannot be empty")
	}

	if req.Region == "" {
		return false, fmt.Errorf("region cannot be empty")
	}

	dist, err := a.distStore.GetDistributorByCode(req.Code)
	if err != nil {
		return false, err
	}

	err = a.checkRegionStringValid(req.Region)
	if err != nil {
		return false, invalidRegionStringError(req.Region)
	}

	err = a.checkRegionValidity(req.Region, dist)
	if err != nil {
		return false, nil
	}
	return true, nil
}
