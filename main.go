package main

import (
	"log"
	"net/http"
	"rest_services_with_http_net/configs"
	"rest_services_with_http_net/handlers"
	"rest_services_with_http_net/persistence"
	"rest_services_with_http_net/service"
)

func main() {

	// Load configuration
	configs.LoadConfig()

	// Initialize MongoDB client
	persistence.InitMongoClient()

	// Initialize Employee repository
	empRepo := persistence.GetEmployeeRepo()

	// Initialize Employee service
	empService := service.NewEmployeeService(empRepo)

	// Initialize HTTP server and handlers
	empHandler := handlers.NewEmployeeHandler(empService)
	http.HandleFunc("/employees", empHandler.Create)
	http.HandleFunc("/employees/{email}", empHandler.GetByID)
	http.HandleFunc("/employees/all", empHandler.GetAll)
	http.HandleFunc("/employees/update/{id}", empHandler.Update)

	// Start the server
	log.Fatal(http.ListenAndServe(":"+configs.AppConfig.Port, nil))
}
