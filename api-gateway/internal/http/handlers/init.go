package handlers

import (
	pb "github.com/dilshodforever/tender/internal/pkg/genprotos"
	"github.com/go-redis/redis/v8"
	//"github.com/minio/minio-go/v7"
)

type HTTPHandler struct {
	//Minio          *minio.Client
	TenderService pb.TenderServiceClient
	BidService    pb.BidServiceClient
	Redis         *redis.Client
}

func NewHandler( //Minio *minio.Client,
	BidService pb.BidServiceClient,
	TenderService pb.TenderServiceClient,
	Redis        *redis.Client) *HTTPHandler {

	return &HTTPHandler{
		TenderService: TenderService,
		BidService:    BidService,
		Redis:         Redis,
		//Minio:              Minio,
	}
}
