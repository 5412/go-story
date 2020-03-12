package ipc

import (
    "testing"
)

type EchoServer struct {
}

func (server *EchoServer) Handle(method, params string) *Response {
    code := method
    body := params
    return &Response{code, body}
}

func (server *EchoServer) Name() string {
    return "EchoServer"
}

func TestIpc(t *testing.T) {
    server := NewIpcServer(&EchoServer{})

    client1 := NewIpcClient(server)
    client2 := NewIpcClient(server)

    resp1, _  := client1.Call("200", "method")
    resp2, _  := client2.Call("300", "method")

    resp3:= Response{"200", "method"}
    resp4 := Response{"300", "method"}

    if *resp1 != resp3 || *resp2 != resp4 {
        t.Error("IpcClient.Call failed. resp1:", resp1, "resp2:", resp2, "resp3:", resp3, "resp4:", resp4)
    }

    client1.Close()
    client2.Close()
}

