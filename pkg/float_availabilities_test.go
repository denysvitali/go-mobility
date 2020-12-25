package mobility

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

func TestParseFloatAvailabilities(t *testing.T) {
	f, err := os.Open("../test/float-availabilities.json")
	if err != nil {
		t.Fatal(err)
	}

	var floatAvail FloatAvailabilitiesResp
	jd := json.NewDecoder(f)
	err = jd.Decode(&floatAvail)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range floatAvail.Availabilities {
		fmt.Printf("%s\n", v.CarDetails.Name)
	}
}
