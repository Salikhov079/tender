package app

import (
	"fmt"
	"log"

	api "github.com/dilshodforever/tender/internal/http"
	"github.com/dilshodforever/tender/internal/pkg/config"
	handler "github.com/dilshodforever/tender/internal/http/handlers"
	_ "github.com/dilshodforever/tender/internal/http/docs"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"github.com/dilshodforever/tender/internal/pkg/genprotos"
)

func Run(cf config.Config) {
	TenderConn, err := grpc.NewClient("localhost:50050", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("Error while connecting to TenderService: ", err.Error())
	}
	defer TenderConn.Close()
	// Redis connection
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	
	// // MinIO connection
	// minioClient, err := minio.New("minio:9000", &minio.Options{
	// 	Creds:  credentials.NewStaticV4("Dior", "isakov05@", ""),
	// 	Secure: false,
	// })
	// if err != nil {
	// 	log.Fatal("Error while connecting to MinIO: ", err.Error())
	// }

	// Creating gRPC clients for the services
	
	
	
	BidService := genprotos.NewBidServiceClient(TenderConn)
	TenderService := genprotos.NewTenderServiceClient(TenderConn)

	// Handler with MinIO and other services
	h := handler.NewHandler(BidService, TenderService, rdb)

	// Setting up the API with the handlers
	r := api.NewGin(h)
	fmt.Println("Server started on port:8080")

	err = r.Run(":8080")
	if err != nil {
		log.Fatal("Error while running server: ", err.Error())
	}
}
