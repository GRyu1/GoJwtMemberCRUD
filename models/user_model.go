package models

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"goLangJwtPrac/db"
	"goLangJwtPrac/structures"
	"goLangJwtPrac/utils"
)

func HandlePostInsertUser(user *structures.User) error {
	user.Authorities = "USER"
	_, err := db.Collection.InsertOne(context.Background(), user)
	return err
}
func HandleGetAllUser() ([]structures.User, error) {
	var users []structures.User
	cursor, err := db.Collection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		var user structures.User
		if err = cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err = cursor.Err(); err != nil {
		return nil, err
	}
	return users, err
}

func HandleGetUser(idStr string) (structures.User, error) {
	var user structures.User
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return user, err
	}
	filter := bson.M{"_id": id}
	err = db.Collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return user, err
	}
	return user, err
}
func HandleDeleteUser(idStr string) error {
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return err
	}
	_, err = db.Collection.DeleteOne(context.Background(), bson.D{{"_id", id}})
	if err != nil {
		return err
	}
	return nil
}
func HandlePatchUser(idStr string, user *structures.User) (structures.User, error) {
	var updatedUser structures.User
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return updatedUser, err
	}
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"username": user.Username, "password": user.Password}}
	result, err := db.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return updatedUser, err
	}
	if result.ModifiedCount > 0 {
		err = db.Collection.FindOne(context.Background(), filter).Decode(&updatedUser)
		if err != nil {
			return updatedUser, err
		}
	}
	return updatedUser, nil
}
func HandleAuthentication(loginRequest *structures.LoginForm) (string, error) {
	var user structures.User
	filter := bson.M{"username": loginRequest.Username}
	err := db.Collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return "", err
	}
	if user.Password != loginRequest.Password {
		return "", fmt.Errorf("로그인 실패")
	}
	accessToken, err := utils.CreateAccessToken(&user)
	if err != nil {
		return "", err
	}
	return accessToken, nil
}
