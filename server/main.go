package main

import (
  pb "github.com/blorente/gofood/server/pb"
  "google.golang.org/grpc"
  "google.golang.org/grpc/reflection"
  "log"
  "net"
)

func main() {
	log.Println("HI!")
  listener, err := net.Listen("tcp", ":8080")
  if err != nil {
    log.Fatalln(err)
  }
  s := grpc.NewServer()
  pb.RegisterMealSuggesterServer(s, &server{})
  reflection.Register(s)
  if err := s.Serve(listener); err != nil {
    log.Fatalf("failed to serve: %v", err)
  }
}
type server struct {
  pb.UnimplementedMealSuggesterServer
}

