package models

import "fmt"

type query struct {
	beach bool
	ski   bool
	month int
	name  string
}

type CityQuery interface {
	Ski() bool
	Month() int
	Name() string
	Beach() bool
}

func NewQuery(beach bool, ski bool, month int, name string) (CityQuery, error) {
	q := &query{
		beach: beach,
		ski:   ski,
		month: month,
		name:  name,
	}
	if err := q.validate(); err != nil {
		return nil, err
	}
	return q, nil
}

func (q query) validate() error {
	if q.month < 1 || q.month > 12 {
		return fmt.Errorf("Month is 1..12, got %v", q.month)
	}
	return nil
}

func (q query) Beach() bool {
	return q.beach
}

func (q query) Ski() bool {
	return q.ski
}

func (q query) Month() int {
	return q.month - 1
}

func (q query) Name() string {
	return q.name
}
