package model

type Address struct {
	City           string      `json:"city"`
	Street_Name    string      `json:"street_name"`
	Street_Address string      `json:"street_address"`
	Zip_Code       string      `json:"zip_code"`
	State          string      `json:"state"`
	Country        string      `json:"country"`
	Coordinates    Coordinates `json:"coordinates"`
}

type Coordinates struct {
	Lat float64 `json:"lat"`
	Lng float64 `json:"lng"`
}
