package types

import (
	"strings"
)

type Distributor struct {
	Code        string                  `json:"code"`
	Permissions *DistributorPermissions `json:"permissions"`
	ParentCode  string                  `json:"parent_code"`

	parent *Distributor `json:"-"`
}

type DistributorPermissions struct {
	Include []string `json:"include"`
	Exclude []string `json:"exclude"`
}

const CountryKey string = "CTRY"
const ProvinceKey string = "PROV"
const CityKey string = "CITY"

type DistributorPermissionsMap struct {
	IncludeMap map[string]map[string][]string `json:"-"` //{CTRY:{IN}, PROV:{MH:[IN]}, CITY:{MUMBAI:[MH,IN]}}
	ExcludeMap map[string]map[string][]string `json:"-"`
}

func (req *DistributorPermissions) ToDistributorPermissionsMap() *DistributorPermissionsMap {
	out := &DistributorPermissionsMap{
		IncludeMap: generateRegionMaps(req.Include),
		ExcludeMap: generateRegionMaps(req.Exclude),
	}

	return out
}

func generateRegionMaps(strArr []string) map[string]map[string][]string {
	out := map[string]map[string][]string{}
	if len(strArr) == 0 {
		return out
	}

	for _, in := range strArr {
		parts := strings.Split(in, "-")
		lenParts := len(parts)
		if lenParts == 3 {
			cityMap, ok := out[CityKey]
			if !ok {
				cityMap = make(map[string][]string, 0)
			}
			cityMap[parts[0]] = []string{parts[1], parts[2]}
			out[CityKey] = cityMap
		} else if lenParts == 2 {
			provMap, ok := out[ProvinceKey]
			if !ok {
				provMap = make(map[string][]string, 0)
			}
			provMap[parts[0]] = []string{parts[1]}
			out[ProvinceKey] = provMap
		} else if lenParts == 1 {
			ctryMap, ok := out[CountryKey]
			if !ok {
				ctryMap = make(map[string][]string, 0)
			}
			ctryMap[parts[0]] = nil
			out[CountryKey] = ctryMap
		}
	}
	return out
}
