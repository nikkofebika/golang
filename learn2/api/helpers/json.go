package helpers

import (
	"encoding/json"
	"net/http"
)

type JSONResponse struct {
	Code int         `json:"code"`
	Data interface{} `json:"data"`
}

func ReadRequestBody(request *http.Request, body any) {
	decoder := json.NewDecoder(request.Body)
	err := decoder.Decode(body)
	PanicIfError(err)
}

func WriteResponseJSON(writer http.ResponseWriter, response any) {
	writer.Header().Set("Content-Type", "application/json")
	encoder := json.NewEncoder(writer)
	err := encoder.Encode(response)
	PanicIfError(err)
}
