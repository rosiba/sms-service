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
    sent_at TIMESTAMP WITH TIME ZONE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE
)`
)

func InitDB(conn *pgxpool.Pool) error {
	_, err := conn.Exec(context.Background(), initSQL)
	if err != nil {
		return fmt.Errorf("failed to run initialization statement: %w ", err)
	}

	return nil
}
