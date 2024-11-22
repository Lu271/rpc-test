package main

import (
	"context"
	"fmt"
	hello "github.com/Lu271/rpc-test/hello-server/kitex_gen/hello"
)

// HelloServiceImpl implements the last service interface defined in the IDL.
type HelloServiceImpl struct{}

// SayHello implements the HelloServiceImpl interface.
func (s *HelloServiceImpl) SayHello(ctx context.Context, req *hello.HelloRequest) (resp *hello.HelloResponse, err error) {
	resp = &hello.HelloResponse{}
	resp.Message = fmt.Sprintf("I am Hello server, I receive %v", req.Message)
	return
}
