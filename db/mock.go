package db

import (
	"github.com/nuareal/matching-partner/model"
)

var Partners = map[string]model.Partner{
	"Pecker":   {NameId: "Pecker", Materials: []model.Material{model.Wood}, Address: model.Coord{Lat: 52.10, Lon: 13.37}, OperatingRadius: 50, Rating: 3.9},
	"Comfy":    {NameId: "Comfy", Materials: []model.Material{model.Carpet}, Address: model.Coord{Lat: 52.50, Lon: 13.80}, OperatingRadius: 80, Rating: 4.5},
	"Pixels":   {NameId: "Pixels", Materials: []model.Material{model.Tiles}, Address: model.Coord{Lat: 52.25, Lon: 13.37}, OperatingRadius: 60, Rating: 4.4},
	"Squared":  {NameId: "Squared", Materials: []model.Material{model.Wood, model.Tiles}, Address: model.Coord{Lat: 52.90, Lon: 13.37}, OperatingRadius: 150, Rating: 3.2},
	"AllinOne": {NameId: "AllinOne", Materials: []model.Material{model.Wood, model.Carpet, model.Tiles}, Address: model.Coord{Lat: 52.50, Lon: 13.37}, OperatingRadius: 50, Rating: 4.5},
}
