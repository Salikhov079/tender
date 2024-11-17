package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	pb "github.com/dilshodforever/nasiya-savdo/genprotos"

	"github.com/google/uuid"
)

type UserStorage struct {
	db *sql.DB
}

func NewUserStorage(db *sql.DB) *UserStorage {
	return &UserStorage{db: db}
}

func (p *UserStorage) Register(user *pb.User) (*pb.User, error) {
	u_id := uuid.NewString()
	query := `
		INSERT INTO users (id, username, password_hash, role, email)
		VALUES ($1, $2, $3, $4, $5)
	`

	_, err := p.db.ExecContext(context.Background(), query, u_id, user.Username, user.Password, user.Role, user.Email)
	if err != nil {
		return nil, err
	}

	return &pb.User{Id: u_id, Username: user.Username, Password: user.Password, Role: user.Role, Email: user.Email}, nil
}

func (p *UserStorage) GetById(id *pb.ById) (*pb.User, error) {
	query := `
		SELECT id, username, password_hash, role, email
		FROM users 
		WHERE id = $1 AND deleted_at = 0
	`

	row := p.db.QueryRowContext(context.Background(), query, id.Id)

	var user pb.User

	err := row.Scan(
		&user.Id,
		&user.Username,
		&user.Password,
		&user.Role,
		&user.Email,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, err
	}

	return &user, nil
}

func (p *UserStorage) GetAll(filter *pb.UserFilter) (*pb.AllUsers, error) {
	users := &pb.AllUsers{}
	query := `
		SELECT id, username, password_hash, role, email 
		FROM users 
		WHERE deleted_at = 0
	`
	var params []interface{}
	count := 1

	if filter.Username != "" {
		query += fmt.Sprintf(" AND username ILIKE $%d", count)
		params = append(params, "%"+filter.Username+"%")
		count++
	}
	if filter.Email != "" {
		query += fmt.Sprintf(" AND email ILIKE $%d", count)
		params = append(params, "%"+filter.Email+"%")
		count++
	}
	if filter.Role != "" {
		query += fmt.Sprintf(" AND role = $%d", count)
		params = append(params, filter.Role)
		count++
	}

	if filter.Limit != 0 || filter.Offset != 0 {
		query += fmt.Sprintf(" LIMIT $%d OFFSET $%d", count, count+1)
		params = append(params, filter.Limit, filter.Offset)
	}

	rows, err := p.db.QueryContext(context.Background(), query, params...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var user pb.User
		err := rows.Scan(
			&user.Id,
			&user.Username,
			&user.Password,
			&user.Role,
			&user.Email,
		)
		if err != nil {
			return nil, err
		}
		users.Users = append(users.Users, &user)
	}

	query = "SELECT COUNT(1) FROM users WHERE deleted_at = 0"
	err = p.db.QueryRowContext(context.Background(), query).Scan(&count)
	if err != nil {
		return nil, err
	}
	users.Count = int32(count)

	return users, nil
}

func (p *UserStorage) Update(user *pb.User) (*pb.User, error) {
	query := `UPDATE users SET `
	args := []interface{}{}
	argCount := 1

	if user.Username != "" {
		query += fmt.Sprintf("username = $%d, ", argCount)
		args = append(args, user.Username)
		argCount++
	}
	if user.Password != "" {
		query += fmt.Sprintf("password_hash = $%d, ", argCount)
		args = append(args, user.Password)
		argCount++
	}
	if user.Role != "" {
		query += fmt.Sprintf("role = $%d, ", argCount)
		args = append(args, user.Role)
		argCount++
	}
	if user.Email != "" {
		query += fmt.Sprintf("email = $%d, ", argCount)
		args = append(args, user.Email)
		argCount++
	}

	query += fmt.Sprintf("updated_at = $%d WHERE id = $%d AND deleted_at = 0", argCount, argCount+1)
	args = append(args, time.Now(), user.Id)

	_, err := p.db.ExecContext(context.Background(), query, args...)
	if err != nil {
		return nil, err
	}

	return &pb.User{Id: user.Id, Username: user.Username, Password: user.Password, Role: user.Role, Email: user.Email}, nil
}

func (p *UserStorage) Delete(id *pb.ById) (*pb.User, error) {
	query := `
		UPDATE users SET deleted_at = $1 
		WHERE id = $2 AND deleted_at = 0
	`
	_, err := p.db.ExecContext(context.Background(), query, time.Now().Unix(), id.Id)
	if err != nil {
		return nil, err
	}
	return &pb.User{}, nil
}

func (p *UserStorage) Login(login *pb.UserLogin) (*pb.User, error) {
	query := `
		SELECT id, username, password_hash, role, email
		FROM users
		WHERE username = $1 AND deleted_at = 0
	`
	row := p.db.QueryRowContext(context.Background(), query, login.Username)

	var user pb.User

	err := row.Scan(
		&user.Id,
		&user.Username,
		&user.Password,
		&user.Role,
		&user.Email,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("User not found")
		}
	} else if login.Password != user.Password {
		return nil, fmt.Errorf("invalid username or password")
	}

	return &user, nil
}
