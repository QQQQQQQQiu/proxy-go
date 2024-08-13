package module

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"proxy-go/types"
	"proxy-go/store"
	"strings"
)

const API_PREFIX_xhr = "/api/xhr"

var Client = http.DefaultClient

func HandlerXHR(resp http.ResponseWriter, req *http.Request) {
	options, errStr := getOptions(req)
	if errStr != "" {
		log.Printf("Error getting options: %s\n", errStr)
		http.Error(resp, fmt.Sprintf("Error getting options: %s", errStr), http.StatusBadRequest)
		return
	}
	var url = options.URL
	var method = options.Method
	var body = strings.NewReader(options.Body)
	var headers = options.Headers
	var isThrowHeaders = options.ThrowHeaders

	log.Printf("Making XHR request: %s %s\n", method, url)

	reqOut, err := http.NewRequest(method, url, body)
	if err != nil {
		log.Printf("Error creating request: %s\n", err)
		http.Error(resp, fmt.Sprintf("Error creating request: %s", err), http.StatusInternalServerError)
		return
	}

	for key, value := range headers {
		reqOut.Header.Set(key, fmt.Sprintf("%v", value))
	}

	respOut, err := Client.Do(reqOut)
	if err != nil {
		log.Printf("Error executing request: %s\n", err)
		http.Error(resp, fmt.Sprintf("Error executing request: %s", err), http.StatusInternalServerError)
		return
	}

	defer respOut.Body.Close()


	responseHeaders := make(map[string]string)
	for key, values := range respOut.Header {
		responseHeaders[key] = values[0]
		resp.Header().Set(key, values[0])
	}

	if isThrowHeaders {
		respBody, err := io.ReadAll(respOut.Body)
		xhrResponse := types.XHRResponseAll{
			StatusCode: respOut.StatusCode,
			Headers:    responseHeaders,
			Body:       string(respBody),
		}
		response, err := json.Marshal(xhrResponse)
		if err != nil {
			log.Println("Error marshaling response:", err)
			http.Error(resp, "Error marshaling response", http.StatusInternalServerError)
			return
		}
		resp.Write(response)
		return
	}
	
	if _, err := io.Copy(resp, respOut.Body); err != nil {
		log.Printf("Error writing response: %s\n", err)
	}
	resp.WriteHeader(respOut.StatusCode)
}

func getOptions(req *http.Request) (types.XHRData, string) {
    url := req.URL.Path
    isPassSecret := HandlerXHR_is_pass_secret(req)
    if req.Method == http.MethodGet {
        var url_str string
        if isPassSecret {
            secret := store.SecretCtl("get", "")
            url_str = url[len(API_PREFIX_xhr+"/"+secret+"/"):]
        } else {
            url_str = url[len(API_PREFIX_xhr+"/"):]
        }
        log.Println("url_str:", url_str)
        if strings.HasPrefix(url_str, "http") {
            options := types.XHRData{
                URL: url_str,
								Method: "GET",
            }
						return options, ""
        } else {
					var data types.XHRData
					reader := strings.NewReader(url_str)
					if err := json.NewDecoder(reader).Decode(&data); err != nil {
							log.Println("Invalid request body:", err)
							return types.XHRData{}, "Invalid request body"
					}
					return data, ""
        }
    } else if req.Method == http.MethodPost {
			var data types.XHRData
			if err := json.NewDecoder(req.Body).Decode(&data); err != nil {
					log.Println("Invalid request body:", err)
					return types.XHRData{}, "Invalid request body"
			}
			return data, ""
    } else {
        log.Println("Invalid request method:", req.Method)
        return types.XHRData{}, "Invalid request method"
    }
}

func HandlerXHR_is_match_route(req *http.Request) bool {
	url := req.URL.Path
	return strings.HasPrefix(url, API_PREFIX_xhr)
}

func HandlerXHR_is_pass_secret(req *http.Request) bool {
	url := req.URL.Path
	secret := store.SecretCtl("get", "")
	if url == fmt.Sprintf("%s/%s", API_PREFIX_xhr, secret) {
		return true
	}
	return strings.HasPrefix(url, fmt.Sprintf("%s/%s/", API_PREFIX_xhr, secret))
}