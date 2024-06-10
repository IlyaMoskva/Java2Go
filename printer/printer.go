package printer

import (
	"fmt"
	"javago/models"
	"os"
	"text/tabwriter"
)

type Printer struct {
	w *tabwriter.Writer
}

func New() *Printer {
	w := tabwriter.NewWriter(os.Stdout, 3, 0, 3, ' ', tabwriter.TabIndent)
	return &Printer{
		w: w,
	}
}

func (p *Printer) CityHeader() {
	fmt.Fprintln(p.w, "Name\tTempC\tTempF\tBeach ready\tSki ready")
}

func (p *Printer) CityDetails(c models.CityTemp, q models.CityQuery) {
	fmt.Fprint(p.w,
		c.Name(), "\t",
		c.TempC(q), "\t",
		c.TempF(q), "\t",
		c.BeachVacationReady(q), "\t",
		c.SkiVacationReady(q), "\n")
}

func (p *Printer) Cleanup() {
	p.w.Flush()
}
