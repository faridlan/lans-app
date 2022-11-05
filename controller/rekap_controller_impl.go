package controller

import (
	"net/http"

	"github.com/faridlan/lans-app/helper"
	"github.com/faridlan/lans-app/model/web"
	"github.com/faridlan/lans-app/service"
	"github.com/julienschmidt/httprouter"
)

type RekapControllerImpl struct {
	RekapService service.RekapService
}

func (controller *RekapControllerImpl) Create(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	rekapCreateReq := web.RekapCreateRequest{}
	helper.ReadFromRequestBody(request, &rekapCreateReq)

	rekapResponse := controller.RekapService.Create(request.Context(), rekapCreateReq)

	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   rekapResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RekapControllerImpl) Update(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("id")
	rekapUpdateReq := web.RekapUpdateRequest{}
	helper.ReadFromRequestBody(request, &rekapUpdateReq)

	rekapUpdateReq.Id = id

	rekapResponse := controller.RekapService.Update(request.Context(), rekapUpdateReq)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   rekapResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RekapControllerImpl) Delete(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("id")

	controller.RekapService.Delete(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}

	helper.WriteToResponseBody(writer, webResponse)
}

func (controller *RekapControllerImpl) FindById(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {
	id := params.ByName("id")

	rekapResponse := controller.RekapService.FindById(request.Context(), id)
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   rekapResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)

}

func (controller *RekapControllerImpl) FindAll(writer http.ResponseWriter, request *http.Request, params httprouter.Params) {

	rekapResponse := controller.RekapService.FindAll(request.Context())
	webResponse := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   rekapResponse,
	}

	helper.WriteToResponseBody(writer, webResponse)
}
