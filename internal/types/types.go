package types

type City struct {
	Code string
	Name string
}

type Province struct {
	Code   string
	Name   string
	Cities map[string]*City
}

type Country struct {
	Code      string
	Name      string
	Provinces map[string]*Province
}
