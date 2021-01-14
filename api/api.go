package api

import (
	"aminuolawale/ajocard/helpers"
	"aminuolawale/ajocard/grpc/status"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net"

	"github.com/gorilla/mux"
	"google.golang.org/grpc"
	
)



func statusHandler(w http.ResponseWriter, r *http.Request){
	response := helpers.CheckStatus()
	json.NewEncoder(w).Encode(response)
	helpers.WriteLogs(r, response)
}


func StartRestServer(){
	router := mux.NewRouter()
	router.Use(helpers.PanicHandler)
	router.HandleFunc("/status", statusHandler).Methods("GET")
	fmt.Println("REST server listening on :8888")
	log.Fatal(http.ListenAndServe(":8888", router))
}

func StartGrpcServer(){
	lis, err := net.Listen("tcp", ":8889")
	helpers.HandleError(err)
	s := status.Server{}
	grpcServer := grpc.NewServer()
	status.RegisterStatusServiceServer(grpcServer, &s)
	fmt.Println("grpc server listening on :8889")
	err = grpcServer.Serve(lis)
	helpers.HandleError(err)

}