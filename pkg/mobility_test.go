package mobility_test

import (
	"fmt"
	mobility "github.com/denysvitali/go-mobility/pkg"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestLogin(t *testing.T) {
	c := mobility.NewClient()
	err := c.Login(os.Getenv("MOBILITY_USERNAME"), os.Getenv("MOBILITY_PASSWORD"))
	if err != nil {
		t.Fail()
	}
	assert.NotNil(t, c.Session)
	fmt.Printf("Session: %v\n", c.Session)
	err = c.Logout()
	if err != nil {
		t.Fail()
	}
	assert.Nil(t, c.Session)
}

func TestLocations(t *testing.T) {
	c := mobility.NewClient()
	locations, err := c.Locations([]mobility.TripMode{mobility.TripOneWay, mobility.TripReturn})
	assert.Nil(t, err)
	assert.Greater(t, len(locations), 0)
	fmt.Printf("locations: %v\n", locations)
}

func TestCategories(t *testing.T) {
	c := mobility.NewClient()
	categories, err := c.Categories([]mobility.TripMode{mobility.TripOneWay, mobility.TripReturn})
	assert.Nil(t, err)
	assert.Greater(t, len(categories), 0)
	fmt.Printf("categories: %v\n", categories)
}