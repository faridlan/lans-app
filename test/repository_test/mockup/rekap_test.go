package mockup

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
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		result, err := rekapRepository.CreateOne(context.Background(), domain.Rekap{
			CsName:      "Udin",
			CusName:     "Jhon",
			RekapStatus: true,
			PrintStatus: false,
			RekapDate:   time.Now().Unix(),
		})

		assert.Nil(t, err)
		assert.Equal(t, &domain.Rekap{
			Id:          result.Id,
			CsName:      result.CsName,
			CusName:     result.CusName,
			RekapStatus: result.RekapStatus,
			PrintStatus: result.PrintStatus,
			RekapDate:   result.RekapDate,
		}, result)
	})
}

func TestUpdateOne(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		rekapCollection := *mt.Coll
		rekapRepository := repository.NewRekapRepository(&rekapCollection)
		rekap := domain.Rekap{
			Id:          primitive.NewObjectID(),
			CsName:      "John",
			CusName:     "Udin",
			RekapStatus: true,
			PrintStatus: false,
			RekapDate:   time.Now().Unix(),
		}
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "value", Value: bson.D{
				{Key: "_id", Value: rekap.Id},
				{Key: "cs_name", Value: rekap.CsName},
				{Key: "cus_name", Value: rekap.CusName},
				{Key: "rekap_status", Value: rekap.RekapStatus},
				{Key: "print_status", Value: rekap.PrintStatus},
				{Key: "rekap_date", Value: rekap.RekapDate},
			}},
		})

		result, err := rekapRepository.UpdateOne(context.Background(), rekap)

		assert.Nil(t, err)
		assert.Equal(t, &rekap, result)
	})
}

func TestDeleteOne(t *testing.T) {

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	mt.Close()

	mt.Run("success", func(mt *mtest.T) {

		rekapCollection := *mt.Coll
		rekapRepository := repository.NewRekapRepository(&rekapCollection)
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "acknowledged", Value: true},
			{Key: "n", Value: 1},
		})
		err := rekapRepository.DeleteOne(context.Background(), domain.Rekap{Id: primitive.NewObjectID()})
		assert.Nil(t, err)
	})

	mt.Run("no document deleted", func(mt *mtest.T) {

		rekapCollection := *mt.Coll
		rekapRepository := repository.NewRekapRepository(&rekapCollection)
		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "acknowledged", Value: true},
			{Key: "n", Value: 0},
		})
		err := rekapRepository.DeleteOne(context.Background(), domain.Rekap{Id: primitive.NewObjectID()})
		assert.NotNil(t, err)
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

func TestFind(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		rekapCollection := *mt.Coll
		rekapRepository := repository.NewRekapRepository(&rekapCollection)
		id1 := primitive.NewObjectID()
		id2 := primitive.NewObjectID()
		epoch := time.Now().Unix()

		first := mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: id1},
			{Key: "cs_name", Value: "jhon"},
			{Key: "cus_name", Value: "udin"},
			{Key: "rekap_status", Value: false},
			{Key: "print_status", Value: true},
			{Key: "rekap_date", Value: epoch},
		})

		second := mtest.CreateCursorResponse(1, "foo.bar", mtest.NextBatch, bson.D{
			{Key: "_id", Value: id2},
			{Key: "cs_name", Value: "doe"},
			{Key: "cus_name", Value: "farid"},
			{Key: "rekap_status", Value: true},
			{Key: "print_status", Value: true},
			{Key: "rekap_date", Value: epoch},
		})
		killCursors := mtest.CreateCursorResponse(0, "foo.bar", mtest.NextBatch)
		mt.AddMockResponses(first, second, killCursors)

		results, err := rekapRepository.FindMany(context.Background())
		assert.Nil(t, err)
		assert.Equal(t, []domain.Rekap{
			{
				Id:          id1,
				CsName:      "jhon",
				CusName:     "udin",
				RekapStatus: false,
				PrintStatus: true,
				RekapDate:   epoch,
			},
			{
				Id:          id2,
				CsName:      "doe",
				CusName:     "farid",
				RekapStatus: true,
				PrintStatus: true,
				RekapDate:   epoch,
			},
		}, results)
	})
}
