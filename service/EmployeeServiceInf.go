package service

import "rest_services_with_http_net/models"

type EmployeeServiceInf interface {
	InsertEmployee(models.Employee) (string, error)
	GetEmployeeByID(string) (models.Employee, error)
	GetAllEmployees() ([]models.Employee, error)
	UpdateEmployee(string, models.Employee) (string, error)
}
