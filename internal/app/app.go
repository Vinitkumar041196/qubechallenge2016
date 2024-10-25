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

// New App initialisation
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

// Creates or Updates a distributor
func (a *App) PutDistributor(dist *types.Distributor) error {
	//Validate input
	if dist == nil {
		return fmt.Errorf("invalid input")
	}

	if dist.Code == "" {
		return fmt.Errorf("code cannot be empty")
	}

	if dist.ParentCode != "" {
		pDist, err := a.distStore.GetDistributorByCode(dist.ParentCode)
		if err != nil {
			if strings.Contains(err.Error(), "not found") || pDist == nil {
				return fmt.Errorf("invalid parent code: %s", dist.ParentCode)
			}
			return err
		}
	}

	//validate distributon permissions
	err := a.validateDistributorPermissions(dist)
	if err != nil {
		return err
	}

	//store distributor
	err = a.distStore.PutDistributorByCode(dist)
	if err != nil {
		return err
	}

	return nil
}

// validates the regions in include list with self and ancestor permissions
func (a *App) validateDistributorPermissions(dist *types.Distributor) error {

	for _, region := range dist.Permissions.Include {
		//Split region string
		rParts := strings.Split(region, "-")

		//check if the codes in region string are valid
		err := checkRegionStringValid(a.countryStore, rParts)
		if err != nil {
			return invalidRegionStringError(region)
		}

		//Check if region is valid for given distributor
		err = checkRegionValidity(a.distStore, rParts, dist)
		if err != nil {
			return err
		}
	}

	return nil
}

// Get distributor by code
func (a *App) GetDistributor(code string) (*types.Distributor, error) {
	//Validate input
	if code == "" {
		return nil, fmt.Errorf("code cannot be empty")
	}

	//get distributor from store
	dist, err := a.distStore.GetDistributorByCode(code)
	if err != nil {
		return nil, err
	}

	return dist, nil
}

// Checks if region is serviceable for a given ditributor
func (a *App) CheckIsServiceable(req *types.IsServiceableRequest) (bool, error) {
	//Validate input
	if req.Code == "" {
		return false, fmt.Errorf("code cannot be empty")
	}

	if req.Region == "" {
		return false, fmt.Errorf("region cannot be empty")
	}

	//get distributor from store
	dist, err := a.distStore.GetDistributorByCode(req.Code)
	if err != nil {
		return false, err
	}

	//Split region string
	rParts := strings.Split(req.Region, "-")

	//check if the codes in region string are valid
	err = checkRegionStringValid(a.countryStore, rParts)
	if err != nil {
		return false, invalidRegionStringError(req.Region)
	}

	//Check if region is valid for given distributor
	err = checkRegionValidity(a.distStore, rParts, dist)
	if err != nil {
		return false, nil
	}

	return true, nil
}

func (a *App) DeleteDistributor(code string) error {
	//Validate input
	if code == "" {
		return fmt.Errorf("code cannot be empty")
	}

	//delete distributor from store
	err := a.distStore.DeleteDistributorByCode(code)
	if err != nil {
		return err
	}

	return nil
}
