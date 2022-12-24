package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Instance struct {
	ObjId       primitive.ObjectID `bson:"_id,omitempty"`
	Id 			 int32								`bson:"id"`
	Week 		int32 								`bson:"week"`
	Center_id int32								`bson:"center_id"`
	Meal_id int32									`bson:"meal_id"`
	Checkout_price float64			`bson:"checkout_price"`
	Base_price float64					`bson:"base_price"`
	Emailer_for_promotion int32		`bson:"emailer_for_promotion"`
	Homepage_featured int32				`bson:"homepage_featured"`
	Num_orders int32							`bson:"num_orders"`
}