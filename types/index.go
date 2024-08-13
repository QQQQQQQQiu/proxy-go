package types

import "encoding/json"

type ReqData struct {
	S string          `json:"s"`
	Type          string          `json:"type"`
	Data          json.RawMessage `json:"data"`
}

type CommandData struct {
	ID     string                 `json:"id"`
	Cmd  string                 `json:"cmd"`
}

type CommandRequest struct {
	Type string    `json:"type"`
	Data []CommandData `json:"data"`
}



type CommandResponse struct {
	ID      string `json:"id"`
	Output  string `json:"output"`
	Error   string `json:"error,omitempty"`
}


type XHRRequest struct {
	Type string `json:"type"`
	Data XHRData `json:"data"`
}

type XHRData struct {
	URL     string                 `json:"url"`
	Method  string                 `json:"method"`
	Headers map[string]interface{} `json:"headers"`
	Body   string                 `json:"body"`
	ThrowHeaders bool `json:"throwHeaders"` // 是否带响应头
}

type XHRResponse string
type XHRResponseAll struct {
    StatusCode int                    `json:"status_code"`
    Headers    map[string]string      `json:"headers"`
    Body       string                 `json:"body"`
}
