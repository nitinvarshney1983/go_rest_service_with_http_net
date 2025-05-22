package service

import (
	"rest_services_with_http_net/models"
	"rest_services_with_http_net/persistence"
)

type EmployeeService struct {
	// Add any dependencies or configurations needed for the service
	employeeRepo *persistence.EmployeeRepo
}

func NewEmployeeService(empRepo *persistence.EmployeeRepo) *EmployeeService {
	return &EmployeeService{
		employeeRepo: empRepo,
	}
}
func (es *EmployeeService) InsertEmployee(employee models.Employee) (models.Employee, error) {
	return es.employeeRepo.InsertEmployee(employee)
}

func (es *EmployeeService) GetEmployeeByID(id string) (models.Employee, error) {
	//empId, err := primitive.ObjectIDFromHex(id)
	/*if err != nil {
		return models.Employee{}, err
	}*/
	return es.employeeRepo.GetEmployeeByID(id)
}
func (es *EmployeeService) GetAllEmployees() ([]models.Employee, error) {
	return es.employeeRepo.GetAllEmployees()
}
func (es *EmployeeService) UpdateEmployee(email string, employee models.Employee) (models.Employee, error) {
	return es.employeeRepo.UpdateEmployee(email, employee)
}
