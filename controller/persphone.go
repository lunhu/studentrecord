package controller

import (
	"net/http"

	"studentrecord/util"
	"studentrecord/service"
	"studentrecord/common"
)

// PersPhoneAdd 增加一条人员邮箱
func PersPhoneAdd(w http.ResponseWriter, req *http.Request) {

	reqJSON := util.HandRequest(req)

	respBody, httpStatusCode := service.PersPhoneAdd(reqJSON)
	if httpStatusCode==201 {
		w.Header().Add("Location", "/api/v1/phone/"+ reqJSON.Get(`template.data.#(name=="pers_id").value`).String())
	}
	common.GetFormatter().JSON(w, httpStatusCode, respBody)
	return
}