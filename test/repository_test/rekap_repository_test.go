package repository_test

import (
	"context"
	"testing"
	"time"

	"github.com/faridlan/lans-app/model/domain"
	"github.com/faridlan/lans-app/repository"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/integration/mtest"
)

func TestInsertOne(t *testing.T) {

	rekapRepository := repository.NewRekapRepository()
	mt := mtest.New(t, mtest.NewOptions().ClientType(mtest.Mock))
	defer mt.Close()

	mt.Run("success", func(t *mtest.T) {
		id := primitive.NewObjectID()
		mt.AddMockResponses(mtest.CreateSuccessResponse())
		_, err := rekapRepository.CreateOne(context.Background(), mt.DB, domain.Rekap{
			Id:          id,
			CsName:      "Udin",
			CusName:     "Jhon",
			RekapStatus: false,
			PrintStatus: true,
			RekapDate:   time.Now().Unix(),
		})

		assert.Nil(t, err)
		// assert.Equal(t, &domain.Rekap{
		// 	Id:          id,
		// 	CsName:      "Udin",
		// 	CusName:     "Jhon",
		// 	RekapStatus: false,
		// 	PrintStatus: true,
		// 	RekapDate:   time.Now().Unix(),
		// }, result)
	})
}
