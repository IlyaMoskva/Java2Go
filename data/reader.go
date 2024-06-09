package data

import (
	"encoding/json"
	"os"
)

type Response struct {
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	HasBeach     bool    `json:"hasBeach"`
	HasMountains bool    `json:"hasMountains"`
	TempC        float64 `json:"tempC"`
}

type DataReader interface {
	ReadData() ([]Response, error)
}

type reader struct {
	path string
}

func NewReader() DataReader {
	return &reader{
		path: "./data/cities.json",
	}
}

func (r *reader) ReadData() ([]Response, error) {
	file, err := os.ReadFile(r.path)
	if err != nil {
		return nil, err
	}

	var data []Response
	err = json.Unmarshal(file, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
