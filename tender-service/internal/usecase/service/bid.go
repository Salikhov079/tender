package service

import (
	"context"

	pb "github.com/dilshodforever/tender/internal/pkg/genprotos"
	"github.com/dilshodforever/tender/internal/storage"
)

type BidService struct {
	stg storage.StorageI
	pb.UnimplementedBidServiceServer
}

func NewBidService(stg storage.StorageI) *BidService {
	return &BidService{stg: stg}
}

func (s *BidService) SubmitBid(ctx context.Context, req *pb.SubmitBidRequest) (*pb.BidResponse, error) {
	return s.stg.Bid().SubmitBid(ctx, req)
}


func (s *BidService) ListBids(ctx context.Context, req *pb.ListBidsRequest) (*pb.ListBidsResponse, error) {
	return s.stg.Bid().ListBids(ctx, req)
}

func (s *BidService) GetAllBidsByTenderId(ctx context.Context, req *pb.GetAllByid) (*pb.ListBidsResponse, error) {
	return s.stg.Bid().GetAllBidsByTenderId(ctx, req)
}

func (s *BidService) ListContractorBids(ctx context.Context, req *pb.GetAllByid) (*pb.GetAllBidsByUserIdRequest, error) {
	return s.stg.Bid().ListContractorBids(ctx, req)
}

func (s *BidService) GetByTenderId(ctx context.Context, id string) (*pb.GetAllBidResponse, error) {
	return s.stg.Bid().GetByTenderId(ctx, id)
}

