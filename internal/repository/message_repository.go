package repository

import (
	"database/sql"
	"errors"
	"messageOK/internal/entity"
)

type MessageRepository interface {
	GetUnsentMessages() ([]entity.Message, error)
	MarkMessageAsSent(id int) error
}

type mysqlMessageRepository struct {
	DB *sql.DB
}

func NewMySQLMessageRepository(db *sql.DB) MessageRepository {
	return &mysqlMessageRepository{DB: db}
}

func (r *mysqlMessageRepository) GetUnsentMessages() ([]entity.Message, error) {
	rows, err := r.DB.Query("SELECT id, content, `to`, status FROM messages WHERE status = 'unsent' order by id limit 2")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var messages []entity.Message
	for rows.Next() {
		var msg entity.Message
		if err := rows.Scan(&msg.ID, &msg.Content, &msg.To, &msg.Status); err != nil {
			return nil, err
		}
		messages = append(messages, msg)
	}
	return messages, nil
}

func (r *mysqlMessageRepository) MarkMessageAsSent(id int) error {
	result, err := r.DB.Exec("UPDATE messages SET status = 'sent' WHERE id = ?", id)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.New("there is no record for update")
	}

	return nil
}
