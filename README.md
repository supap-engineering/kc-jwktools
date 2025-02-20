![Supap Logo](https://supap.co.uk/images/logo.svg)

# Keycloak Public Key Retriever

A Go package for fetching and formatting public keys from Keycloak JWK endpoints. This package helps in retrieving RSA256 public keys from Keycloak realms for JWT validation.

## Installation

```bash
go get github.com/supap-engineering/kc-jwktools
```

## Configuration

The package accepts configuration in the following format:

```go
type Config struct {
    Type   string `json:"type"`   // Currently supports "RS256"
    JWKUrl string `json:"jwk_url"` // Full URL to Keycloak's JWK endpoint
}
```

### Basic Usage

```go
import "github.com/supap-engineering/kc-jwktools/jwktools"

func main() {
    // Configure using JSON string
    configJSON := `{
        "type": "RS256",
        "jwk_url": "https://your-keycloak-domain.com/realms/your-realm/protocol/openid-connect/certs"
    }`
    
    config, err := pkg.NewConfigFromJSON(configJSON)
    if err != nil {
        log.Fatal(err)
    }

    keyFetcher := pkg.NewKeyFetcher(config)
    publicKey, err := keyFetcher.GetPublicKey()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Public Key: %s\n", publicKey)
}
```

### Using with Environment Variables (Viper)

```go
import (
    "github.com/supap-engineering/kc-jwktools/jwktools"
    "github.com/spf13/viper"
)

func main() {
    // Load config from environment variable
    keycloakConfig := viper.GetString("KEYCLOAK_REALMRS256PUBLICKEY")
    
    config, err := pkg.NewConfigFromJSON(keycloakConfig)
    if err != nil {
        log.Fatal(err)
    }

    keyFetcher := pkg.NewKeyFetcher(config)
    publicKey, err := keyFetcher.GetPublicKey()
    if err != nil {
        log.Fatal(err)
    }
}
```

### Environment Variable Format

```env
KEYCLOAK_REALMRS256PUBLICKEY={"type":"RS256","jwk_url":"https://your-domain.com/realms/your-realm/protocol/openid-connect/certs"}
```

### Keycloak URL Format

The JWK URL should follow this pattern:
```
https://<keycloak-domain>/realms/<realm-name>/protocol/openid-connect/certs
```

## Features

- Fetches public keys from Keycloak JWK endpoints
- Supports RS256 signing algorithm
- Formats public keys for JWT validation
- Environment variable configuration support
- Error handling and validation

## Maintainer

This project is maintained by [Supap](https://supap.co.uk)

## Contributing

Your contributions are welcome! Please [create a pull request](https://github.com/supap-engineering/kc-jwktools/pulls) on GitHub. Much larger changes need to be discussed with the development team via the [issues section at GitHub](https://github.com/supap-engineering/kc-jwktools/issues).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.


---
⭐ If you find this package helpful, please consider giving it a star