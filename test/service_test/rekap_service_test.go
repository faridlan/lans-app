package service_test

import (
	"context"
	"testing"
	"time"

	"github.com/faridlan/lans-app/app"
	"github.com/faridlan/lans-app/model/web"
	"github.com/faridlan/lans-app/repository"
	"github.com/faridlan/lans-app/service"
	"github.com/stretchr/testify/assert"
)

func TestServiceInsert(t *testing.T) {

	client := app.NewDatabase()
	Collection := client.Database("lans_app").Collection("rekap")
	repository := repository.NewRekapRepository(Collection)
	service := service.NewRekapService(repository)

	response := service.Create(context.Background(), web.RekapCreateRequest{
		CsName:      "Farid",
		CusName:     "Lan",
		RekapStatus: true,
		PrintStatus: true,
		RekapDate:   time.Now().Unix(),
	})

	assert.Equal(t, web.RekapResponse{
		Id:          response.Id,
		CsName:      response.CsName,
		CusName:     response.CusName,
		RekapStatus: response.RekapStatus,
		PrintStatus: response.PrintStatus,
		RekapDate:   response.RekapDate,
	}, response)
}

func TestServiceUpdate(t *testing.T) {

}
