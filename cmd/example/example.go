package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	"github.com/emicklei/go-restful/v3"
	"github.com/jemtucker/protogengosvc/proto"
)

func main() {
	// 1. Create the service using our handler implementation
	service := proto.NewHelloWorldService(&MyHandler{})
	service.RegisterRoutes()

	// 2. Retrieve the web service and register with go-restful
	ws := service.GetWebService()
	restful.Add(ws)

	// 3. Start the server
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// MyHandler is a concrete implementation of proto.HelloWorldHandler
type MyHandler struct{}

// SayHello is the implementation of proto.HelloWorldHandler.SayHello
func (h *MyHandler) SayHello(req *proto.SayHelloRequest) (*proto.SayHelloResponse, error) {
	if req.Name == "" {
		return nil, errors.New("missing name")
	}

	return &proto.SayHelloResponse{
		Message: fmt.Sprintf("Hello %s!", req.Name),
	}, nil
}

// SayGoodbye is the implementation of proto.HelloWorldHandler.SayGoodbye
func (h *MyHandler) SayGoodbye(req *proto.SayGoodbyeRequest) (*proto.SayGoodbyeResponse, error) {
	return &proto.SayGoodbyeResponse{
		Message: "👋",
	}, nil
}
