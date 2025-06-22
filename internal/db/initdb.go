package db

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	initSQL = `CREATE TABLE IF NOT EXISTS messages(
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    recipient VARCHAR(25) NOT NULL,
    content VARCHAR(200) NOT NULL,
    status VARCHAR(25) NOT NULL,
    sent_at TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
)`
)

func InitDB(conn *pgxpool.Pool) error {
	_, err := conn.Exec(context.Background(), initSQL)
	if err != nil {
		return fmt.Errorf("failed to run initialization statement: %w ", err)
	}

	return nil
}
