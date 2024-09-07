package types

type CommandData struct {
	ID     string                 `json:"id"`
	Cmd  string                 `json:"cmd"`
}

type CommandResponse struct {
	ID      string `json:"id"`
	Output  string `json:"output"`
	Error   string `json:"error,omitempty"`
}

type XHRData struct {
	URL     string                 `json:"url"`
	Method  string                 `json:"method"`
	Headers map[string]interface{} `json:"headers"`
	Body   string                 `json:"body"`
	ThrowHeaders bool `json:"throwHeaders"` // 是否带响应头
}

type XHRResponseAll struct {
    StatusCode int                    `json:"status_code"`
    Headers    map[string]string      `json:"headers"`
    Body       string                 `json:"body"`
}
