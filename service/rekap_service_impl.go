package service

import (
	"context"

	"github.com/faridlan/lans-app/helper/model"
	"github.com/faridlan/lans-app/model/domain"
	"github.com/faridlan/lans-app/model/web"
	"github.com/faridlan/lans-app/repository"
)

type RekapServiceImpl struct {
	RekapRepo repository.RekapRepository
}

func NewRekapService(RekapRepo repository.RekapRepository) RekapService {
	return &RekapServiceImpl{}
}

func (service *RekapServiceImpl) Create(ctx context.Context, request web.RekapCreateRequest) web.RekapResponse {
	rekap := domain.Rekap{
		CsName:      request.CsName,
		CusName:     request.CsName,
		RekapStatus: request.RekapStatus,
		PrintStatus: request.PrintStatus,
		RekapDate:   request.RekapDate,
	}

	result, err := service.RekapRepo.CreateOne(ctx, rekap)
	if err != nil {
		panic(err)
	}

	return model.RekapResponse(*result)
}
