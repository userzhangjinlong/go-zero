// Code generated by goctl. DO NOT EDIT!
// Source: hello.proto

package server

import (
	"context"

	greetlogic "github.com/userzhangjinlong/go-zero/tools/goctl/example/rpc/hello/internal/logic/greet"
	"github.com/userzhangjinlong/go-zero/tools/goctl/example/rpc/hello/internal/svc"
	"github.com/userzhangjinlong/go-zero/tools/goctl/example/rpc/hello/pb/hello"
)

type GreetServer struct {
	svcCtx *svc.ServiceContext
	hello.UnimplementedGreetServer
}

func NewGreetServer(svcCtx *svc.ServiceContext) *GreetServer {
	return &GreetServer{
		svcCtx: svcCtx,
	}
}

func (s *GreetServer) SayHello(ctx context.Context, in *hello.HelloReq) (*hello.HelloResp, error) {
	l := greetlogic.NewSayHelloLogic(ctx, s.svcCtx)
	return l.SayHello(in)
}
