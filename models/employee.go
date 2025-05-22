package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Employee struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	FirstName string             `bson:"firstName" json:"first_name"`
	LastName  string             `bson:"lastName" json:"last_name"`
	Email     string             `bson:"email" json:"email"`
	Phone     string             `bson:"phone" json:"phone"`
	Position  string             `bson:"position" json:"position"`
	Salary    int                `bson:"salary" json:"salary"`
	Address   string             `bson:"address" json:"address"`
	CreatedAt string             `bson:"createdAt" json:"created_at"`
	UpdatedAt string             `bson:"updatedAt" json:"updated_at"`
}
