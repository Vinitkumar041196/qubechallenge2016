package app

import (
	"distributor-manager/internal/store"
	"fmt"
	"strings"
)

func invalidInclusionError(region string) error {
	return fmt.Errorf("invalid inclusion : %s", region)
}

func checkRegionValid(store store.DistributorStorage, region string, code string, includeArrayError bool) error {
	if code == "" {
		if includeArrayError {
			return invalidInclusionError(region)
		}
		return nil
	}

	dist, err := store.GetDistributorByCode(code)
	if err != nil {
		return err
	}

	if checkRegionInSlice(region, dist.Permissions.Exclude) {
		return invalidInclusionError(region)
	}

	if checkRegionInSlice(region, dist.Permissions.Include) {
		return nil
	} else {
		includeArrayError = true
	}

	return checkRegionValid(store, region, dist.ParentCode, includeArrayError)
}

func checkRegionInSlice(region string, sl []string) bool {
	rParts := strings.Split(region, "-")
	for _, exReg := range sl {
		for i := 0; i <= len(rParts)-1; i++ {
			if exReg == strings.Join(rParts[i:], "-") {
				return true
			}
		}
	}

	return false
}
