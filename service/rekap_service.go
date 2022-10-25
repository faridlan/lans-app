package service

import (
	"context"

	"github.com/faridlan/lans-app/model/web"
)

type RekapService interface {
	Create(ctx context.Context, request web.RekapCreateRequest) web.RekapResponse
	Update(ctx context.Context, requset web.RekapUpdateRequest) web.RekapResponse
	Delete(ctx context.Context, rekapId string)
	FindById(ctx context.Context, rekapId string) web.RekapResponse
	FindAll(ctx context.Context) []web.RekapResponse
}
