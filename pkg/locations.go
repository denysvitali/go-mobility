package mobility

type LocationsReq = TripModeReq

type LocationsResp struct {
	Locations []Location `json:"locations"`
}