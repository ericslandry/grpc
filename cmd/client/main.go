package main

import (
	"context"
	"log"
	"os"
	"time"

	pb "grpc/pb"

	"github.com/urfave/cli/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	app := &cli.App{
		Name:  "client",
		Usage: "This is the client part of the CLI application",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "addr",
				Value: "localhost:8080",
				Usage: "Address of the server",
			},
			&cli.StringFlag{
				Name:  "name",
				Value: "John Doe",
				Usage: "Name to greet",
			},
		},
		Action: func(cCtx *cli.Context) error {
			addr := cCtx.String("addr")
			conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
			if err != nil {
				log.Fatalf("did not connect: %v", err)
			}
			defer conn.Close()
			c := pb.NewGreeterClient(conn)
			ctx, cancel := context.WithTimeout(context.Background(), time.Second)
			defer cancel()
			name := cCtx.String("name")
			r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
			if err != nil {
				log.Fatalf("could not greet: %v", err)
			}
			log.Printf("Greeting: %s", r.GetGreeting())
			return nil
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
