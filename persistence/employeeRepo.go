package persistence

import (
	"context"
	"fmt"
	"rest_services_with_http_net/models"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type EmployeeRepo struct {
	coll *mongo.Collection
}

var (
	employeeRepoInstance *EmployeeRepo
	employeeRepoOnce     sync.Once
)

func GetEmployeeRepo() *EmployeeRepo {
	employeeRepoOnce.Do(func() {
		employeeRepoInstance = &EmployeeRepo{
			coll: GetCollection("employees"),
		}
	})
	return employeeRepoInstance
}

func (er *EmployeeRepo) InsertEmployee(employee models.Employee) (models.Employee, error) {
	employee.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	result, err := er.coll.InsertOne(context.Background(), employee)
	if err != nil {
		return models.Employee{}, err
	}
	if oid, ok := result.InsertedID.(primitive.ObjectID); ok {
		employee.ID = oid
	} else {
		return models.Employee{}, fmt.Errorf("failed to convert InsertedID to ObjectID")
	}
	return employee, nil
}

func (er *EmployeeRepo) GetEmployeeByID(id string) (models.Employee, error) {
	var employee models.Employee
	err := er.coll.FindOne(context.Background(), bson.M{"email": id}).Decode(&employee)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return models.Employee{}, fmt.Errorf("employee not found")
		}
		return models.Employee{}, err
	}
	return employee, nil
}

func (er *EmployeeRepo) GetAllEmployees() ([]models.Employee, error) {
	var employees []models.Employee
	cursor, err := er.coll.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var employee models.Employee
		if err := cursor.Decode(&employee); err != nil {
			return nil, err
		}
		employees = append(employees, employee)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return employees, nil
}

func (er *EmployeeRepo) UpdateEmployee(id string, updated models.Employee) (models.Employee, error) {
	var updatedEmployee models.Employee
	filter := bson.M{"email": id}
	update := bson.M{
		"$set": bson.M{
			"firstName": updated.FirstName,
			"lastName":  updated.FirstName,
			"phone":     updated.Phone,
			"position":  updated.Position,
			"salary":    updated.Salary,
			"address":   updated.Address,
			"updatedAt": time.Now().Format("2006-01-02 15:04:05"),
		},
	}
	result, err := er.coll.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return updatedEmployee, err
	}
	if result.MatchedCount == 0 {
		return updatedEmployee, fmt.Errorf("no employee found with id: %s", id)
	}

	return result.UpsertedID.(models.Employee), nil
}
