package handler

import (
	"context"
	"fmt"

	"github.com/micro/go-log"
	hello "mysamples/microsample/proto/hello"
)

type HelloSvc struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *HelloSvc) SayHello(ctx context.Context, req *hello.Request, rsp *hello.Response) error {
	log.Log("Received HelloSvc.SayHello request")

	rsp.Msg = fmt.Sprintf("\"{\"res\":\"hello %s\"}\"", req.Name)

	return nil
}
