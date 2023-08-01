package db

import (
	"context"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"testing"
)

func TestInitMongoDB(t *testing.T) {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	var err error
	Client, err = mongo.Connect(context.Background(), clientOptions)
	assert.Nil(t, err, "Client 획득 실패")
	Collection = Client.Database("local").Collection("jwtPrac")
	assert.Nil(t, err, "Collecton 획득 실패")
}
