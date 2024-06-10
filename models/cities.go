package models

import (
	"javago/data"
	"sort"
	"strings"
)

type cities struct {
	cityMap map[string]CityTemp
}

type Cities interface {
	Filter(q CityQuery) []CityTemp
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

func (c cities) Filter(q CityQuery) []CityTemp {
	if !q.Beach() && !q.Ski() && q.Name() != "" {
		return c.listAll()
	}
	return c.filterHelper(q)
}

func (c cities) filterHelper(q CityQuery) []CityTemp {
	var cs []CityTemp
	for _, rc := range c.cityMap {
		if matchFilter(rc, q) {
			cs = append(cs, rc)
		}
	}
	SortAlphabetically(cs)
	return cs
}

func matchFilter(rc CityTemp, q CityQuery) bool {
	if q.Beach() && rc.BeachVacationReady(q) {
		return true
	}
	if q.Ski() && rc.SkiVacationReady(q) {
		return true
	}
	if q.Name() != "" && strings.Contains(
		strings.ToLower(rc.Name()), q.Name()) {
		return true
	}
	return false
}

func (c cities) listAll() []CityTemp {
	var cs []CityTemp
	for _, rc := range c.cityMap {
		cs = append(cs, rc)
	}
	SortAlphabetically(cs)
	return cs
}

// Sorts the slice in abc order, as map is not ordered
func SortAlphabetically(cs []CityTemp) {
	sort.Slice(cs, func(i, j int) bool {
		return cs[i].Name() < cs[j].Name()
	})
}
