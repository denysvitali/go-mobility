package mobility

type EnergyType struct {
	Electric bool
	Id int
	Name string
}

type OnboardUnit struct {
	BluetoothEnabled bool `json:"bluetooth_enabled"`
	Id string `json:"id"`
	KeyExternal bool `json:"key_external"`
}

type CarDetail struct {
	BrandName string `json:"brand_name"`
	Color *string `json:"color"`
	EnergyType EnergyType `json:"energy_type"`
	ExternalNumber int `json:"external_number"`
	FuelCard *string `json:"fuel_card"`
	HolderId int `json:"holder_id"`
	HolderName string `json:"holder_name"`
	Id int `json:"id"`
	Identification string `json:"identification"`
	IdentificationInternal string `json:"identification_internal"`
	Infos *string `json:"infos"`
	ModelName string `json:"model_name"`
	Name string `json:"name"`
	Number int `json:"number"`
	OnboardUnit OnboardUnit `json:"onboard_unit"`
	RoTypeId int `json:"ro_type_id"`
}

type Availability struct {
	CarDetails CarDetail `json:"car_details"`
	Category Category
}

type FloatAvailabilitiesResp struct {
	Availabilities []Availability
}