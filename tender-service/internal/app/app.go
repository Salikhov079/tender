package app

import (
	"log"
	"net"

	"github.com/dilshodforever/tender/internal/pkg/config"
	pb "github.com/dilshodforever/tender/internal/pkg/genprotos"
	"github.com/dilshodforever/tender/internal/pkg/postgres"
	"github.com/dilshodforever/tender/internal/storage/repo"
	"github.com/dilshodforever/tender/internal/usecase/service"

	"google.golang.org/grpc"
)

func Run(cf *config.Config) {
	// Connect to Postgres
	pgm, err := postgres.New(cf)
	if err != nil {
		log.Fatal(err)
	}
	defer pgm.Close()

	// Initialize storage
	db := repo.NewStorage(pgm.DB)

	// Set up gRPC server
	lis, err := net.Listen("tcp", cf.GRPCPort)
	if err != nil {
		log.Fatal("Failed to listen: ", err)
	}
	server := grpc.NewServer()

	// Register services
	pb.RegisterBidServiceServer(server, service.NewBidService(db))
	pb.RegisterTenderServiceServer(server, service.NewTenderService(db))

	log.Println("Server started on port " + cf.GRPCPort)
	if err = server.Serve(lis); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	defer lis.Close()
}
