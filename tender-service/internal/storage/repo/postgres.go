package repo

import (
	"database/sql"

	"github.com/dilshodforever/tender/internal/storage"
)

type Storage struct {
	BidS          storage.BidI    // BidI for bid-related operations
	TenderS       storage.TenderI // TenderI for tender-related operations
	NotificationS storage.NotificationI
	DB            *sql.DB
}

func NewStorage(db *sql.DB) *Storage {
	return &Storage{
		BidS:          NewBidStorage(db),    // Create a function to initialize Bid storage
		TenderS:       NewTenderStorage(db), // Create a function to initialize Tender storage
		NotificationS: NewNotificationStorage(db),
		DB:            db,
	}
}

func (s *Storage) Bid() storage.BidI {
	return s.BidS
}

func (s *Storage) Tender() storage.TenderI {
	return s.TenderS
}

func (s *Storage) Notification() storage.NotificationI {
	return s.NotificationS
}
