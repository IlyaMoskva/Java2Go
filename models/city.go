package models

var beachVacationThreshold float64 = 25
var skiVacationThreshold float64 = -2

type city struct {
	name         string
	tempC        []float64
	hasBeach     bool
	hasMountains bool
}

// BeachVacationReady implements CityTemp.
func (c *city) BeachVacationReady(q CityQuery) bool {
	return c.hasBeach && c.tempC[q.Month()] > beachVacationThreshold
}

// SkiVacationReady implements CityTemp.
func (c *city) SkiVacationReady(q CityQuery) bool {
	return c.hasMountains &&
		c.tempC[q.Month()] < skiVacationThreshold
}

type CityTemp interface {
	Name() string
	TempC(q CityQuery) float64
	TempF(q CityQuery) float64
	BeachVacationReady(q CityQuery) bool
	SkiVacationReady(q CityQuery) bool
}

func NewCity(n string, t []float64, hasBeach bool, hasMountains bool) CityTemp {
	return &city{
		name:         n,
		tempC:        t,
		hasBeach:     hasBeach,
		hasMountains: hasMountains,
	}
}

func (c city) Name() string {
	return c.name
}

func (c city) TempC(q CityQuery) float64 {
	return c.tempC[q.Month()]
}

func (c city) TempF(q CityQuery) float64 {
	return (c.tempC[q.Month()] * 9 / 5) + 32
}
