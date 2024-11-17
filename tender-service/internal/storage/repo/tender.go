package repo

import (
	"bytes"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	pb "github.com/dilshodforever/tender/internal/pkg/genprotos"
	"github.com/google/uuid"
)

// TenderStorage handles database operations for tenders.
type TenderStorage struct {
	db *sql.DB
}

// NewTenderStorage creates a new TenderStorage instance.
func NewTenderStorage(db *sql.DB) *TenderStorage {
	return &TenderStorage{db: db}
}

// CreateTender creates a new tender.
func (s *TenderStorage) CreateTender(ctx context.Context, req *pb.CreateTenderRequest) (*pb.TenderResponse, error) {
	id := uuid.NewString()
	query := `
	INSERT INTO tenders (id, client_id, title, description, deadline, budget, file_url)
	VALUES ($1, $2, $3, $4, $5, $6, $7)
`
	_, err := s.db.ExecContext(ctx, query, id, req.ClientId, req.Title, req.Description, req.Deadline, req.Budget, req.FileUrl)

	if err != nil {
		return nil, err
	}
	return &pb.TenderResponse{Message: "Tender created successfully"}, nil
}

// TenderAward updates the client and status for a tender.
func (s *TenderStorage) TenderAward(ctx context.Context, req *pb.CreatTenderAwardRequest) (*pb.TenderResponse, error) {
	id := uuid.NewString()
	
	query := `
	INSERT INTO winners (id, tender_id, bid_id) VALUES ($1, $2, $3)
`
	result, err := s.db.ExecContext(ctx, query, id, req.TenderId, req.BidId)

	if err != nil {
		return nil, err
	}
	query = `UPDATE tenders 
	        SET  status='awarded'
			WHERE id=$1`

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return nil, errors.New("no tender found with the given ID")
	}
	result, err = s.db.ExecContext(ctx, query, req.TenderId)
	
	if err != nil {
		return nil, err
	}
	rowsAffected, _ = result.RowsAffected()
	if rowsAffected == 0 {
		return nil, errors.New("no tender found with the given ID")
	}
	
	return &pb.TenderResponse{Message: "Tender awarded successfully"}, nil
}

// DeleteTender deletes a tender by ID.
func (s *TenderStorage) DeleteTender(ctx context.Context, req *pb.TenderIdRequest) (*pb.TenderResponse, error) {
	query := `
		UPDATE tenders
		SET deleted_at=$1
		WHERE id = $2
	`
	result, err := s.db.ExecContext(ctx, query, time.Now().Unix(), req.Id)
	if err != nil {
		return nil, err
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return nil, errors.New("no tender found with the given ID")
	}
	return &pb.TenderResponse{Message: "Tender deleted successfully"}, nil
}

// ListTenders retrieves tenders with dynamic filtering.
func (s *TenderStorage) ListTenders(ctx context.Context, req *pb.ListTendersRequest) (*pb.ListTendersResponse, error) {
	var (
		queryBuilder bytes.Buffer
		args         []interface{}
		argCounter   = 1
	)

	queryBuilder.WriteString(`
		SELECT id, client_id, title, description, deadline, budget, status, file_url, created_at
		FROM tenders
		WHERE deleted_at=0
	`)

	if req.Title != "" {
		queryBuilder.WriteString(fmt.Sprintf(" AND title ILIKE $%d", argCounter))
		args = append(args, "%"+req.Title+"%")
		argCounter++
	}
	if req.Deadline != "" {
		queryBuilder.WriteString(fmt.Sprintf(" AND deadline >= $%d", argCounter))
		args = append(args, req.Deadline)
		argCounter++
	}
	if req.Limit > 0 {
		queryBuilder.WriteString(fmt.Sprintf(" LIMIT $%d", argCounter))
		args = append(args, req.Limit)
		argCounter++
	}
	if req.Offset > 0 {
		queryBuilder.WriteString(fmt.Sprintf(" OFFSET $%d", argCounter))
		args = append(args, req.Offset)
		argCounter++
	}

	rows, err := s.db.QueryContext(ctx, queryBuilder.String(), args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	response := &pb.ListTendersResponse{}
	for rows.Next() {
		var tender pb.GetTenderResponse
		err := rows.Scan(
			&tender.Id, &tender.ClientId, &tender.Title, &tender.Description,
			&tender.Deadline, &tender.Budget, &tender.Status, &tender.FileUrl, &tender.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		response.Tenders = append(response.Tenders, &tender)
	}

	return response, nil
}

// UpdateTender updates tender details.
func (s *TenderStorage) UpdateTender(ctx context.Context, req *pb.UpdateTenderRequest) (*pb.TenderResponse, error) {
	query := `
		UPDATE tenders
		SET 
			title = $1,
			description = $2,
			deadline = $3,
			budget = $4,
			updated_at = now()
		WHERE id = $5
	`
	result, err := s.db.ExecContext(ctx, query, req.Title, req.Description, req.Deadline, req.Budget, req.Id)
	if err != nil {
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil || rowsAffected == 0 {
		return nil, errors.New("no tender found with the given ID")
	}

	return &pb.TenderResponse{Message: "Tender updated successfully"}, nil
}

// ListUserTenders retrieves tenders created by a specific user.
func (s *TenderStorage) ListUserTenders(ctx context.Context, req *pb.TenderIdRequest) (*pb.ListTendersResponse, error) {
	query := `
		SELECT id, client_id, title, description, deadline, budget, status, file_url, created_at
		FROM tenders
		WHERE client_id = $1
	`
	rows, err := s.db.QueryContext(ctx, query, req.Id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	response := &pb.ListTendersResponse{}
	for rows.Next() {
		var tender pb.GetTenderResponse
		err := rows.Scan(
			&tender.Id, &tender.ClientId, &tender.Title, &tender.Description, &tender.Deadline, &tender.Budget, &tender.Status, &tender.FileUrl, &tender.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		response.Tenders = append(response.Tenders, &tender) // `Binds` to'g'ri ishlatilgan
	}

	return response, nil
}
