package structures

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Username    string             `json:"username" bson:"username,omitempty"`
	Password    string             `json:"password" bson:"password,omitempty"`
	Authorities string             `json:"authorizes" bson:"authorizes,omitempty"`
}

type LoginForm struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
