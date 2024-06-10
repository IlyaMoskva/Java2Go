package models_test

import (
	"javago/models"
	"testing"
)

func TestCity(t *testing.T) {
	expectedName := "Test City"
	temperatures := []float64{-5, 10, 20, 35}
	city := models.NewCity(expectedName, temperatures, true, true)

	t.Run("name", func(t *testing.T) {
		got := city.Name()
		if got != expectedName {
			t.Errorf("Expected '%v', got '%v'", expectedName, got)
		}
	})

	t.Run("temC", func(t *testing.T) {
		cq := createQuery(t, true, true, 2, "")
		got := city.TempC(cq)
		want := temperatures[1]
		if got != want {
			t.Errorf("Expected '%v', got '%v'", expectedName, got)
		}
	})
}

func createQuery(t *testing.T, beach bool, ski bool, month int, name string) models.CityQuery {
	t.Helper()
	cq, err := models.NewQuery(beach, ski, month, name)
	if err != nil {
		t.Fatalf("Error creating city query: %v", err)
	}
	return cq
}
