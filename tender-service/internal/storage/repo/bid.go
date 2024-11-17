package repo

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/google/uuid"
	pb "github.com/dilshodforever/tender/internal/pkg/genprotos"
)

type BidStorage struct {
	db *sql.DB
}

func NewBidStorage(db *sql.DB) *BidStorage {
	return &BidStorage{db: db}
}

// Submit a new bid
func (b *BidStorage) SubmitBid(ctx context.Context, req *pb.SubmitBidRequest) (*pb.BidResponse, error) {
	id := uuid.NewString()
	query := `
		INSERT INTO bids (id, tender_id, contractor_id, price, delivery_time, comments, status, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, 'pending', now())
	`

	_, err := b.db.ExecContext(ctx, query, id, req.TenderId, req.ContractorId, req.Price, req.DeliveryTime, req.Comments)
	if err != nil {
		return nil, err
	}

	return &pb.BidResponse{Message: "Bid submitted successfully"}, nil
}

// List all bids with optional filtering
func (b *BidStorage) ListBids(ctx context.Context, req *pb.ListBidsRequest) (*pb.ListBidsResponse, error) {
	var bids pb.ListBidsResponse
	var queryBuilder strings.Builder

	queryBuilder.WriteString(`
		SELECT id, tender_id, price, delivery_time, comments, status, created_at
		FROM bids
		WHERE deleted_at = 0
	`)

	var args []interface{}
	argCounter := 1

	// Filtering by delivery time
	if req.DeliveryTime > 0 {
		queryBuilder.WriteString(fmt.Sprintf(" AND delivery_time = $%d", argCounter))
		args = append(args, req.DeliveryTime)
		argCounter++
	}

	// Adding LIMIT clause (default: 10)
	if req.Limit > 0 {
		queryBuilder.WriteString(fmt.Sprintf(" LIMIT $%d", argCounter))
		args = append(args, req.Limit)
		argCounter++
	} else {
		queryBuilder.WriteString(" LIMIT 10")
	}

	// Adding OFFSET clause (default: 0)
	if req.Offset > 0 {
		queryBuilder.WriteString(fmt.Sprintf(" OFFSET $%d", argCounter))
		args = append(args, req.Offset)
		argCounter++
	}

	// Execute the query
	rows, err := b.db.QueryContext(ctx, queryBuilder.String(), args...)
	if err != nil {
		return nil, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	// Process the rows
	for rows.Next() {
		var bid pb.GetAllBidResponse
		err := rows.Scan(&bid.Id, &bid.TenderId, &bid.Price, &bid.DeliveryTime, &bid.Comments, &bid.Status, &bid.CreatedAt)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row: %w", err)
		}
		
		bids.Bids = append(bids.Bids, &bid)
	}

	// Check for row iteration errors
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during row iteration: %w", err)
	}

	return &bids, nil
}

func (b *BidStorage) GetAllBidsByTenderId(ctx context.Context, req *pb.GetAllByid) (*pb.ListBidsResponse, error) {
	query := `
		SELECT 
			b.tender_id, b.price, b.delivery_time, b.comments, b.status, b.created_at,
			t.id, t.client_id, t.title, t.description, t.deadline, t.budget, t.status, t.file_url, t.created_at
		FROM bids b
		JOIN tenders t ON b.tender_id = t.id
		WHERE b.tender_id = $1 and deleted_at=0
	`

	rows, err := b.db.QueryContext(ctx, query, req.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	response := &pb.ListBidsResponse{}
	for rows.Next() {
		var bid pb.GetAllBidResponse
		var tender pb.GetTenderResponse
		err := rows.Scan(
			&bid.TenderId, &bid.Price, &bid.DeliveryTime, &bid.Comments, &bid.Status, &bid.CreatedAt,
			&tender.Id, &tender.ClientId, &tender.Title, &tender.Description, &tender.Deadline, &tender.Budget, &tender.Status, &tender.FileUrl, &tender.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		bid.Tenders = &tender
		response.Bids = append(response.Bids, &bid)
	}
	return response, nil
}

func (b *BidStorage) ListContractorBids(ctx context.Context, req *pb.GetAllByid) (*pb.GetAllBidsByUserIdRequest, error) {
	query := `
		SELECT 
			b.contractor_id, b.contactor_id,b.price, b.delivery_time, b.comments,
			t.id, t.client_id, t.title, t.description, t.deadline, t.budget, t.status, t.file_url, t.created_at
		FROM bids b
		JOIN tenders t ON b.tender_id = t.id
		WHERE b.contractor_id = $1 and deleted_at=0
	`

	rows, err := b.db.QueryContext(ctx, query, req.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	response := &pb.GetAllBidsByUserIdRequest{}
	for rows.Next() {
		var bid pb.GetAllBidsByUser
		var tender pb.GetTenderResponse
		err := rows.Scan(
			&bid.ContractorId,&bid.ContractorId, &bid.Price, &bid.DeliveryTime, &bid.Comments,
			&tender.Id, &tender.ClientId, &tender.Title, &tender.Description, &tender.Deadline, &tender.Budget, &tender.Status, &tender.FileUrl, &tender.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		bid.Tenders = &tender
		response.Bids = append(response.Bids, &bid)
	}
	return response, nil
}



func (b *BidStorage) GetByTenderId(ctx context.Context, id string) (*pb.GetAllBidResponse, error) {
	fmt.Println(id)
    query := `
        SELECT contractor_id, price, delivery_time, comments, status, created_at
        FROM bids 
        WHERE id=$1 AND deleted_at=0
    `
    var bid pb.GetAllBidResponse

    // Execute the query to fetch a single row
    err := b.db.QueryRowContext(ctx, query, id).Scan(
        &bid.ContactrorId,
        &bid.Price,
        &bid.DeliveryTime,
        &bid.Comments,
        &bid.Status,
        &bid.CreatedAt,
    )
    if err != nil {
        if err == sql.ErrNoRows {
            // Return nil if no rows are found (optional)
            return nil, err
        }
        return nil, err
    }


    return &bid, nil
}



