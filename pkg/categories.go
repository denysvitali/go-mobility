package mobility

type CategoriesReq = TripModeReq

type CategoriesResp struct {
	Categories []Category `json:"categories"`
}