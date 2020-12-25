package mobility

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type StationSearchQuery struct {
	SearchString *string `json:"search_string,omitempty"`
	TripModes *[]TripMode `json:"trip_modes,omitempty"`
}

func (c *Client) Stations(ssq StationSearchQuery) ([]Location, error) {
	locationsRequestBytes, err := json.Marshal(&ssq)
	if err != nil {
		return []Location{}, err
	}
	req, err := c.createPOST("public/locations/v2/stations", locationsRequestBytes)
	if err != nil {
		return []Location{}, err
	}

	res, err := c.httpClient.Do(req)
	if err != nil {
		return []Location{}, err
	}

	if res.StatusCode != http.StatusOK {
		return []Location{}, fmt.Errorf("invalid status code received: %d, 200 was expected", res.StatusCode)
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return []Location{}, err
	}

	var locations LocationsResp
	err = json.Unmarshal(bodyBytes, &locations)
	if err != nil {
		return []Location{}, err
	}

	return locations.Locations, nil
}
