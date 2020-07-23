package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	pb "github.com/dezson/text-to-speech/api"
	"google.golang.org/grpc"
)

func main() {
	backend := flag.String("b", "localhost:8080", "address of the say backend")
	output := flag.String("o", "output.wav", "Wav file where the output will be written")
	flag.Parse()

	if flag.NArg() < 1 {
		fmt.Println("Usage: \n\t say <text> ")
		os.Exit(1)
	}

	conn, err := grpc.Dial(*backend, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to %s: %v", *backend, err)
	}
	defer conn.Close()

	client := pb.NewTextToSpeechClient(conn)
	text := &pb.Text{Text: flag.Arg(0)}
	res, err := client.Say(context.Background(), text)
	if err != nil {
		log.Fatalf("Could not say %s: %v", text.Text, err)
	}

	if err := ioutil.WriteFile(*output, res.Audio, 0666); err != nil {
		log.Fatalf("Could not write to %s: %v", *output, err)
	}
}
