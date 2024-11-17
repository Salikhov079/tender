package repo

import (
	"context"
	"database/sql"
	"fmt"

	pb "github.com/dilshodforever/tender/internal/pkg/genprotos"
	"github.com/google/uuid"
	
)

type NotificationStorage struct {
	db *sql.DB
}

func NewNotificationStorage(db *sql.DB) *NotificationStorage {
	return &NotificationStorage{db: db}
}

func (n *NotificationStorage) CreateNotification(ctx context.Context, req *pb.CreateNotificationRequest) (*pb.NotificationResponse, error) {
	id := uuid.NewString()
	fmt.Println("not",n.db.Stats())
	query := `
		INSERT INTO notifications (id, user_id, message, relation_id, type, created_at)
		VALUES ($1, $2, $3, $4, $5, now())
	`
	_, err := n.db.ExecContext(ctx, query, id, req.UserId, req.Message, req.RelationId, req.Type)
	if err != nil {
		return nil, err
	}

	return &pb.NotificationResponse{Message: "Notification send successfully"}, nil
}
