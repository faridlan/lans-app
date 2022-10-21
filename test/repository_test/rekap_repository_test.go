package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/faridlan/lans-app/app"
	"github.com/faridlan/lans-app/model/domain"
	"github.com/faridlan/lans-app/repository"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestRekapInsert(t *testing.T) {

	client := app.NewDatabase()
	Collection := client.Database("lans_app").Collection("rekap")
	repository := repository.NewRekapRepository(Collection)
	id := primitive.NewObjectID()

	rekap, err := repository.CreateOne(context.Background(), domain.Rekap{
		Id:          id,
		CsName:      "Joko",
		CusName:     "Moro",
		RekapStatus: true,
		PrintStatus: true,
		RekapDate:   time.Now().Unix(),
	})

	assert.Nil(t, err)
	assert.Equal(t, &domain.Rekap{
		Id:          id,
		CsName:      "Joko",
		CusName:     "Moro",
		RekapStatus: true,
		PrintStatus: true,
		RekapDate:   time.Now().Unix(),
	}, rekap)
}
