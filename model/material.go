package model

import (
	"bytes"
	"encoding/json"
)

// enum: wood
type Material uint8

const (
	Undefined Material = iota
	Wood
	Carpet
	Tiles
)

var (
	materialFromString = map[string]Material{
		"undefined": Undefined,
		"wood":      Wood,
		"carpet":    Carpet,
		"tiles":     Tiles,
	}

	materialToString = map[Material]string{
		Undefined: "undefined",
		Wood:      "wood",
		Carpet:    "carpet",
		Tiles:     "tiles",
	}
)

// MarshalJSON marshals the enum as a quoted json string
func (m Material) MarshalJSON() ([]byte, error) {
	buffer := bytes.NewBufferString(`"`)
	buffer.WriteString(materialToString[m])
	buffer.WriteString(`"`)
	return buffer.Bytes(), nil
}

// UnmarshalJSON unmashals a quoted json string to the enum value
func (m *Material) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}

	*m = materialFromString[j]
	return nil
}
