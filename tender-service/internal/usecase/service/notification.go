package service

import (
	"context"
	"fmt"

	pb "github.com/dilshodforever/tender/internal/pkg/genprotos"
	"github.com/dilshodforever/tender/internal/storage"
)

type NotificationService struct {
	stg storage.StorageI
	pb.UnimplementedNotificationServiceServer
}

func NewNotificationService(stg storage.StorageI) *NotificationService {
	return &NotificationService{stg: stg}
}

func (s *NotificationService) CreateNotification(ctx context.Context, req *pb.CreateNotificationRequest) (*pb.NotificationResponse, error) {
	fmt.Println(11111)
	return s.stg.Notification().CreateNotification(ctx, req)
}