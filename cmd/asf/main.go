package main

import (
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	"github.com/rezamt/asf-server/pkg/asf/asf/restapi"
	"github.com/rezamt/asf-server/pkg/asf/asf/restapi/operations"
	"log"
)

func main() {

	// Initialize Swagger
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewHelloAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	server.Port = 8080

	api.GetHelloUserHandler = operations.GetHelloUserHandlerFunc(GetHelloUser)
	api.CheckHealthHandler = operations.CheckHealthHandlerFunc(Health)

	// Start listening using having the handlers and port
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}

// Validating Access Token
func isAuthorized(authz string) bool {
	if authz == "" {
		return false
	}
	log.Printf("Authorization: %s", authz)
	return true
}

//Health route returns OK
func Health(param operations.CheckHealthParams) middleware.Responder {

	if !isAuthorized(param.HTTPRequest.Header.Get("Authorization")) {
		return operations.NewGetHelloUserBadRequest()
	}

	return operations.NewCheckHealthOK().WithPayload("OK")
}

//GetHelloUser returns Hello + your name
func GetHelloUser(param operations.GetHelloUserParams) middleware.Responder {
	if !isAuthorized(param.HTTPRequest.Header.Get("Authorization")) {
		return operations.NewGetHelloUserBadRequest()
	}

	return operations.NewGetHelloUserOK().WithPayload("Hello " + param.User + "!")
}
