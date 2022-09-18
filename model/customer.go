package model

type Customer struct {
	Material    Material `json:"material"`  // available: wood, carpet, tiles
	Address     Coord    `json:"address"`   // Earth coordinates
	FloorArea   float64  `json:"floorArea"` // unit: meters
	PhoneNumber string   `json:"phoneNumber"`
} // @name Customer

func (c Customer) Valid() (string, bool) {

	if c.Material == Undefined {
		return "material not available", false
	}

	if !c.Address.Valid() {
		return "invalid address", false
	}

	return "", true
}
