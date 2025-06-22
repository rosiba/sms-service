package postgres

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/jackc/pgx/v5/pgxpool"
	"sms-service/internal/model"
	"time"
)

type PostgresRepository struct {
	conn *pgxpool.Pool
}

func NewMessageRepository(conn *pgxpool.Pool) *PostgresRepository {
	return &PostgresRepository{
		conn: conn,
	}
}

func (p *PostgresRepository) GetPendingMessages(count uint) ([]model.Message, error) {
	ctx := context.Background()
	rows, err := p.conn.Query(ctx, "SELECT id, recipient, content FROM messages WHERE status = $1 ORDER BY id LIMIT $2", model.MessageStatusPending, count)
	if err != nil {
		return nil, fmt.Errorf("failed to query: %w", err)
	}
	defer rows.Close()

	messages := make([]model.Message, 0)
	for rows.Next() {
		var m model.Message
		if err := rows.Scan(&m.ID, &m.Recipient, &m.Content); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		messages = append(messages, m)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to query: %w", err)
	}

	return messages, nil
}

func (p *PostgresRepository) GetSentMessages(count uint) ([]model.Message, error) {
	ctx := context.Background()
	rows, err := p.conn.Query(ctx, "SELECT id, recipient, content, sent_at FROM messages WHERE deleted_at IS NULL AND status = $1 ORDER BY created_at DESC LIMIT $2", model.MessageStatusSent, count)
	if err != nil {
		return nil, fmt.Errorf("failed to query rows: %v", err)
	}
	defer rows.Close()

	messages := make([]model.Message, 0)
	for rows.Next() {
		var m model.Message
		if err := rows.Scan(&m.ID, &m.Recipient, &m.Content, &m.SentAt); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}
		messages = append(messages, m)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed to scan rows: %v", err)
	}

	return messages, nil
}

func (p *PostgresRepository) SaveMessage(message model.Message) (string, error) {
	ctx := context.Background()
	tx, err := p.conn.Begin(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to start transaction: %w", err)
	}

	var insertedID string
	err = tx.QueryRow(ctx, "INSERT INTO messages(recipient, content, status) values($1, $2, $3) RETURNING id", message.Recipient, message.Content, model.MessageStatusPending).Scan(&insertedID)
	if err != nil {
		return "", fmt.Errorf("failed to insert message: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		if err := tx.Rollback(ctx); err != nil {
			return "", fmt.Errorf("failed to rollback transaction: %w", err)
		}
		return "", fmt.Errorf("failed to commit transaction(rollback applied): %w", err)
	}

	return insertedID, nil
}

func (p *PostgresRepository) SetMessageAsSent(id pgtype.UUID, sentAt time.Time) error {
	ctx := context.Background()
	tx, err := p.conn.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to start transaction: %w", err)
	}
	t, err := tx.Exec(ctx, "UPDATE messages SET status = $1, sent_at = $2 WHERE id = $3", model.MessageStatusSent, sentAt, id)
	if err != nil {
		if err := tx.Rollback(ctx); err != nil {
			return fmt.Errorf("failed to rollback transaction: %w", err)
		}
		return fmt.Errorf("failed to set status: %w", err)
	}

	if err := tx.Commit(ctx); err != nil {
		if err := tx.Rollback(ctx); err != nil {
			return fmt.Errorf("failed to rollback transaction: %w", err)
		}
		return fmt.Errorf("failed to commit transaction(rollback applied): %w", err)
	}

	if t.RowsAffected() == 0 {
		return fmt.Errorf("no lines affected")
	}

	return nil
}
