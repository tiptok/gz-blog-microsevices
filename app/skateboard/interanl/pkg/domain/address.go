package domain

type Address struct {
	Address string  `json:"address,omitempty"`
	Lon     float64 `json:"lon,omitempty"`
	Lat     float64 `json:"lat,omitempty"`
}
