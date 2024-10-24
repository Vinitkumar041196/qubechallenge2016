package app

import (
	"distributor-manager/internal/store"
	"distributor-manager/internal/types"
	"fmt"
	"strings"
)

// generates invalid inclusion error
func invalidInclusionError(rParts []string) error {
	return fmt.Errorf("invalid inclusion : %s", strings.Join(rParts, "-"))
}

// generates invalid region string error
func invalidRegionStringError(region string) error {
	return fmt.Errorf("invalid region string : %s", region)
}

// checks if the region in include list of distributor is valid
// check exclude list at the given distributor
// recursively checks include and exclude list of parents and other ancestors
func checkRegionValidity(distStore store.DistributorStorage, rParts []string, dist *types.Distributor) error {
	//1. check if region is not in own exclude list
	if checkRegionInSlice(rParts, dist.Permissions.Exclude) {
		return invalidInclusionError(rParts)
	}

	//2. check if region is valid based on parent's and ancestor's permissions
	err := checkRegionValidWithAncestorPermissions(distStore, rParts, dist.ParentCode)
	if err != nil {
		return err
	}

	//region is valid
	return nil
}

func checkRegionValidWithAncestorPermissions(store store.DistributorStorage, rParts []string, parentCode string) error {
	//stop recursion on reaching root distributor
	if parentCode == "" {
		return nil
	}

	//get parent distributor
	parent, err := store.GetDistributorByCode(parentCode)
	if err != nil {
		return err
	}

	//check region in exclude list
	if checkRegionInSlice(rParts, parent.Permissions.Exclude) {
		return invalidInclusionError(rParts)
	}

	//check region in include list
	if !checkRegionInSlice(rParts, parent.Permissions.Include) {
		return invalidInclusionError(rParts)
	}

	//call for parent of current parent
	return checkRegionValidWithAncestorPermissions(store, rParts, parent.ParentCode)
}

// check is region or substring of region is in slice sl
func checkRegionInSlice(rParts []string, sl []string) bool {
	for _, exReg := range sl {
		for i := 0; i <= len(rParts)-1; i++ {
			if exReg == strings.Join(rParts[i:], "-") {
				//found
				return true
			}
		}
	}
	//not found
	return false
}

// check if the region string contains valid codes for city, province, state
func checkRegionStringValid(countryStore store.CountryStorage, parts []string) error {
	var err error
	if lenParts := len(parts); lenParts == 1 {
		//check country exists
		_, err = countryStore.GetCountryByCode(parts[0])
	} else if lenParts == 2 {
		//check province exists
		_, err = countryStore.GetProvinceByCode(parts[1], parts[0])
	} else if lenParts == 3 {
		//check city exists
		_, err = countryStore.GetCityByCode(parts[2], parts[1], parts[0])
	}

	return err
}
