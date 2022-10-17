package model

import (
	"github.com/faridlan/lans-app/model/domain"
	"github.com/faridlan/lans-app/model/web"
)

func RekapResponse(rekap domain.Rekap) web.RekapResponse {
	return web.RekapResponse{
		Id:          rekap.Id.Hex(),
		CsName:      rekap.CsName,
		CusName:     rekap.CusName,
		RekapStatus: rekap.RekapStatus,
		PrintStatus: rekap.PrintStatus,
		RekapDate:   rekap.RekapDate,
	}

}