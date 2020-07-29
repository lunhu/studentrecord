package common

import (
	"reflect"
)

type CjData struct {
	Name	string	`json:"name"`
	Value	string	`json:"value"`
	Prompt	string	`json:"prompt"`
}

type CjItem struct {
	Rel		string		`json:"rel"`
	Href	string		`json:"href"`
	Data	[]*CjData	`json:"data"`
	
}

type CjError struct {
	Title   string	`json:"title"`
	Code	string	`json:"code"`
	Message	string	`json:"message"`
}


type CjCollection struct {
	Version	string		`json:"version"`
	Href	string		`json:"href"`
	Items	[]*CjItem	`json:"items"`
	Error	*CjError	`json:"error"`
}

type CjResponse struct {
	Collection	*CjCollection	`json:"collection"`
}


func HandItem(stru interface{}, httpStatusCode int) (*CjResponse, int) {

	var cjDatas []*CjData

	t := reflect.TypeOf(stru)
	v := reflect.ValueOf(stru)

	
	for i := 0; i < t.NumField(); i++ {
        // Get the field, returns https://golang.org/pkg/reflect/#StructField
		field := t.Field(i)
		value := v.Field(i).Interface()

        //Get the field tag value
		cj := field.Tag.Get("cj")
		prompt := field.Tag.Get("prompt")
		cjData := &CjData{}
		cjData.Name = cj
		cjData.Value = value.(string)
		cjData.Prompt = prompt
		cjDatas = append(cjDatas, cjData)
        
	}
	
	cjItem := &CjItem{}
	cjItem.Data = cjDatas

	cjCollection := &CjCollection{}
	cjCollection.Items =append(cjCollection.Items, cjItem)
	
	return HandResponse(cjCollection), httpStatusCode

}

func HandError(entry, catalog string) (*CjResponse, int) {

	cjError := &CjError{}
    httpStatusCode := 200
	switch entry {
	case "400":
		cjError.Title = "The request cannot be completed"
		cjError.Code = "Bad Request Code x400"
		cjError.Message = "The request was rejected by the server. Check the parameters, please!"
		httpStatusCode = 400
	case "500":
		cjError.Title = "Unable to search item"
		cjError.Code = "Internal Error Code x500"
		cjError.Message = "The server encountered an internal error or misconfiguration and was unable to complete your request"
		httpStatusCode = 500
	default:
		cjError.Title = "An unknown error occurred"
		cjError.Code = "Internal Error Code x000"
		cjError.Message = "The server encountered an unknown error or misconfiguration and was unable to complete your request"
		httpStatusCode = 500
	}

	cjCollection :=&CjCollection{}
	
	cjCollection.Version = "1.0"

	cjCollection.Error = cjError

	return HandResponse(cjCollection), httpStatusCode

}

func HandResponse(cjCollection *CjCollection) *CjResponse {
	cjResponse := &CjResponse{}
	cjResponse.Collection = cjCollection
	return cjResponse
}


