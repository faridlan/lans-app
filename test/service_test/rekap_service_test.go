package service_test

import (
	"context"
	"testing"
	"time"

	"github.com/faridlan/lans-app/model/web"
	"github.com/faridlan/lans-app/repository"
	"github.com/faridlan/lans-app/service"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestCreate(t *testing.T) {

	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(t *mtest.T) {
		rekapCollection := *mt.Coll
		rekapRepository := repository.NewRekapRepository(&rekapCollection)
		rekapService := service.NewRekapService(rekapRepository)

		mt.AddMockResponses(mtest.CreateSuccessResponse())

		response := rekapService.Create(context.Background(), web.RekapCreateRequest{
			CsName:      "Jhon",
			CusName:     "Udin",
			RekapStatus: false,
			PrintStatus: false,
			RekapDate:   time.Now().Unix(),
		})

		assert.Equal(t, &web.RekapResponse{
			CsName:      "Jhon",
			CusName:     "Udin",
			RekapStatus: false,
			PrintStatus: false,
			RekapDate:   time.Now().Unix(),
		}, response)
	})
}
