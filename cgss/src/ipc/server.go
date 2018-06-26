package ipc

import (
    "encoding/json"
    "fmt"
)

type Request struct {
    // 属性名 类型 tag 系统对接时可以使用tag代表属性名
    Method string "method"
    Params string "params"
}

type Response struct {
    Code string "code"
    Body string "body"
}

type Server interface {
    Name() string
    Handle(method, params string) *Response
}

type IpcServer struct {
    Server
}

func NewIpcServer(server Server) *IpcServer {
    return &IpcServer{server}
}

func (server *IpcServer)Connect() chan string {
    session := make(chan string, 0)

    // 匿名函数由参数直接调用方法 func (arge int) {} (arge)
    go func (c chan string) {
        for {
            request := <-c

            if request == "CLOSE" {
                break
            }
            var req Request
            err := json.Unmarshal([]byte(request), &req)
            if err != nil {
                fmt.Println("Invalid request format:", request)
            }

            resp := server.Handle(req.Method, req.Params)

            b, err := json.Marshal(resp)
            c <- string(b)
        }

        fmt.Println("Session closed.")
    }(session)

    fmt.Println("A new session has been created successfully.")
    return session
}

