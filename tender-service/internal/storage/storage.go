package storage

import (
	"context"

	pb "github.com/dilshodforever/tender/internal/pkg/genprotos"
)

type StorageI interface {
	Tender() TenderI
	Bid() BidI
	Notification() NotificationI
}

type TenderI interface {
	CreateTender(context.Context, *pb.CreateTenderRequest) (*pb.TenderResponse, error)
	UpdateTender(context.Context, *pb.UpdateTenderRequest) (*pb.TenderResponse, error)
	DeleteTender(context.Context, *pb.TenderIdRequest) (*pb.TenderResponse, error)
	ListTenders(context.Context, *pb.ListTendersRequest) (*pb.ListTendersResponse, error)
	TenderAward(context.Context, *pb.CreatTenderAwardRequest) (*pb.TenderResponse, error)
	ListUserTenders(context.Context, *pb.TenderIdRequest) (*pb.ListTendersResponse, error)
}

type BidI interface {
	SubmitBid(ctx context.Context, req *pb.SubmitBidRequest) (*pb.BidResponse, error)
	ListBids(ctx context.Context, req *pb.ListBidsRequest) (*pb.ListBidsResponse, error)
	GetAllBidsByTenderId(ctx context.Context, req *pb.GetAllByid) (*pb.ListBidsResponse, error)
	ListContractorBids(ctx context.Context, req *pb.GetAllByid) (*pb.GetAllBidsByUserIdRequest, error)
	GetByTenderId(ctx context.Context, id string) (*pb.GetAllBidResponse, error)

}

type NotificationI interface {
	CreateNotification(ctx context.Context, req *pb.CreateNotificationRequest) (*pb.NotificationResponse, error)
}
