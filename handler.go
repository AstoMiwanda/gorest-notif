package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/pushm0v/gorest-notif/model"
	"github.com/pushm0v/gorest-notif/service"
)

type NotifHandler struct {
	notifService service.NotifService
}

func NewNotifHandler(notifService service.NotifService) *NotifHandler {
	return &NotifHandler{notifService: notifService}
}

func (s *NotifHandler) responseBuilder(w http.ResponseWriter, message interface{}) {
	w.Header().Set("Content-Type", "application/json")
	m := model.Response{
		Message: message,
	}

	err := json.NewEncoder(w).Encode(m)
	if err != nil {
		log.Fatalf("Response builder error : %v", err)
	}
}

func (s *NotifHandler) SendEmailNotif(w http.ResponseWriter, r *http.Request) {

	var m = &model.Message{}
	err := json.NewDecoder(r.Body).Decode(m)
	if err != nil {
		errMsg := fmt.Sprintf("Request decoder error : %v", err)

		w.WriteHeader(http.StatusBadRequest)
		s.responseBuilder(w, errMsg)
		return
	}

	if m.Destination == "" || m.DestinationName == "" {
		errMsg := fmt.Sprintf("Destination or destination name is empty")

		w.WriteHeader(http.StatusBadRequest)
		s.responseBuilder(w, errMsg)
		return
	}

	m.Type = model.Email
	m.From = "no-reply@kitabisa.com"
	m.FromName = "No Reply"

	err = s.notifService.SendEmail(m)
	if err != nil {
		errMsg := fmt.Sprintf("Send Email Notif error : %v", err)

		w.WriteHeader(http.StatusInternalServerError)
		s.responseBuilder(w, errMsg)
		return
	}
	w.WriteHeader(http.StatusCreated)
	s.responseBuilder(w, "email sent")
}

func (s *NotifHandler) NotFound(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)
	s.responseBuilder(w, "not found")
}
