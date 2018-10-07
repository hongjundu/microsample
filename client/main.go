package main

import (
	"context"
	"fmt"
	micro "github.com/micro/go-micro"
	proto "mysamples/microsample/proto/hello"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// Create a new service. Optionally include some options here.

	service := micro.NewService(micro.Name("go.micro.srv.hello"))
	service.Init()

	// Create new hello client
	hello := proto.NewHelloService("go.micro.srv.hello", service.Client())

	for i := 0; i < 100; i++ {

		// Call the greeter

		rsp, err := hello.SayHello(context.TODO(), &proto.Request{Name: "Du Hongjun"})
		if err != nil {
			fmt.Println(err)
			return
		}

		// Print response
		fmt.Println(rsp.Msg)
		time.Sleep(time.Second)

	}

	exitchan := make(chan os.Signal, 1)
	signal.Notify(exitchan, os.Interrupt)
	signal.Notify(exitchan, syscall.SIGTERM)
	signal.Notify(exitchan, syscall.SIGKILL)
	<-exitchan

}
