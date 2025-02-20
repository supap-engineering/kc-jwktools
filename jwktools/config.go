package jwktools

import (
	"encoding/json"
	"fmt"
)

// Config represents the JWK configuration
type Config struct {
    Type   string `json:"type"`
    JWKUrl string `json:"jwk_url"`
}

// NewConfigFromJSON creates a new Config from JSON string
func NewConfigFromJSON(jsonStr string) (*Config, error) {
    var config Config
    if err := json.Unmarshal([]byte(jsonStr), &config); err != nil {
        return nil, fmt.Errorf("failed to parse config: %w", err)
    }
    
    if err := config.validate(); err != nil {
        return nil, err
    }
    
    return &config, nil
}

func (c *Config) validate() error {
    if c.Type == "" {
        return fmt.Errorf("type cannot be empty")
    }
    if c.JWKUrl == "" {
        return fmt.Errorf("jwk_url cannot be empty")
    }
    return nil
}
