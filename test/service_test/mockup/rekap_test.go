package mockup

import (
	"context"
	"testing"
	"time"

	"github.com/faridlan/lans-app/model/web"
	"github.com/faridlan/lans-app/repository"
	"github.com/faridlan/lans-app/service"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestCreate(t *testing.T) {

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		rekapCollection := *mt.Coll
		rekapRepository := repository.NewRekapRepository(&rekapCollection)
		rs := service.NewRekapService(rekapRepository)
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		result := rs.Create(context.Background(), web.RekapCreateRequest{
			CsName:      "Udin",
			CusName:     "Jhon",
			RekapStatus: true,
			PrintStatus: false,
			RekapDate:   time.Now().Unix(),
		})

		assert.Equal(t, web.RekapResponse{
			Id:          result.Id,
			CsName:      result.CsName,
			CusName:     result.CusName,
			RekapStatus: result.RekapStatus,
			PrintStatus: result.PrintStatus,
			RekapDate:   result.RekapDate,
		}, result)
	})
}

func TestUpdate(t *testing.T) {

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		rekapCollection := *mt.Coll
		rekapRepository := repository.NewRekapRepository(&rekapCollection)
		rekapService := service.NewRekapService(rekapRepository)

		id := primitive.NewObjectID()

		rekapReq := web.RekapUpdateRequest{
			Id:          id.Hex(),
			CsName:      "Udin",
			CusName:     "Jhon",
			RekapStatus: false,
			PrintStatus: false,
			RekapDate:   time.Now().Unix(),
		}

		mt.AddMockResponses(bson.D{
			{Key: "ok", Value: 1},
			{Key: "value", Value: bson.D{
				{Key: "_id", Value: rekapReq.Id},
				{Key: "cs_name", Value: rekapReq.CsName},
				{Key: "cus_name", Value: rekapReq.CusName},
				{Key: "rekap_status", Value: rekapReq.RekapStatus},
				{Key: "print_status", Value: rekapReq.PrintStatus},
				{Key: "rekap_date", Value: rekapReq.RekapDate},
			}},
		})

		result := rekapService.Update(context.Background(), rekapReq)
		assert.NotEqual(t, rekapReq, result)
	})

}

func TestFindById(t *testing.T) {
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(mt *mtest.T) {
		rekapCollection := *mt.Coll
		rerkapRepository := repository.NewRekapRepository(&rekapCollection)
		rekapService := service.NewRekapService(rerkapRepository)
		id := primitive.NewObjectID()

		// mt.AddMockResponses(mtest.CreateSuccessResponse())

		expectedResponse := web.RekapResponse{
			Id:          id.Hex(),
			CsName:      "Udin",
			CusName:     "farid",
			RekapStatus: false,
			PrintStatus: false,
			RekapDate:   time.Now().Unix(),
		}

		mt.AddMockResponses(mtest.CreateCursorResponse(1, "foo.bar", mtest.FirstBatch, bson.D{
			{Key: "_id", Value: expectedResponse.Id},
			{Key: "cs_name", Value: expectedResponse.CsName},
			{Key: "cus_name", Value: expectedResponse.CusName},
			{Key: "rekap_status", Value: expectedResponse.RekapStatus},
			{Key: "print_status", Value: expectedResponse.PrintStatus},
			{Key: "rekap_date", Value: expectedResponse.RekapDate},
		}))
		response := rekapService.FindById(context.Background(), id.Hex())
		assert.Equal(t, expectedResponse, response)
	})
}
