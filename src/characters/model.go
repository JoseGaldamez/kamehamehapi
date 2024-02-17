package characters

import "go.mongodb.org/mongo-driver/bson/primitive"

type Character struct {
	ID         primitive.ObjectID `json:"id" bson:"_id"`
	Name       string             `json:"name" bson:"name"`
	Race       string             `json:"race" bson:"race"`
	Age        int                `json:"age" bson:"age"`
	Universe   int                `json:"universe" bson:"universe"`
	Techniques []string           `json:"techniques" bson:"techniques"`
	Transforms []string           `json:"transforms" bson:"transforms"`
	Image      string             `json:"image" bson:"image"`
}
