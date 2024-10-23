package parser

import (
	"distributor-manager/internal/types"
	"encoding/csv"
	"os"
)

func ParseCSVToCountries(fileName string) (map[string]*types.Country, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	csvReader := csv.NewReader(file)
	data, err := csvReader.ReadAll()
	if err != nil {
		return nil, err
	}

	var cm = make(map[string]*types.Country)

	if len(data) > 1 {
		for _, row := range data[1:] {
			if len(row) == 6 {
				city := &types.City{
					Code: row[0],
					Name: row[3],
				}

				if country, ok1 := cm[row[2]]; ok1 {

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

					cm[country.Code] = country
				}
			}
		}
	}
	return cm, nil
}
