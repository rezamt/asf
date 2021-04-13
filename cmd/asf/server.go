package main

import (
	"github.com/go-openapi/runtime/middleware"
	"log"

	"github.com/go-openapi/loads"
	"github.com/rezamt/asf-server/pkg/swagger/server/restapi"
	"github.com/rezamt/asf-server/pkg/swagger/server/restapi/operations"
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

	// Implement the CheckHealth handler
	//api.CheckHealthHandler = operations.CheckHealthHandlerFunc(
	//	func(user operations.CheckHealthParams) middleware.Responder {
	//		return operations.NewCheckHealthOK().WithPayload("OK")
	//	})

	// Implement the GetHelloUser handler
	//api.GetHelloUserHandler = operations.GetHelloUserHandlerFunc(
	//	func(user operations.GetHelloUserParams) middleware.Responder {
	//		return operations.NewGetHelloUserOK().WithPayload("Hello " + user.User + "!")
	//	})

	// Start listening using having the handlers and port
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}

//Health route returns OK
func Health(param operations.CheckHealthParams) middleware.Responder {
	return operations.NewCheckHealthOK().WithPayload("OK")
}


//GetHelloUser returns Hello + your name
func GetHelloUser(user operations.GetHelloUserParams) middleware.Responder {
	return operations.NewGetHelloUserOK().WithPayload("Hello " + user.User + "!")
}

