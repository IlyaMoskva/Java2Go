package models

import "javago/data"

type cities struct {
	cityMap map[string]CityTemp
}

type Cities interface {
	ListAll() []CityTemp
}

func NewCities(reader data.DataReader) (Cities, error) {
	d, err := reader.ReadData()
	if err != nil {
		return nil, err
	}
	cmap := make(map[string]CityTemp)
	for _, r := range d {
		cmap[r.Id] = NewCity(r.Name, r.TempC, r.HasBeach, r.HasMountains)
	}

	return &cities{
		cityMap: cmap,
	}, nil
}

func (c cities) ListAll() []CityTemp {
	var cs []CityTemp
	for _, rc := range c.cityMap {
		cs = append(cs, rc)
	}
	return cs
}
