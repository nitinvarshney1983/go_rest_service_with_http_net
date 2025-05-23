package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"rest_services_with_http_net/models"
	"rest_services_with_http_net/service"
	"testing"
)

// Mock service for testing
type mockEmployeeService struct {
	InsertEmployeeFunc  func(models.Employee) (models.Employee, error)
	GetEmployeeByIDFunc func(string) (models.Employee, error)
	GetAllEmployeesFunc func() ([]models.Employee, error)
	UpdateEmployeeFunc  func(string, models.Employee) (models.Employee, error)
}

var _ service.EmployeeServiceInf = (*mockEmployeeService)(nil)

func (m *mockEmployeeService) InsertEmployee(emp models.Employee) (models.Employee, error) {
	return m.InsertEmployeeFunc(emp)
}
func (m *mockEmployeeService) GetEmployeeByID(id string) (models.Employee, error) {
	return m.GetEmployeeByIDFunc(id)
}
func (m *mockEmployeeService) GetAllEmployees() ([]models.Employee, error) {
	return m.GetAllEmployeesFunc()
}
func (m *mockEmployeeService) UpdateEmployee(id string, emp models.Employee) (models.Employee, error) {
	return m.UpdateEmployeeFunc(id, emp)
}

func TestGetByID_Success(t *testing.T) {
	mockService := &mockEmployeeService{
		GetEmployeeByIDFunc: func(id string) (models.Employee, error) {
			return models.Employee{FirstName: "Test", LastName: "Test", Email: "test@example.com", Phone: "99999999"}, nil
		},
	}
	handler := NewEmployeeHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/employees/test@example.com", nil)
	rr := httptest.NewRecorder()

	handler.GetByID(rr, req)

	if rr.Code != http.StatusOK {
		t.Errorf("expected status 200, got %d", rr.Code)
	}
	var emp models.Employee
	if err := json.NewDecoder(rr.Body).Decode(&emp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if emp.Email != "test@example.com" {
		t.Errorf("expected email 'test@example.com', got '%s'", emp.Email)
	}
}

func TestGetByID_NotFound(t *testing.T) {
	mockService := &mockEmployeeService{
		GetEmployeeByIDFunc: func(id string) (models.Employee, error) {
			return models.Employee{}, errors.New("not found")
		},
	}
	handler := NewEmployeeHandler(mockService)

	req := httptest.NewRequest(http.MethodGet, "/employees/unknown@example.com", nil)
	rr := httptest.NewRecorder()

	handler.GetByID(rr, req)

	if rr.Code != http.StatusNotFound {
		t.Errorf("expected status 404, got %d", rr.Code)
	}
}

func TestCreate_Success(t *testing.T) {
	mockService := &mockEmployeeService{
		InsertEmployeeFunc: func(emp models.Employee) (models.Employee, error) {
			return emp, nil
		},
	}
	handler := NewEmployeeHandler(mockService)

	emp := models.Employee{FirstName: "Test", LastName: "Test", Email: "test@example.com", Phone: "99999999", Position: "Developer", Salary: 50000, Address: "123 Test St", CreatedAt: "2023-10-01", UpdatedAt: "2023-10-01"}
	body, _ := json.Marshal(emp)
	req := httptest.NewRequest(http.MethodPost, "/employees", bytes.NewReader(body))
	rr := httptest.NewRecorder()

	handler.Create(rr, req)

	if rr.Code != http.StatusCreated {
		t.Errorf("expected status 201, got %d", rr.Code)
	}
	var resp map[string]string
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if resp["message"] != "Employee created" {
		t.Errorf("expected message 'Employee created', got '%s'", resp["message"])
	}
}
