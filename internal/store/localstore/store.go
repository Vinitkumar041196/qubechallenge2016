package localstore

import (
	"distributor-manager/internal/parser"
	"distributor-manager/internal/store"
	"distributor-manager/internal/types"
	"errors"
)

type localStore struct {
	filePath  string
	countries map[string]*types.Country
}

var ErrNotFound error = errors.New("record not found")

func NewLocalCountryStore(filepath string) store.CountryStorage {
	return &localStore{
		filePath:  filepath,
		countries: make(map[string]*types.Country),
	}
}

func (s *localStore) LoadData() error {
	data, err := parser.ParseCSVFile(s.filePath)
	if err != nil {
		return err
	}

	if len(data) > 1 {
		for _, row := range data[1:] {
			if len(row) == 6 {
				city := &types.City{
					Code: row[0],
					Name: row[3],
				}

				if country, ok1 := s.countries[row[2]]; ok1 {

					if province, ok2 := country.Provinces[row[1]]; ok2 {

						if _, ok3 := province.Cities[row[0]]; !ok3 {
							province.Cities[city.Code] = city
						}

					} else {
						country.Provinces[row[1]] = &types.Province{
							Code: row[1],
							Name: row[4],
							Cities: map[string]*types.City{
								city.Code: city,
							},
						}
					}
				} else {
					country = &types.Country{
						Code: row[2],
						Name: row[5],
						Provinces: map[string]*types.Province{
							row[1]: {
								Code: row[1],
								Name: row[4],
								Cities: map[string]*types.City{
									city.Code: city,
								},
							},
						},
					}

					s.countries[country.Code] = country
				}
			}
		}
	}

	return nil
}

func (s *localStore) GetCountryByCode(code string) (*types.Country, error) {
	country, ok := s.countries[code]
	if !ok {
		return nil, ErrNotFound
	}
	return country, nil
}

func (s *localStore) GetProvinceByCode(countryCode, provinceCode string) (*types.Province, error) {
	country, ok := s.countries[countryCode]
	if !ok {
		return nil, ErrNotFound
	}

	province, ok := country.Provinces[provinceCode]
	if !ok {
		return nil, ErrNotFound
	}
	return province, nil
}

func (s *localStore) GetCityByCode(countryCode, provinceCode, cityCode string) (*types.City, error) {
	country, ok := s.countries[countryCode]
	if !ok {
		return nil, ErrNotFound
	}

	province, ok := country.Provinces[provinceCode]
	if !ok {
		return nil, ErrNotFound
	}

	city, ok := province.Cities[cityCode]
	if !ok {
		return nil, ErrNotFound
	}

	return city, nil
}