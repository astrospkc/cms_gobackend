// Connects to MongoDB and sets a Stable API version
package main

import (
	"fmt"
	"gobackend/connect"
	"gobackend/controller"

	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	
	connect.Connect()
	
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
	// r := router.Router()
	r:= mux.NewRouter()
	// fmt.Println("the router : ", r)
	// r.Use(handlers.CORS(
	// 	handlers.AllowedOrigins([]string{"*"}),
	// 	handlers.AllowedHeaders([]string{"X-Requested-With","Content-Type","Authorization"}),
	// 	handlers.AllowedMethods([]string{"GET", "POST","DELETE","PUT", "HEAD","OPTIONS"}),
	// ))
	r.HandleFunc("/api/createuser", controller.CreateUser).Methods("POST")
	
	
		err:=http.ListenAndServe(":8000", r)
		if err!= nil{
			log.Fatal("there is an error setting up: ", err)
		}
	// established connection
	// srv:=&http.Server{
	// 	Addr: "127.0.0.1:8000",
	// 	WriteTimeout: time.Second*15,
	// 	ReadTimeout: time.Second*15,
	// 	IdleTimeout: time.Second*60,
	// 	Handler:r,
	// }
	// fmt.Println("Server is up and running")
	// if err:=srv.ListenAndServe();err!=nil{
	// 	log.Println(err)
	// }
	
}
