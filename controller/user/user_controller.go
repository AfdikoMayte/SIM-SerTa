package controller

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

type UserController interface {
	Index(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Add(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Store(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Edit(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Update(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	Destroy(w http.ResponseWriter, r *http.Request, params httprouter.Params)
	FindById(w http.ResponseWriter, r *http.Request, params httprouter.Params)
}
