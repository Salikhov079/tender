package server

import (
	"fmt"
	"log"
	"net"

	"github.com/dilshodforever/nasiya-savdo/api"
	"github.com/dilshodforever/nasiya-savdo/api/handler"
	"github.com/dilshodforever/nasiya-savdo/config"
	pb "github.com/dilshodforever/nasiya-savdo/genprotos"
	"github.com/dilshodforever/nasiya-savdo/service"
	"github.com/dilshodforever/nasiya-savdo/storage/postgres"
	r "github.com/dilshodforever/nasiya-savdo/storage/redis"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func Run() error {
	cnf := config.Load()
	db, err := postgres.NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
		return err
	}

	liss, err := net.Listen("tcp", cnf.GrpcPort)
	if err != nil {
		log.Fatal("Error while connection on tcp: ", err.Error())
		return err
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, service.NewUserService(db))
	log.Printf("server listening at %v", liss.Addr())
	go func() {
		if err := s.Serve(liss); err != nil {
			log.Fatalf("failed to serve: %v", err)
		}
	}()

	userConn, err := grpc.NewClient(fmt.Sprintf("localhost%s", cnf.GrpcPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while NewClient: ", err.Error())
		return err
	}
	defer userConn.Close()

	client := redis.NewClient(&redis.Options{
		Addr: cnf.RedisHost + cnf.RedisPort,
	})

	rdb := r.NewInMemoryStorage(client)

	us := pb.NewUserServiceClient(userConn)
	h := handler.NewHandler(us, rdb)
	r := api.NewGin(h)

	fmt.Println("Server started on port:", cnf.HTTPPort)

	err = r.Run(cnf.HTTPPort)
	if err != nil {
		log.Fatal("Error while Run: ", err.Error())
		return err
	}

	return nil
}
