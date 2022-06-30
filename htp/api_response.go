package htp

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func ApiResponse(w http.ResponseWriter, statusCode int, statusTxt string, data interface{}) {
	response, err := json.Marshal(struct {
		StatusCode int         `json:"status_code"`
		StatusTxt  string      `json:"status_txt"`
		Data       interface{} `json:"data"`
	}{
		statusCode,
		statusTxt,
		data,
	})
	if err != nil {
		response = []byte(fmt.Sprintf(`{"status_code":500, "status_txt":"%s", "data":null}`, err.Error()))
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Content-Length", strconv.Itoa(len(response)))
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	w.Header().Set("Access-Control-Allow-Methods", "*")
	w.WriteHeader(statusCode)
	w.Write(response)
}

func Api_ResponseErr(statusTxt interface{}) string {
	response, err := json.Marshal(struct {
		StatusCode int         `json:"status_code"`
		StatusTxt  string      `json:"status_txt"`
		Data       interface{} `json:"data"`
	}{
		700,
		fmt.Sprint(statusTxt),
		"[]",
	})
	if err != nil {
		log.Println(err)
		return fmt.Sprintf("{\"status_code\":\"%v\",\"status_txt\":\"%v\",\"data\":\"[]\"", 700, err)
	}
	return string(response)
}

func Api_Response(statusCode int, statusTxt string, data interface{}) string {
	m := new(struct {
		Status_Code int         `json:"status_code"`
		Status_Txt  string      `json:"status_txt"`
		Data        interface{} `json:"data"`
	})
	m.Data = data
	m.Status_Code = statusCode
	m.Status_Txt = statusTxt
	jsonstr, err := json.Marshal(m)
	if err != nil {
		log.Println(err)
		return fmt.Sprintf("{\"status_code\":\"%v\",\"status_txt\":\"%v\",\"data\":\"[]\"", 700, err)
	}
	return string(jsonstr)
}
