package mockup

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

	mt.Run("success", func(mt *mtest.T) {
		rekapCollection := *mt.Coll
		rekapRepository := repository.NewRekapRepository(&rekapCollection)
		rs := service.NewRekapService(rekapRepository)
		mt.AddMockResponses(mtest.CreateSuccessResponse())

		resutl := rs.Create(context.Background(), web.RekapCreateRequest{
			CsName:      "Udin",
			CusName:     "Jhon",
			RekapStatus: true,
			PrintStatus: false,
			RekapDate:   time.Now().Unix(),
		})

		assert.Equal(t, web.RekapResponse{
			Id:          resutl.Id,
			CsName:      "Udin",
			CusName:     "Jhon",
			RekapStatus: true,
			PrintStatus: false,
			RekapDate:   time.Now().Unix(),
		}, resutl)
	})
}

// func TestUpdate(t *testing.T) {

// 	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
// 	defer mt.Close()

// 	mt.Run("success", func(mt *mtest.T) {
// 		rekapCollection := *mt.Coll
// 		rerkapRepository := repository.NewRekapRepository(&rekapCollection)
// 		rekapService := service.NewRekapService(rerkapRepository)
// 	})

// }
