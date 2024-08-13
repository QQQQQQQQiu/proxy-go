package module

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"proxy-go/types"
	"strings"
)


func HandleXHR( body types.ReqData, w http.ResponseWriter, r *http.Request) {
	var xhrData types.XHRData
    if err := json.Unmarshal(body.Data, &xhrData); err != nil {
        log.Println("Invalid xhr data:", err)
        http.Error(w, "Invalid xhr data", http.StatusBadRequest)
        return
    }
    log.Printf("Making XHR request: %s %s\n", xhrData.Method, xhrData.URL)

    
    // 设置请求参数
    req, err := http.NewRequest(xhrData.Method, xhrData.URL, strings.NewReader(xhrData.Body))
    if err != nil {
        log.Printf("Error creating request: %s\n", err)
        http.Error(w, fmt.Sprintf("Error creating request: %s", err), http.StatusInternalServerError)
        return
    }
    // 设置请求头
    for key, value := range xhrData.Headers {
        req.Header.Set(key, fmt.Sprintf("%v", value))
    }

	
	// 发送自定义请求
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
			log.Printf("Error executing request: %s\n", err)
			http.Error(w, fmt.Sprintf("Error executing request: %s", err), http.StatusInternalServerError)
			return
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
			log.Printf("Error reading response body: %s\n", err)
			http.Error(w, fmt.Sprintf("Error reading response body: %s", err), http.StatusInternalServerError)
			return
	}

	w.WriteHeader(resp.StatusCode)
	responseHeaders := make(map[string]string)
	for key, value := range resp.Header {
			responseHeaders[key] = value[0]
			w.Header().Set(key, value[0])
	}

	var isThrowHeaders bool = xhrData.ThrowHeaders
	if isThrowHeaders {
			xhrResponse := types.XHRResponseAll{
					StatusCode: resp.StatusCode,
					Headers:    responseHeaders,
					Body:       string(respBody),
			}
			response, err := json.Marshal(xhrResponse)
			if err != nil {
					log.Println("Error marshaling response:", err)
					http.Error(w, "Error marshaling response", http.StatusInternalServerError)
					return
			}
			w.Write(response)
			return
	}

	w.Write(respBody)
}