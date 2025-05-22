package handlers

import (
	"encoding/json"
	"net/http"
	"rest_services_with_http_net/models"
	"rest_services_with_http_net/service"
	"strings"
)

type EmployeeHandler struct {
	employeeService *service.EmployeeService
}

func NewEmployeeHandler(empService *service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{
		employeeService: empService,
	}
}

func (eh *EmployeeHandler) Create(w http.ResponseWriter, r *http.Request) {
	// Call the service layer to insert the employee
	var emp models.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if _, err := eh.employeeService.InsertEmployee(emp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Employee created"})
}
func (eh *EmployeeHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	// Call the service layer to get the employee by ID
	path := r.URL.Path
	segments := splitPath(path)
	// Check if the path has the correct structure
	if len(segments) < 2 || segments[0] != "employees" {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}

	id := segments[1] // Extract the ID from the second segment
	if id == "" {
		http.Error(w, "emailId is required", http.StatusBadRequest)
		return
	}

	emp, err := eh.employeeService.GetEmployeeByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(emp)
}
func (eh *EmployeeHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	// Call the service layer to get all employees
	employees, err := eh.employeeService.GetAllEmployees()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(employees)
}
func (eh *EmployeeHandler) Update(w http.ResponseWriter, r *http.Request) {
	// Call the service layer to update the employee
	path := r.URL.Path
	segments := splitPath(path)
	// Check if the path has the correct structure
	if len(segments) < 2 || segments[0] != "employees" {
		http.Error(w, "Invalid path", http.StatusBadRequest)
		return
	}

	email := segments[1] // Extract the ID from the second segment
	if email == "" {
		http.Error(w, "emailId is required", http.StatusBadRequest)
		return
	}

	var emp models.Employee
	if err := json.NewDecoder(r.Body).Decode(&emp); err != nil {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	if _, err := eh.employeeService.UpdateEmployee(email, emp); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{"message": "Employee updated"})
}

// Helper function to split the path into segments
func splitPath(path string) []string {
	segments := []string{}
	for _, segment := range strings.Split(path, "/") {
		if segment != "" {
			segments = append(segments, segment)
		}
	}
	return segments
}
