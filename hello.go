package main

import (
	"flag"
	"fmt"
	"javago/data"
	"javago/models"
	"javago/printer"
)

func main() {
	fmt.Println("Welcome to Vacation Planner. v.0.2")

	beachReady := flag.Bool("beach", false, "Display only beach ready places")
	skiReady := flag.Bool("ski", false, "Display only ski ready places")
	name := flag.String("name", "", "Display the city with given Name")
	month := flag.Int("month", 0, "Show availability in month")
	flag.Parse()

	cq, err := models.NewQuery(*beachReady, *skiReady, *month, *name)
	if err != nil {
		fmt.Println("Fatal error while reading params: ", err)
		return
	}

	cities, err := models.NewCities(data.NewReader())
	if err != nil {
		fmt.Println("Fatal error while reading json file")
		return
	}

	p := printer.New()
	defer p.Cleanup()
	p.CityHeader()

	// Filter by flags
	cs := cities.Filter(cq)
	for _, c := range cs {
		p.CityDetails(c, cq)
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
