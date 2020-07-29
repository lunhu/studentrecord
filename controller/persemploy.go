package controller

import (
	"net/http"

	"studentrecord/util"
	"studentrecord/service"
	"studentrecord/common"
)

// PersEmployAdd 增加一条人员工作
func PersEmployAdd(w http.ResponseWriter, req *http.Request) {

	reqJSON := util.HandRequest(req)

	respBody, httpStatusCode := service.PersEmployAdd(reqJSON)
	if httpStatusCode==201 {
		w.Header().Add("Location", "/api/v1/employ/"+ reqJSON.Get(`template.data.#(name=="pers_id").value`).String())
	}
	common.GetFormatter().JSON(w, httpStatusCode, respBody)
	return
}