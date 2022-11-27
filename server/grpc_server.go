package main

import (
	"log"
	"net"

	pb "github.com/blorente/gofood/server/pb"
	"github.com/fatih/color"
	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedMealSuggesterServer
	backend *echo.Echo
}

func CreateAndStartServer(app pocketbase.PocketBase, port string) {
	router, err := apis.InitApi(app)
	if err != nil {
		panic(err)
	}

	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalln(err)
	}
	s := grpc.NewServer()
	pb.RegisterMealSuggesterServer(s, &server{backend: router})
	reflection.Register(s)

	bold := color.New(color.Bold).Add(color.FgGreen)
	bold.Printf("> Server started at: %s\n", color.CyanString("localhost:%s", port))
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func (*server) 
