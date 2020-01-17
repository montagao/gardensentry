// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"net/http"

	errors "github.com/go-openapi/errors"
	runtime "github.com/go-openapi/runtime"
	middleware "github.com/go-openapi/runtime/middleware"

	"gardensentry.v1/gen/restapi/operations"
)

//go:generate swagger generate server --target ../../gen --name Gardensentry --spec ../../api.yml --exclude-main

func configureFlags(api *operations.GardensentryAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.GardensentryAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.JSONConsumer = runtime.JSONConsumer()
	api.JSONProducer = runtime.JSONProducer()

	if api.AddEventHandler == nil {
		api.AddEventHandler = operations.AddEventHandlerFunc(func(params operations.AddEventParams) middleware.Responder {
			return middleware.NotImplemented("operation .AddEvent has not yet been implemented")
		})
	}
	if api.DeleteEventHandler == nil {
		api.DeleteEventHandler = operations.DeleteEventHandlerFunc(func(params operations.DeleteEventParams) middleware.Responder {
			return middleware.NotImplemented("operation .DeleteEvent has not yet been implemented")
		})
	}
	if api.GetEventByIDHandler == nil {
		api.GetEventByIDHandler = operations.GetEventByIDHandlerFunc(func(params operations.GetEventByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation .GetEventByID has not yet been implemented")
		})
	}
	if api.GetEventsHandler == nil {
		api.GetEventsHandler = operations.GetEventsHandlerFunc(func(params operations.GetEventsParams) middleware.Responder {
			return middleware.NotImplemented("operation .GetEvents has not yet been implemented")
		})
	}
	if api.UpdateEventHandler == nil {
		api.UpdateEventHandler = operations.UpdateEventHandlerFunc(func(params operations.UpdateEventParams) middleware.Responder {
			return middleware.NotImplemented("operation .UpdateEvent has not yet been implemented")
		})
	}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix"
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
