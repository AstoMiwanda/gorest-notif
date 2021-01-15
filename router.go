package main

import (
	"net/http"

	"github.com/pushm0v/gorest-notif/client"

	"github.com/gorilla/mux"
	"github.com/pushm0v/gorest-notif/service"
)

func RestRouter() *mux.Router {
	r := mux.NewRouter()
	api := r.PathPrefix("/api/v1").Subrouter()

	notifRouter(api)
	r.Use(LoggingMiddleware)
	return r
}

func notifRouter(r *mux.Router) {

	var sendgridClient = client.NewSendgridClient()
	var notifService = service.NewNotifService(sendgridClient)
	var handler = NewNotifHandler(notifService)

	r.HandleFunc("/email", handler.SendEmailNotif).Methods(http.MethodPost)
	r.HandleFunc("/", handler.NotFound)
}
