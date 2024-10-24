package types

import "strings"

type Distributor struct {
	Code            string                  `json:"code"`
	Permissions     *DistributorPermissions `json:"permissions"`
	SubDistributors []*Distributor          `json:"sub_distributors"`
}

type DistributorPermissions struct {
	Include    []string                       `json:"include"`
	Exclude    []string                       `json:"exclude"`
	IncludeMap map[string]map[string]struct{} `json:"-"`
	ExcludeMap map[string]map[string]struct{} `json:"-"`
}

const CountryKey string = "CTRY"
const ProvinceKey string = "PROV"
const CityKey string = "CITY"

type DistributorPermissionsMap struct {
	IncludeRegions map[string]map[string]struct{}
	ExcludeRegions map[string]map[string]struct{}
}

func (req *DistributorPermissions) ToDistributorPermissionsMap() *DistributorPermissionsMap {
	out := &DistributorPermissionsMap{
		IncludeRegions: generateRegionMaps(req.Include),
		ExcludeRegions: generateRegionMaps(req.Exclude),
	}

	return out
}

func generateRegionMaps(strArr []string) map[string]map[string]struct{} {

	out := map[string]map[string]struct{}{
		CountryKey:  make(map[string]struct{}),
		ProvinceKey: make(map[string]struct{}),
		CityKey:     make(map[string]struct{}),
	}

	if len(strArr) == 0 {
		return out
	}
	for _, in := range strArr {
		parts := strings.Split(in, "-")

		if lenParts := len(parts); lenParts == 1 {
			out[CountryKey][parts[0]] = struct{}{}

		} else if len(parts) == 2 {
			out[ProvinceKey][parts[0]] = struct{}{}
			out[CountryKey][parts[1]] = struct{}{}

		} else if len(parts) == 3 {
			out[CityKey][parts[0]] = struct{}{}
			out[ProvinceKey][parts[1]] = struct{}{}
			out[CountryKey][parts[2]] = struct{}{}
		}
	}
	return out
}
