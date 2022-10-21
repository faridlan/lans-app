package database

import (
	"context"
	"testing"

	"github.com/faridlan/lans-app/app"
	"go.mongodb.org/mongo-driver/bson"
)

func TestConnection(t *testing.T) {
	c := app.NewDatabase()
	Collection := c.Database("rekap").Collection("rekap")
	Collection.Find(context.Background(), bson.M{})
}
