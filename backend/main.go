package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net"
	"os/exec"

	pb "github.com/dezson/text-to-speech/api"
	"github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct{}

func (server) Say(ctx context.Context, text *pb.Text) (*pb.Speech, error) {
	f, err := ioutil.TempFile("", "")
	if err != nil {
		return nil, fmt.Errorf("Could not create file: %v", err)
	}
	if err := f.Close(); err != nil {
		return nil, fmt.Errorf("Could not close %s file: %v", f.Name(), err)
	}

	cmd := exec.Command("flite", "-t", text.Text, "-o", f.Name())
	if data, err := cmd.CombinedOutput(); err != nil {
		return nil, fmt.Errorf("flite failed: %s", data)
	}

	data, err := ioutil.ReadFile(f.Name())
	if err != nil {
		return nil, fmt.Errorf("Could not read %s file: %v", f.Name(), err)
	}

	return &pb.Speech{Audio: data}, nil
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
