package main

import (
	"log"
	"flag"
	"net/http"
	module "proxy-go/module"
	"proxy-go/store"
)

func main() {
    // 定义命令行参数-s
    var secret string
    var port string
    flag.StringVar(&secret, "s", "", "接口令牌")
    flag.StringVar(&port, "p", "801", "端口")
    flag.Parse()
    store.SecretCtl("set", secret)
    store.PortCtl("set", port)
    if secret == "" {
        log.Fatal("高危程序接口令牌不能为空，运行程序加 -s 参数")
    }
    // 启动 HTTP 服务器
    server := &http.Server{
        Addr:           ":"+ port,
        Handler:        http.HandlerFunc(handlerRequest),
    }
    log.Println("监听端口: ", port, " 接口令牌: ", secret)
    log.Fatal(server.ListenAndServe())
}

func handlerRequest(resp http.ResponseWriter, req *http.Request) {
    resp.Header().Set("Access-Control-Allow-Origin", "*")
    resp.Header().Set("Access-Control-Allow-Methods", "*")
    resp.Header().Set("Access-Control-Allow-Headers", "*")

    if req.Method == http.MethodOptions {
        resp.WriteHeader(http.StatusOK)
        return
    }
    var is_secret_pass bool = false
    switch {
        case module.HandlerDoc_is_match_route(req):
            is_secret_pass = true
        case module.HandlerXHR_is_pass_secret(req):
            is_secret_pass = true
    }
    var is_check_secret_header = check_secret_header(req) 
    if !is_check_secret_header && !is_secret_pass {
        resp.WriteHeader(http.StatusUnauthorized)
        return
    }
    switch {
        case module.HandlerDoc_is_match_route(req):
            module.HandlerDoc(resp, req)
        case module.HandleCommand_is_match_route(req):
            module.HandleCommand(resp, req)
        case module.HandlerXHR_is_match_route(req):
            module.HandlerXHR(resp, req)
        default:
            resp.WriteHeader(http.StatusNotFound)
    }
    return
}

func check_secret_header(req *http.Request) bool {
    var s = req.Header.Get("s")
    var secret = store.SecretCtl("get", "")
    return s == secret
}
