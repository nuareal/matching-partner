package model

import (
	"math"
)

const earthRaidusKm = 6371

// Coord
type Coord struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
} // @name Coord

// Distance calculates beeline between two coordinates in kilometers.
func Distance(c1 Coord, c2 Coord) (km float64) {
	lat1 := degreesToRadians(c1.Lat)
	lon1 := degreesToRadians(c1.Lon)
	lat2 := degreesToRadians(c2.Lat)
	lon2 := degreesToRadians(c2.Lon)

	diffLat := lat2 - lat1
	diffLon := lon2 - lon1

	distCalc := math.Pow(math.Sin(diffLat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*
		math.Pow(math.Sin(diffLon/2), 2)

	distCalc = 2 * math.Atan2(math.Sqrt(distCalc), math.Sqrt(1-distCalc))

	distance := distCalc * earthRaidusKm

	return distance
}

// Valid checks if coordinate is valid.
func (c Coord) Valid() bool {
	return c.Lat > -90 && c.Lat < 90 && c.Lon > -180 && c.Lon < 180
}

// degreesToRadians converts from degrees to radians.
func degreesToRadians(d float64) float64 {
	return d * math.Pi / 180
}
