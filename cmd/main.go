package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"go.mongodb.org/mongo-driver/bson"
	 contact "contact.com"
  "github.com/go-chi/chi"
)

var mh *contact.MongoHandler

func registerRoutes() http.Handler {
	r := chi.NewRouter()
	r.Route("/contacts", func(r chi.Router) {
		r.Get("/", getAllContact)                 //GET /contacts
		r.Get("/{phonenumber}", getContact)       //GET /contacts/0147344454
		r.Post("/", addContact)                   //POST /contacts
		r.Put("/{phonenumber}", updateContact)    //PUT /contacts/0147344454
		r.Delete("/{phonenumber}", deleteContact) //DELETE /contacts/0147344454
	})
	return r
}


func main(){
	mongoDbConnection := "mongodb://localhost:27017"
	mh = contact.NewHandler(mongoDbConnection) //Create an instance of MongoHander with the connection string provided
	r := registerRoutes()
	log.Fatal(http.ListenAndServe(":3060", r)) //You can modify to run on a different port
}
