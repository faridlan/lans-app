package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/faridlan/lans-app/model/domain"
	"github.com/faridlan/lans-app/repository"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestInsertOne(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		rekapCollection := *mt.Coll
		rekapRepository := repository.NewRekapRepository(&rekapCollection)
		id := primitive.NewObjectID()
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		resutl, err := rekapRepository.CreateOne(context.Background(), domain.Rekap{
			Id:          id,
			CsName:      "Udin",
			CusName:     "Jhon",
			RekapStatus: true,
			PrintStatus: false,
			RekapDate:   time.Now().Unix(),
		})

		assert.Nil(t, err)
		assert.Equal(t, &domain.Rekap{
			Id:          id,
			CsName:      "Udin",
			CusName:     "Jhon",
			RekapStatus: true,
			PrintStatus: false,
			RekapDate:   time.Now().Unix(),
		}, resutl)
	})
}

func TestFindOne(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		rekapCollection := *mt.Coll
		rekapRepository := repository.NewRekapRepository(&rekapCollection)
		expectedRekap := domain.Rekap{
			Id:          primitive.NewObjectID(),
			CsName:      "john",
			CusName:     "udin",
			RekapStatus: true,
			PrintStatus: false,
			RekapDate:   time.Now().Unix(),
		}
		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: expectedRekap.Id},
			{Key: "cs_name", Value: expectedRekap.CsName},
			{Key: "cus_name", Value: expectedRekap.CusName},
			{Key: "rekap_status", Value: expectedRekap.RekapStatus},
			{Key: "print_status", Value: expectedRekap.PrintStatus},
			{Key: "rekap_date", Value: expectedRekap.RekapDate},
		}))
		response, err := rekapRepository.FindOne(context.Background(), expectedRekap.Id)
		assert.Nil(t, err)
		assert.Equal(t, &expectedRekap, response)
	})
}
