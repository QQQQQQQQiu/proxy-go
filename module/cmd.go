package module

import (
	"context"
	"encoding/json"
	"io"
	"fmt"
	"log"
	"net/http"
	"os/exec"
	"proxy-go/types"
	"strings"
	"time"
)

var API_PREFIX_cmd = "/api/cmd";

func HandleCommand(resp http.ResponseWriter, req *http.Request) {
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Println("Failed to read request body:", err)
		http.Error(resp, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer req.Body.Close()

	var commandData []types.CommandData
	if err := json.Unmarshal(body, &commandData); err != nil {
		log.Println("Invalid request body:", err)
		http.Error(resp, "Invalid request body", http.StatusBadRequest)
		return
	}

	results := make([]types.CommandResponse, 0)

	for _, cmd := range commandData {
			log.Printf("Executing command: %s\n", cmd.Cmd)

			// 一些字段校验
			if cmd.ID == "" || cmd.Cmd == "" {
					log.Println("Invalid command:", cmd)
					http.Error(resp, "Invalid command", http.StatusBadRequest)
					return
			}

			// 动态处理 ping 命令
			if strings.HasPrefix(cmd.Cmd, "ping") {
					// 这里可以根据需要添加更多的参数处理逻辑
					cmd.Cmd = fmt.Sprintf("%s -c 4", cmd.Cmd) // 默认限制为 4 次
			}

			// 创建一个带超时的 context
			ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
			defer cancel()

			// 使用 context 来执行命令
			out, err := exec.CommandContext(ctx, "sh", "-c", cmd.Cmd).CombinedOutput()
			response := types.CommandResponse{ID: cmd.ID, Output: "" }

			if ctx.Err() == context.DeadlineExceeded {
					log.Printf("Command %s timed out\n", cmd.ID)
					response.Error = "Error: command timed out"
			} else if err != nil {
					log.Printf("Command execution error for %s: %s\n", cmd.ID, err)
					response.Error = fmt.Sprintf("Error: %s", err)
			} else {
					response.Output = string(out)
			}

			results = append(results, response)
	}

	response, err := json.Marshal(results)
	if err != nil {
			log.Println("Error marshaling response:", err)
			http.Error(resp, "Error marshaling response", http.StatusInternalServerError)
			return
	}

	resp.WriteHeader(http.StatusOK)
	resp.Write(response)
}

func HandleCommand_is_match_route(req *http.Request) bool {
	var url = req.URL.Path
	return strings.HasPrefix(url, API_PREFIX_cmd)
}