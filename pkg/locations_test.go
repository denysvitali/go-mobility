package mobility

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestParseLocation(t *testing.T) {
	f, err := os.Open("../test/locations.json")
	if err != nil {
		t.Fatal(err)
	}

	var locations LocationsResp
	jd := json.NewDecoder(f)
	err = jd.Decode(&locations)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range locations.Locations {
		fmt.Printf("%s\n", v.Name)
	}
}
