package config

import (
	"encoding/json"
	"fmt"
	"os"

	v1 "github.com/nokamoto/analyticapr-go/pkg/api/v1"
	"google.golang.org/protobuf/encoding/protojson"
	"gopkg.in/yaml.v3"
)

// NewConfig returns a new config from a YAML file.
func NewConfig(file string) (*v1.Config, error) {
	bs, err := os.ReadFile(file)
	if err != nil {
		return nil, fmt.Errorf("read file: %w", err)
	}
	var y map[string]any
	if err := yaml.Unmarshal(bs, &y); err != nil {
		return nil, fmt.Errorf("unmarshal yaml: %w", err)
	}
	bs, err = json.Marshal(&y)
	if err != nil {
		return nil, fmt.Errorf("marshal json: %w", err)
	}
	var c v1.Config
	if err := protojson.Unmarshal(bs, &c); err != nil {
		return nil, fmt.Errorf("unmarshal proto: %w", err)
	}
	return &c, nil
}
