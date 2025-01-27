package main

import (
	"github.com/ivansukach/book-service/protocol"
	"github.com/ivansukach/book-service/repositories"
	"github.com/ivansukach/book-service/server"
	"github.com/ivansukach/book-service/service"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

func main() {
	redisClient := repositories.NewRedisClient()
	rps := repositories.New(redisClient)
	bs := service.New(rps)
	srv := server.New(bs)
	s := grpc.NewServer()
	protocol.RegisterBookServiceServer(s, srv)
	listener, err := net.Listen("tcp", ":1323")
	if err != nil {
		log.Error(err)
		return
	}
	s.Serve(listener)
}
