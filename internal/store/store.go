package store

import (
	"distributor-manager/internal/types"
)

type CountryStorage interface {
	LoadData() error

	GetCountryByCode(code string) (*types.Country, error)
	GetProvinceByCode(countryCode, provinceCode string) (*types.Province, error)
	GetCityByCode(countryCode, provinceCode, cityCode string) (*types.City, error)
}

type DistributorStorage interface {
	GetDistributorByCode(code string) (*types.Distributor, error)
	PutDistributorByCode(code string, dist types.Distributor) error
	DeleteDistributorByCode(code string) error
}
