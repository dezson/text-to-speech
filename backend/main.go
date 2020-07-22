package main

import (
	"flag"
	"fmt"
	"net"

	pb "github.com/dezson/text-to-speech/api"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct{}

func (server) Say(ctx context.Context, text *pb.Text) (*pb.Speech, error) {
	return nil, fmt.Errorf("Not implemented")
}

func main() {
	port := flag.Int("p", 8080, "Port to listen to")
	flag.Parse()

	logrus.Infof("Listening port: %d", *port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		logrus.Fatalf("Could not listen to port %d: %v", *port, err)
	}

	s := grpc.NewServer()
	pb.RegisterTextToSpeechServer(s, server{})
	err = s.Serve(lis)
	if err != nil {
		logrus.Fatalf("Could not serve %v", err)
	}
}
