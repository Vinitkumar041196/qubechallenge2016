package store

import (
	"distributor-manager/internal/types"
)

type CountryStorage interface {
	LoadData() error

	GetCitiesByProvinceCode(provinceCode string) (map[string]*types.City, error)
	GetProvincesByCountryCode(countryCode string) (map[string]*types.Province, error)

	GetCountryByCode(code string) (*types.Country, error)
	GetProvinceByCode(countryCode, provinceCode string) (*types.Province, error)
	GetCityByCode(countryCode, provinceCode, cityCode string) (*types.City, error)
}
