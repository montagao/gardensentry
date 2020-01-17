package main

import (
	"flag"
	"log"

	"gardensentry.v1/gen/models"
	"gardensentry.v1/gen/restapi"
	"gardensentry.v1/gen/restapi/operations"
	"gardensentry.v1/internal/store"
	_ "github.com/GoogleCloudPlatform/cloudsql-proxy/proxy/dialers/postgres"
	"github.com/go-openapi/loads"
	"github.com/go-openapi/runtime/middleware"
	_ "github.com/lib/pq"
)

func main() {
	eventStore, err := store.New()

	//ctx := context.Background()
	portFlag := flag.Int("port", 3000, "Port to run this service on")

	// load embedded swagger file
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// create new service API
	api := operations.NewGardensentryAPI(swaggerSpec)
	server := restapi.NewServer(api)
	defer server.Shutdown()

	// parse flags
	flag.Parse()
	// set the port this service will be run on
	server.Port = *portFlag

	api.AddEventHandler = operations.AddEventHandlerFunc(
		func(params operations.AddEventParams) middleware.Responder {
			newEvent := &models.Event{
				Description: params.Body.Description,
				Timestamp:   params.Body.Timestamp,
				Type:        params.Body.Type,
				VidURL:      params.Body.VidURL,
			}
			err := eventStore.Put(newEvent)
			// TODO: Return errors
			if err != nil {
				log.Fatal(err)
			}
			return operations.NewAddEventCreated().WithPayload(newEvent)
		})

	api.GetEventsHandler = operations.GetEventsHandlerFunc(
		func(params operations.GetEventsParams) middleware.Responder {
			events, err := eventStore.GetAll(int(*params.Limit))
			if err != nil {
				log.Fatal(err)
			}
			return operations.NewGetEventsOK().WithPayload(events)
		})

	api.GetEventByIDHandler = operations.GetEventByIDHandlerFunc(
		func(params operations.GetEventByIDParams) middleware.Responder {
			event, err := eventStore.GetByID(params.ID)
			if err != nil {
				log.Fatal(err)
			}

			if event != nil {
				return operations.NewGetEventByIDOK().WithPayload(event)
			} else {
				return operations.NewGetEventByIDDefault(404)
			}
		})

	/* should handle video uploading directly from Pi. Then make vidUrl a required field for event POSTs
	api.UploadVideoToEventHandler = operations.UploadVideoToEventHandlerFunc(
		func(params operations.UploadVideoToEventParams) middleware.Responder {
			var event *models.Event
			for _, e := range events {
				if e.ID == params.ID {
					event = e
				}
			}
			if event == nil {
				// no associated event found
				return operations.NewGetEventByIDDefault(404)
			}

			buf := []byte{}
			_, err := params.Upfile.Read(buf)
			if err != nil {
				log.Fatal("could not read file")
			}

			log.Printf("blah")
			log.Printf("got bytes: %s", string(buf))
			// upload vid to Google Cloud then update the URL in DB
			event.VidURL = fmt.Sprintf("%d-vidURL", params.ID)
			return operations.NewGetEventByIDOK().WithPayload(event)
		})
	*/

	// serve API
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}
