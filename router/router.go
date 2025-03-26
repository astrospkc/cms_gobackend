package router

import (
	"gobackend/controller"

	"github.com/gorilla/mux"
)

func Router()*mux.Router{
	r:= mux.NewRouter()
	// user router
	r.HandleFunc("api/createUser",controller.CreateUser).Methods("POST")
	// r.HandleFunc("/tea", ).Methods("GET")
	
	return r
}

