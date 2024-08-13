package utils

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"proxy-go/types"
	"strings"
)

func GetBody(r *http.Request) (types.ReqData, string) {
	var body io.Reader

	if r.Method == http.MethodGet {
			p := r.URL.Query().Get("p")
			if p == "" {
					log.Println("Invalid request: p field is empty")
					return types.ReqData{}, "Invalid request: p field is empty"
			}
			body = strings.NewReader(p)
	} else if r.Method == http.MethodPost {
			body = r.Body
	} else {
			log.Println("Invalid request method:", r.Method)
			return types.ReqData{}, "Invalid request method"
	}

	log.Println("request body:", body)

	var data types.ReqData
	if err := json.NewDecoder(body).Decode(&data); err != nil {
			log.Println("Invalid request body:", err)
			return types.ReqData{}, "Invalid request body"
	}

	return data, ""
}
