package types

type Distributor struct {
	Code            string                 `json:"code"`
	SubDistributors []*Distributor         `json:"sub_distributors"`
	Permissions     DistributorPermissions `json:"permissions"`
}

type DistributorPermissions struct {
	IncludeRegions map[string]map[string]struct{}
	ExcludeRegions map[string]map[string]struct{}
}
