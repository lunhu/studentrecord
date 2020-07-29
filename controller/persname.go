package controller

import (
	"net/http"

	"studentrecord/util"
	"studentrecord/service"
	"studentrecord/common"
)

// PersNameAdd 增加一条人员姓名
func PersNameAdd(w http.ResponseWriter, req *http.Request) {

	reqJSON := util.HandRequest(req)
	respBody, httpStatusCode := service.PersNameAdd(reqJSON)
	if httpStatusCode==201 {
		w.Header().Add("Location", "/api/v1/name/"+ respBody.Collection.Items[0].Data[0].Value)
	}
	
	common.GetFormatter().JSON(w, httpStatusCode, respBody)
	return
}