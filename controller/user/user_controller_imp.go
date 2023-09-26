package controller

import (
	"encoding/json"
	"github.com/afdikomayte/sim-layanan-sertifikat-tanah/helper"
	"github.com/afdikomayte/sim-layanan-sertifikat-tanah/model/web"
	service "github.com/afdikomayte/sim-layanan-sertifikat-tanah/service/user"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type UserControllerImpl struct {
	UserService service.UserService
}

func (controller *UserControllerImpl) Index(w http.ResponseWriter, r *http.Request, params httprouter.Params) {

	controller.UserService.FindAll(r.Context())

	datas := web.DataToPage{
		"pageName": "Halaman User",
	}
	web.HtmlResponse(w, "user/index.html", "index", datas)

}

func (controller *UserControllerImpl) Add(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	datas := web.DataToPage{
		"pageName": "Halaman Tambah User",
	}
	web.HtmlResponse(w, "user/add.html", "add", datas)
}

func (controller *UserControllerImpl) Store(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	userRequest := web.UserRequest{}
	err := decoder.Decode(&userRequest)
	helper.PanicIfError(err)
	panic("implement me")
}

func (controller *UserControllerImpl) Edit(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//TODO implement me
	panic("implement me")
}

func (controller *UserControllerImpl) Update(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//TODO implement me
	panic("implement me")
}

func (controller *UserControllerImpl) Destroy(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//TODO implement me
	panic("implement me")
}

func (controller *UserControllerImpl) FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	//TODO implement me
	panic("implement me")
}
