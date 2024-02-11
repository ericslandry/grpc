package main

import (
	"context"
	"log"
	"net"
	"os"

	pb "grpc/pb"

	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
)

func main() {
	app := &cli.App{
		Name:  "server",
		Usage: "This is the server part of the CLI application",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "port",
				Value: "8080",
				Usage: "Port to listen on",
			},
		},
		Action: func(cCtx *cli.Context) error {
			port := cCtx.String("port")
			lis, err := net.Listen("tcp", ":"+port)
			if err != nil {
				log.Fatalf("failed to listen: %v", err)
			}
			s := grpc.NewServer()
			pb.RegisterGreeterServer(s, &server{})
			log.Printf("server listening at %v", lis.Addr())
			if err := s.Serve(lis); err != nil {
				log.Fatalf("failed to serve: %v", err)
			}
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloResponse{
		Greeting: "Hello, " + in.GetName(),
	}, nil
}
