package module

import (
	"strings"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
)

var API_PREFIX_doc = "/doc";

func HandlerDoc (w http.ResponseWriter, r *http.Request)  {
	// 读取文件内容
	cwd,_ := os.Getwd()
	readmePath := filepath.Join(cwd, "readme.md")
	data, err := os.ReadFile(readmePath)
	if err != nil {
			fmt.Println("Error reading file:", err)
			return
	}

	// 将文件内容转换为字符串
	content := string(data)
	w.WriteHeader(http.StatusOK)
	w.Write( []byte(content) )
}

func HandlerDoc_is_match_route(req *http.Request) bool {
	var url = req.URL.Path
	return strings.HasPrefix(url, API_PREFIX_doc)
}