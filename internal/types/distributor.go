package types

import "strings"

type Distributor struct {
	Code            string
	Permissions     *DistributorPermissions
	SubDistributors []*Distributor
}

type DistributorPermissions struct {
	IncludeRegions map[string]map[string]struct{}
	ExcludeRegions map[string]map[string]struct{}
}

type PutDistributorRequest struct {
	Code            string                     `json:"code"`
	Permissions     *DistributorPermissionsReq `json:"permissions"`
	SubDistributors []*PutDistributorRequest   `json:"sub_distributors"`
}

func (in *PutDistributorRequest) ToDistributor() *Distributor {
	distributor := new(Distributor)
	distributor.Code = in.Code
	distributor.Permissions = in.Permissions.ToDistributorPermissions()
	if len(in.SubDistributors) == 0 {
		for _, subD := range in.SubDistributors {
			distributor.SubDistributors = append(distributor.SubDistributors, subD.ToDistributor())
		}
	}
	return distributor
}

type DistributorPermissionsReq struct {
	Include []string `json:"include"`
	Exclude []string `json:"exclude"`
}

const CountryKey string = "CTRY"
const ProvinceKey string = "PROV"
const CityKey string = "CITY"

func (req *DistributorPermissionsReq) ToDistributorPermissions() *DistributorPermissions {
	out := &DistributorPermissions{
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
			out[CityKey][parts[0]] = struct{}{}

		} else if len(parts) == 2 {
			out[CityKey][parts[0]] = struct{}{}
			out[ProvinceKey][parts[1]] = struct{}{}

		} else if len(parts) == 3 {
			out[CityKey][parts[0]] = struct{}{}
			out[ProvinceKey][parts[1]] = struct{}{}
			out[CountryKey][parts[2]] = struct{}{}
		}
	}
	return out
}
