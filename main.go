package main

import (
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"mysamples/microsample/handler"
	hello "mysamples/microsample/proto/hello"
)

func main() {
	// New Service
	service := micro.NewService(
		micro.Name("go.micro.srv.hello"),
		micro.Version("latest"),
	)

	// Initialise service
	service.Init()

	// Register Handler
	hello.RegisterHelloHandler(service.Server(), new(handler.HelloSvc))

	// Run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
