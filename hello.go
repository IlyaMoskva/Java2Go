package main

import (
	"fmt"
	"javago/data"
	"javago/models"
	"javago/printer"
)

func main() {
	fmt.Println("Welcome to Vacation Planner. v.0.2")

	cities, err := models.NewCities(data.NewReader())
	if err != nil {
		fmt.Println("Fatal error while reading json file")
		return
	}

	p := printer.New()
	defer p.Cleanup()
	p.CityHeader()

	for _, c := range cities.ListAll() {
		p.CityDetails(c)
	}

	/* set cities
	lon := models.NewCity("London", 23, false, false)
	bar := models.NewCity("Barcelona", 30, true, false)
	nyc := models.NewCity("New York", 28, true, false)
	ant := models.NewCity("St. Anton", -3, false, true)
	asp := models.NewCity("Aspen", -5, false, true)

	p.CityDetails(lon)
	p.CityDetails(bar)
	p.CityDetails(nyc)
	p.CityDetails(ant)
	p.CityDetails(asp)
	*/

}
