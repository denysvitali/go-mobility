package mobility

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const BaseURL = "https://api.mobility.ch/classic/10/v2/"
const AppKey = "8e2c37eb348204d48736c3724e00487d"

type IdName struct {
	Id int32 `json:"id"`
	Name string `json:"name"`
}

type Role = IdName
type Function = IdName

type Session struct {
	Id int32 `json:"id"`
	Roles []Role `json:"roles"`
	Functions []Function `json:"functions"`
	SessionToken string `json:"session_token"`
}

type Client struct {
	httpClient *http.Client
	appKey     string
	baseUrl    *url.URL

	Session *Session
}

type SessionLogin struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (c *Client) Login(username string, password string) error {
	sessionLogin := SessionLogin{
		Username: username,
		Password: password,
	}
	userPassReq, err := json.Marshal(&sessionLogin)
	if err != nil {
		return err
	}
	req, err := c.createPOST("public/sessions/v2", userPassReq)
	if err != nil {
		return err
	}
	
	res, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid status code: received %d but 200 was expected", res.StatusCode)
	}
	
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	
	return json.Unmarshal(bodyBytes, &c.Session)
}

type TripMode = string
const (
	TripReturn TripMode = "RETURN"
	TripOneWay TripMode = "ONEWAY"
)

type Category struct {
	Count *int `json:"count,omitempty"`
	Description string `json:"description"`
	Id int `json:"id"`
	Key int `json:"key"`
	Name string `json:"name"`
	ShowCarOnResOffers bool `json:"show_car_on_res_offers"`
	SortOrder *int `json:"sort_order"`
	TripModes *[]TripMode
}

type Location struct {
	Categories []Category
	GeoX float64 `json:"geo_x"`
	GeoY float64 `json:"geo_y"`
	Id string `json:"id"`
	IsStation bool `json:"is_station"`
	LocationNameNotInSync bool `json:"location_name_not_in_sync"`
	MaxReservationMin int `json:"max_reservation_min"`
	MinReservationMin int `json:"min_reservation_min"`
	Name string `json:"name"`
	Number int `json:"number"`
	OnewayCapable bool `json:"oneway_capable"`
}

type TripModeReq struct {
	TripModes []TripMode `json:"trip_modes"`
}

func (c *Client) Locations(tripModes []TripMode) ([]Location, error) {
	locationsRequest := LocationsReq{
		TripModes: tripModes,
	}
	locationsRequestBytes, err := json.Marshal(&locationsRequest)
	if err != nil {
		return []Location{}, err
	}
	req, err := c.createPOST("public/locations/v2/all", locationsRequestBytes)
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

func (c *Client) Categories(tripMode []TripMode) ([]Category, error) {
	tripModeBytes, err := json.Marshal(&TripModeReq{TripModes: tripMode})
	if err != nil {
		return nil, err
	}
	categorisesReq, err := c.createPOST("public/settings/v3/categories", tripModeBytes)
	resp, err := c.httpClient.Do(categorisesReq)
	if err != nil {
		return nil, err
	}
	
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid status code %d, expected 200", resp.StatusCode)
	}

	var categoriesResp CategoriesResp
	jd := json.NewDecoder(resp.Body)
	err = jd.Decode(&categoriesResp)
	if err != nil {
		return nil, err
	}
	
	return categoriesResp.Categories, nil
}

func (c *Client) Logout() error {
	if c.Session == nil {
		return nil
	}
	
	logoutReq, err := c.createDELETE("private/sessions");
	if err != nil {
		return err
	}
	
	res, err := c.httpClient.Do(logoutReq)
	if err != nil {
		return err
	}
	
	if res.StatusCode != http.StatusNoContent {
		return fmt.Errorf("invalid status code returned: %d returned but 204 expected", res.StatusCode)
	}

	c.Session = nil
	
	return nil
}

func (c *Client) createDELETE(path string) (*http.Request, error) {
	finalUrl, err := c.baseUrl.Parse(path)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("DELETE", finalUrl.String(), nil)
	if err != nil {
		return nil, err
	}

	if c.Session != nil {
		req.Header.Add("Mobisys-Session-Token", c.Session.SessionToken)
	}

	return req, nil
}

func (c *Client) createPOST(path string, requestBody []byte) (*http.Request, error) {
	finalUrl, err := c.baseUrl.Parse(path)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", finalUrl.String(), bytes.NewReader(requestBody))
	if err != nil {
		return nil, err
	}
	
	if c.Session != nil {
		req.Header.Add("Mobisys-Session-Token", c.Session.SessionToken)
	} else {
		req.Header.Add("App-Key", c.appKey)
	}
	req.Header.Add("Content-Type", "application/json")
	return req, nil
}

func NewClient() Client {
	baseUrl, err := url.Parse(BaseURL)
	if err != nil {
		panic(err)
	}
	return Client{
		httpClient: http.DefaultClient,
		appKey: AppKey,
		baseUrl: baseUrl,
	}
}
