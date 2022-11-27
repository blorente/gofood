package main

import (
	"context"
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

type Server struct {
	pb.UnimplementedMealSuggesterServer
	backend *echo.Echo
	app pocketbase.PocketBase
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
	pb.RegisterMealSuggesterServer(s, &Server{backend: router, app: app})
	reflection.Register(s)

	bold := color.New(color.Bold).Add(color.FgGreen)
	bold.Printf("> Server started at: %s\n", color.CyanString("localhost:%s", port))
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func ToMealKind(kind string) pb.MealKind {
	switch kind {
	case "breakfast": return pb.MealKind_Breakfast
	case "dinner": return pb.MealKind_Dinner
	default: log.Fatalf("Bad mealkind! %s", kind)
	}
	return pb.MealKind_Breakfast
}

func (server *Server) SuggestMeal(ctx context.Context, req *pb.SuggestMealRequest) (*pb.SuggestMealResponse, error) {
	log.Printf("Received request for meal with filter: %v", req.GetFilters())
	meals, err := server.app.Dao().FindRecordsByExpr("meals", nil)
	if err != nil {
		return nil, err
	}
	log.Printf("Got %d meals!", len(meals))
	suggestions := make([]*pb.MealSuggestion, 0, 2)
	for _, mealRecord := range meals {
		data := mealRecord.SchemaData()
		suggestion := &pb.MealSuggestion{
			Kind: req.GetKind(),
			Name: data["name"].(string),
			TotalKcal: data["total_kcal"].(float64),
		}
		suggestions = append(suggestions, suggestion)
	}
	suggestions = append(suggestions, 
			&pb.MealSuggestion{
				Kind: pb.MealKind_Breakfast,
				Name: "Omelette",
				TotalKcal: 19,
				Foods: []*pb.MealFood{},
			},
	)
	return &pb.SuggestMealResponse{
		Suggestions: suggestions,
	}, nil
}
