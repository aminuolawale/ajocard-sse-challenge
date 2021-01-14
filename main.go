package main

import (
	"aminuolawale/ajocard/api"
	"os"
)

func main(){
	if len(os.Args) > 1 && os.Args[1] == "grpc" {
		api.StartGrpcServer()
	} else {
		api.StartRestServer()
	}	
	
}