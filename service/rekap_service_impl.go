package service

import (
	"context"

	"github.com/faridlan/lans-app/helper"
	"github.com/faridlan/lans-app/helper/model"
	"github.com/faridlan/lans-app/model/domain"
	"github.com/faridlan/lans-app/model/web"
	"github.com/faridlan/lans-app/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type RekapServiceImpl struct {
	RekapRepo repository.RekapRepository
}

func NewRekapService(RekapRepo repository.RekapRepository) RekapService {
	return &RekapServiceImpl{
		RekapRepo: RekapRepo,
	}
}

func (service *RekapServiceImpl) Create(ctx context.Context, request web.RekapCreateRequest) web.RekapResponse {
	rekap := domain.Rekap{
		CsName:      request.CsName,
		CusName:     request.CusName,
		RekapStatus: request.RekapStatus,
		PrintStatus: request.PrintStatus,
		RekapDate:   request.RekapDate,
	}

	result, err := service.RekapRepo.CreateOne(ctx, rekap)
	helper.PanicIfError(err)

	return model.RekapResponse(result)
}

func (service *RekapServiceImpl) Update(ctx context.Context, requset web.RekapUpdateRequest) web.RekapResponse {

	id, err := primitive.ObjectIDFromHex(requset.Id)
	helper.PanicIfError(err)

	rekap, err := service.RekapRepo.FindOne(ctx, id)
	helper.PanicIfError(err)

	rekap.CsName = requset.CsName
	rekap.CusName = requset.CusName
	rekap.RekapStatus = requset.RekapStatus
	rekap.PrintStatus = requset.PrintStatus
	rekap.RekapDate = requset.RekapDate

	rekap, err = service.RekapRepo.UpdateOne(ctx, *rekap)
	helper.PanicIfError(err)

	return model.RekapResponse(rekap)
}

func (service *RekapServiceImpl) Delete(ctx context.Context, rekapId string) {
	id, err := primitive.ObjectIDFromHex(rekapId)
	helper.PanicIfError(err)

	rekap, err := service.RekapRepo.FindOne(ctx, id)
	helper.PanicIfError(err)

	service.RekapRepo.DeleteOne(ctx, *rekap)
}

func (service *RekapServiceImpl) FindById(ctx context.Context, rekapId string) web.RekapResponse {
	id, err := primitive.ObjectIDFromHex(rekapId)
	helper.PanicIfError(err)

	rekap, err := service.RekapRepo.FindOne(ctx, id)
	helper.PanicIfError(err)

	return model.RekapResponse(rekap)
}

func (service *RekapServiceImpl) FindAll(ctx context.Context) []web.RekapResponse {
	rekap, err := service.RekapRepo.FindMany(ctx)
	helper.PanicIfError(err)
	return model.RekapResponses(rekap)
}
