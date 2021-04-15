package main

import (
	"encoding/json"
	"fmt"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/rezamt/asf-server/cmd/auth"
	"github.com/rezamt/asf-server/pkg/efood/efood/models"
	"github.com/rezamt/asf-server/pkg/efood/efood/restapi"
	"github.com/rezamt/asf-server/pkg/efood/efood/restapi/operations"
	"github.com/rezamt/asf-server/pkg/efood/efood/restapi/operations/user"
	"log"
)

func main() {

	// Initialize Swagger
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewEfoodAPI(swaggerSpec)

	api.OauthSecurityAuth = verifyOauthAccessToken

	// Implement Other methods
	api.UserGetCartHandler = user.GetCartHandlerFunc(Cards)

	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.Port = 8080

	// Start listening using having the handlers and port
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}

func verifyOauthAccessToken(token string, scopes []string) (interface{}, error) {
	fmt.Printf("Validating Access Token Passed Through ")

	headers, claims, err := auth.VerifyAccessToken(token, "jupitercm.auth0.com")

	if err != nil {
		return nil, errors.New(401, "error authenticate")
	}

	printJSON(headers)
	printJSON(claims)

	prin := models.Principal(token)
	return &prin, nil
}

func Cards(params user.GetCartParams, i interface{}) middleware.Responder {
	cp := models.CartPreview{}
	cp = append(cp, &models.CartItem{
		Currency:    "AUD",
		ImageURL:    "https://card.img",
		ProductID:   1000,
		ProductName: "FirstProd",
		Quantity:    0,
		UnitPrice:   0,
	})

	return user.NewGetCartOK().WithPayload(cp)
}

func printJSON(j interface{}) error {
	var out []byte
	var err error

	out, err = json.Marshal(j)

	if err == nil {
		fmt.Println(string(out))
	}

	return err
}
