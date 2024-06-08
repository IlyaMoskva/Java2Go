package models

var beachVacationThreshold float64 = 32
var skiVacationThreshold float64 = -2

type city struct {
	name         string
	tempC        float64
	hasBeach     bool
	hasMountains bool
}

// BeachVacationReady implements CityTemp.
func (c *city) BeachVacationReady() bool {
	return c.hasBeach && c.tempC > beachVacationThreshold
}

// SkiVacationReady implements CityTemp.
func (c *city) SkiVacationReady() bool {
	return c.hasMountains &&
		c.tempC < skiVacationThreshold
}

type CityTemp interface {
	Name() string
	TempC() float64
	TempF() float64
	BeachVacationReady() bool
	SkiVacationReady() bool
}

func NewCity(n string, t float64, hasBeach bool, hasMountains bool) CityTemp {
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

func (c city) TempC() float64 {
	return c.tempC
}

func (c city) TempF() float64 {
	return (c.tempC * 9 / 5) + 32
}

/*
func BeachVacationReady(c city) bool {
	return c.tempC > beachVacationThreshold && c.hasBeach
}

func SkiVacationReady(c city) bool {
	return c.tempC < skiVacationThreshold && c.hasMountains
}
*/
