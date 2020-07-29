package controller

import (
	"net/http"

	"studentrecord/util"
	"studentrecord/service"
	"studentrecord/common"
)

// PersOprDefnByPersID 根据 pers_id 获取 PerOprDefn 
func PersOprDefnByPersID(w http.ResponseWriter, req *http.Request) {
	reqJSON := util.HandRequest(req)
	respBody, httpStatusCode := service.SelectPersOprDefnByPersID(reqJSON)
	common.GetFormatter().JSON(w, httpStatusCode, respBody)
	return
}

// PersOprDefnByOprID 根据 opr_id 获取 PerOprDefn 
func PersOprDefnByOprID(w http.ResponseWriter, req *http.Request) {
	reqJSON := util.HandRequest(req)
	respBody, httpStatusCode := service.SelectPersOprDefnByOprID(reqJSON)
	common.GetFormatter().JSON(w, httpStatusCode, respBody)
	return
}
