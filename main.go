package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"
	module "proxy-go/module"
	"proxy-go/utils"
)

// 令牌
var SECRET string
// 端口
var PORT string = "801"

func main() {
    // 定义命令行参数-s
    flag.StringVar(&SECRET, "s", "", "Specify the value for m")
    flag.Parse()
    if SECRET == "" {
        log.Fatal("高危程序接口密码不能为空，运行程序加 -s 参数")
    }
    log.Println("接口请求密码: ", SECRET)

    log.Println("Listening on port " + PORT)
    http.HandleFunc("/doc", handlerDoc)
    http.HandleFunc("/api", handlerApi)
    if err := http.ListenAndServe(":" + PORT, nil); err != nil {
        log.Fatalf("Error starting server: %v", err)
    }
}

func handlerApi(w http.ResponseWriter, r *http.Request) {
    
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "*")
    w.Header().Set("Access-Control-Allow-Headers", "*")

    if r.Method == http.MethodOptions {
        w.WriteHeader(http.StatusOK)
        return
    }

    // 调用 getBody 函数解析请求体
    body, err := utils.GetBody(r)
    if err != "" {
        http.Error(w, "Invalid request body："+ err, http.StatusBadRequest)
        return
    }
    reqAuthorization := r.Header.Get("Authorization")
    if reqAuthorization == "" {
        reqAuthorization = body.S
    }
    var actionType = body.Type
    

    // 校验令牌
    if reqAuthorization != SECRET {
        log.Println("Invalid secret: " + reqAuthorization)
        http.Error(w, "Invalid secret", http.StatusUnauthorized)
        return
    }

    switch actionType {
        case "command":
            module.HandleCommand(body, w, r)
        case "xhr":
            module.HandleXHR(body, w, r)
        default:
            log.Println("Invalid type:", body.Type)
            http.Error(w, "Invalid type", http.StatusBadRequest)
    }
}




func handlerDoc (w http.ResponseWriter, r *http.Request)  {
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

