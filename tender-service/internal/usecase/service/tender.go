package service

import (
	"context"
	"fmt"

	pb "github.com/dilshodforever/tender/internal/pkg/genprotos"
	"github.com/dilshodforever/tender/internal/storage"
)

type TenderService struct {
	stg storage.StorageI
	pb.UnimplementedTenderServiceServer
}

func NewTenderService(stg storage.StorageI) *TenderService {
	return &TenderService{stg: stg}
}

func (s *TenderService) CreateTender(ctx context.Context, req *pb.CreateTenderRequest) (*pb.TenderResponse, error) {
	return s.stg.Tender().CreateTender(ctx, req)
}

func (s *TenderService) UpdateTender(ctx context.Context, req *pb.UpdateTenderRequest) (*pb.TenderResponse, error) {
	return s.stg.Tender().UpdateTender(ctx, req)
}

func (s *TenderService) DeleteTender(ctx context.Context, req *pb.TenderIdRequest) (*pb.TenderResponse, error) {
	return s.stg.Tender().DeleteTender(ctx, req)
}

func (s *TenderService) ListTenders(ctx context.Context, req *pb.ListTendersRequest) (*pb.ListTendersResponse, error) {
	return s.stg.Tender().ListTenders(ctx, req)
}

func (s *TenderService) TenderAward(ctx context.Context, req *pb.CreatTenderAwardRequest) (*pb.TenderResponse, error) {
	res, err := s.stg.Tender().TenderAward(ctx, req)
	if err != nil {
		return nil, err
	}
	response, err := s.stg.Bid().GetByTenderId(ctx, req.BidId)
	if err != nil {
		return nil, err
	}
	fmt.Println(response)
	_, err = s.stg.Notification().CreateNotification(ctx, &pb.CreateNotificationRequest{UserId: response.ContactrorId,
		Message: "Your Bid Accepted!!!", RelationId: req.BidId})
	if err != nil {

		return nil, err
	}
	return res, nil

}

func (s *TenderService) ListUserTenders(ctx context.Context, req *pb.TenderIdRequest) (*pb.ListTendersResponse, error) {
	return s.stg.Tender().ListUserTenders(ctx, req)
}
