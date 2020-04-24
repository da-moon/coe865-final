package config

import (
	"encoding/json"
	"io"

	"github.com/mitchellh/mapstructure"
)

// DecodeJSONConfig ...
func DecodeJSONConfig(r io.Reader) (*Config, error) {

	var raw interface{}
	dec := json.NewDecoder(r)
	if err := dec.Decode(&raw); err != nil {
		return nil, err
	}
	var md mapstructure.Metadata
	var result Config
	msdec, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Metadata:    &md,
		Result:      &result,
		ErrorUnused: true,
	})
	if err != nil {
		return nil, err
	}
	if err := msdec.Decode(raw); err != nil {
		return nil, err
	}
	return &result, nil
}
