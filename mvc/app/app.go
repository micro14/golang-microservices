package app

import (
	"net/http"

	"github.com/micro14/golang-microservices/mvc/controllers"
)

func StartApp() {
	// Initializing the handlers that we have
	http.HandleFunc("/users", controllers.GetUser)

	// Launch the Application
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
