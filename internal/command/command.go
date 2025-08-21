package command

import "encoding/json"

// Command Generic action received by the client
type Command struct {
	Type   string  `json:"type"`
	Dx     float64 `json:"dx,omitempty"`
	Dy     float64 `json:"dy,omitempty"`
	Button string  `json:"button,omitempty"`
}

// Parse receives the raw data and converts them to a Command
func Parse(data []byte) (*Command, error) {
	var cmd Command

	// Parse json and fill out cmd struct
	err := json.Unmarshal(data, &cmd)
	if err != nil {
		return nil, err
	}

	// Return pointer for command e no error
	return &cmd, nil
}
