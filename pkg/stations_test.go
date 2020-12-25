package mobility_test

import (
	"fmt"
	mobility "github.com/denysvitali/go-mobility/pkg"
	"testing"
)

func TestSearchStation(t *testing.T) {
	c := mobility.NewClient()
	searchString := "Zurich"
	stations, err := c.Stations(mobility.StationSearchQuery{
		SearchString: &searchString,
		TripModes:    &[]mobility.TripMode{mobility.TripReturn},
	})
	
	if err != nil {
		t.Fatal(err)
	}
	
	for _, v := range stations {
		fmt.Printf("%s\n", v.Name)
	}
}
