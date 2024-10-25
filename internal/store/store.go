package store

import (
	"distributor-manager/internal/types"
)

// Country Store Interface
type CountryStorage interface {
	LoadData() error
	GetCountryByCode(code string) (*types.Country, error)
	GetProvinceByCode(countryCode, provinceCode string) (*types.Province, error)
	GetCityByCode(countryCode, provinceCode, cityCode string) (*types.City, error)
}

// Distributor Store Interface
type DistributorStorage interface {
	GetDistributorByCode(code string) (*types.Distributor, error)
	PutDistributorByCode(dist *types.Distributor) error
	DeleteDistributorByCode(code string) error
}
