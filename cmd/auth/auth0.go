package auth

import (
	"encoding/json"
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"net/http"
)

type Jwks struct {
	Keys []JSONWebKeys `json:"keys"`
}

type JSONWebKeys struct {
	Kty string   `json:"kty"`
	Kid string   `json:"kid"`
	Use string   `json:"use"`
	N   string   `json:"n"`
	E   string   `json:"e"`
	X5c []string `json:"x5c"`
}

func VerifyAccessToken(accessToken string, domain string) (headers map[string]interface{}, claims jwt.Claims, err error) {

	// Parse Token but Don't Validate it for now
	parser := jwt.Parser{}
	parser.SkipClaimsValidation = false
	token, _, _ := parser.ParseUnverified(accessToken, jwt.StandardClaims{})

	// Getting RSA Public Key for Auth0 Access Token
	pubKey, _ := getTokenPublicKey(token, domain)

	// Parse the token.  Load the key from command line option
	token, err = parser.ParseWithClaims(accessToken, jwt.MapClaims{}, func(t *jwt.Token) (interface{}, error) {
		return jwt.ParseRSAPublicKeyFromPEM([]byte(pubKey))
	})

	// if there is any error during parsing token
	if err != nil {
		return nil, nil, err
	}

	return token.Header, token.Claims, nil
}

func getTokenPublicKey(token *jwt.Token, domain string) (string, error) {
	cert := ""
	resp, err := http.Get("https://" + domain + "/.well-known/jwks.json")

	if err != nil {
		return cert, err
	}
	defer resp.Body.Close()

	var jwks = Jwks{}
	err = json.NewDecoder(resp.Body).Decode(&jwks)

	if err != nil {
		return cert, err
	}

	for k, _ := range jwks.Keys {
		if token.Header["kid"] == jwks.Keys[k].Kid {
			cert = "-----BEGIN CERTIFICATE-----\n" + jwks.Keys[k].X5c[0] + "\n-----END CERTIFICATE-----"
		}
	}

	if cert == "" {
		err := errors.New("Unable to find appropriate key.")
		return cert, err
	}

	return cert, nil
}
