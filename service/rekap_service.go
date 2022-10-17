package service

import (
	"context"

	"github.com/faridlan/lans-app/model/web"
)

type RekapService interface {
	Create(ctx context.Context, request web.RekapCreateRequest) web.RekapResponse
}
