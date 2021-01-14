package main

import (
	"log"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	"aminuolawale/ajocard/grpc/status"
	"aminuolawale/ajocard/helpers"
)

func main() {

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":8889", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	s := status.NewStatusServiceClient(conn)

	response, err := s.GetStatus(context.Background(), &status.Status{})
	helpers.HandleError(err)
	log.Printf("Response Body: %s", response)

}