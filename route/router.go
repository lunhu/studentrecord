package route

import (
	"github.com/gorilla/mux"
	"studentrecord/controller"
	"studentrecord/util"
)

func InitRoutes(mx *mux.Router) {
	InitRoutesForOprDefn(mx)
	InitRoutesForNames(mx)
	InitRoutesForEmails(mx)
	InitRoutesForPhones(mx)
	InitRoutesForEmploys(mx)
	InitRoutesForInfos(mx)
}

// 初始化账号相关的路由
func InitRoutesForOprDefn(mx *mux.Router) {
	submx := mx.PathPrefix("/api/v1").Subrouter()
	submx.HandleFunc("/oprdefns/pers/{pers_id}", controller.PersOprDefnByPersID).Methods("GET")
	submx.HandleFunc("/oprdefns/opr/{opr_id}", controller.PersOprDefnByOprID).Methods("GET")
	submx.Use(util.NewAuth().Middleware)
}

func InitRoutesForNames(mx *mux.Router) {
	submx := mx.PathPrefix("/api/v1").Subrouter()
	submx.HandleFunc("/names", controller.PersNameAdd).Methods("POST")
	submx.Use(util.NewAuth().Middleware)
}

func InitRoutesForEmails(mx *mux.Router) {
	submx := mx.PathPrefix("/api/v1").Subrouter()
	submx.HandleFunc("/emails", controller.PersEmailAdd).Methods("POST")
	submx.Use(util.NewAuth().Middleware)
}

func InitRoutesForPhones(mx *mux.Router) {
	submx := mx.PathPrefix("/api/v1").Subrouter()
	submx.HandleFunc("/phones", controller.PersPhoneAdd).Methods("POST")
	submx.Use(util.NewAuth().Middleware)
}

func InitRoutesForEmploys(mx *mux.Router) {
	submx := mx.PathPrefix("/api/v1").Subrouter()
	submx.HandleFunc("/employs", controller.PersEmployAdd).Methods("POST")
	submx.Use(util.NewAuth().Middleware)
}

func InitRoutesForInfos(mx *mux.Router) {
	submx := mx.PathPrefix("/api/v1").Subrouter()
	submx.HandleFunc("/infos", controller.PersInfoAdd).Methods("POST")
	submx.Use(util.NewAuth().Middleware)
}