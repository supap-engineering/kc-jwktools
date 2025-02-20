package jwktools

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

// JWK represents the JSON Web Key structure
type JWK struct {
    Keys []Key `json:"keys"`
}

// Key represents a single JWK key
type Key struct {
    Kid string   `json:"kid"`
    Kty string   `json:"kty"`
    Alg string   `json:"alg"`
    Use string   `json:"use"`
    N   string   `json:"n"`
    E   string   `json:"e"`
    X5c []string `json:"x5c"`
}

// KeyFetcher handles JWK key fetching operations
type KeyFetcher struct {
    config *Config
    client *http.Client
}

// NewKeyFetcher creates a new KeyFetcher instance
func NewKeyFetcher(config *Config) *KeyFetcher {
    return &KeyFetcher{
        config: config,
        client: &http.Client{},
    }
}

// GetPublicKey fetches and formats the public key from the JWK endpoint
func (kf *KeyFetcher) GetPublicKey() (string, error) {
    jwk, err := kf.fetchJWK()
    if err != nil {
        return "", fmt.Errorf("failed to fetch JWK: %w", err)
    }

    return kf.formatPublicKey(jwk)
}

func (kf *KeyFetcher) fetchJWK() (*JWK, error) {
    resp, err := kf.client.Get(kf.config.JWKUrl)
    if err != nil {
        return nil, err
    }
    defer resp.Body.Close()

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return nil, err
    }

    var jwk JWK
    if err := json.Unmarshal(body, &jwk); err != nil {
        return nil, err
    }

    return &jwk, nil
}

func (kf *KeyFetcher) formatPublicKey(jwk *JWK) (string, error) {
    for _, key := range jwk.Keys {
        if key.Alg == "RS256" && key.Use == "sig" {
            decodedN, err := base64.RawURLEncoding.DecodeString(key.N)
            if err != nil {
                return "", fmt.Errorf("failed to decode key: %w", err)
            }

            standardBase64 := strings.TrimRight(base64.StdEncoding.EncodeToString(decodedN), "=")
            return fmt.Sprintf("MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA%sIDAQAB", standardBase64), nil
        }
    }

    return "", fmt.Errorf("no RS256 signing key found")
}
