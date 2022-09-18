package model

type Partner struct {
	NameId          string     `json:"nameId"`
	Materials       []Material `json:"materials"`
	Address         Coord      `json:"address"`
	OperatingRadius float64    `json:"operatingRadius"` // unit: kilometers
	Rating          float64    `json:"rating"`          // range: [0.0,5]
} // @name Partner
