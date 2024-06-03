package helper

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

type Response struct {
	Meta Meta        `json:"meta"`
	Data interface{} `json:"data"`
}

type Meta struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
	Success bool   `json:"success"`
}

func APIResponse(w http.ResponseWriter, resp Response) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(resp.Meta.Code)
	jsonData, _ := json.Marshal(&resp)
	w.Write(jsonData)
}

func APIResponseFailed(w http.ResponseWriter, meta Meta) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(meta.Code)
	jsonData, _ := json.Marshal(&meta)
	w.Write(jsonData)
}

func APIResponseSuccessWithoutData(w http.ResponseWriter, meta Meta) {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(meta.Code)
	jsonData, _ := json.Marshal(&meta)
	w.Write(jsonData)
}
