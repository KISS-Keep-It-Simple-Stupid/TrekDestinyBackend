package db

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/chat/models"
)


func TestNewPostgresRepository(t *testing.T) {
	db, _, _ := sqlmock.New()
	repo := NewPostgresRepository(db)
	_ , ok :=  repo.(*PostgresRepository)
	assert.True(t, ok, "Expected a *PostgresRepository")
	assert.Equal(t, db, repo.(*PostgresRepository).DB, "DB field should be set to the provided database connection")
}
func TestInsertChatMessage(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database connection: %v", err)
	}
	defer db.Close()
	repo := &PostgresRepository{DB: db}
	message := &models.Message{
		Message:  "Hello, test!",
		UserID:   123,
		UserName: "testuser",
		Time:     "2001-10-10",
		ChatID:   "10-12",
	}

	mock.ExpectExec(`^insert into messages`).WithArgs(message.Message, message.UserID, message.UserName, message.Time, message.ChatID).WillReturnResult(sqlmock.NewResult(1, 1))

	repo.InsertChatMessage(message)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestGetCount(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database connection: %v", err)
	}
	defer db.Close()
	repo := &PostgresRepository{DB: db}
	chatID := "123"
	expectedCount := 42
	mock.ExpectQuery(`^select count`).WithArgs(chatID).WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(expectedCount))

	count := repo.GetCount(chatID)
	assert.Equal(t, expectedCount, count, "Unexpected count value")
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestGetHistory(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database connection: %v", err)
	}
	defer db.Close()
	viper.Set("PAGESIZE", "10")
	repo := &PostgresRepository{DB: db}
	chatID := "123"
	pageNumber := 1
	numberOfMessages := 30

	mock.ExpectQuery(`^SELECT  user_id , username , message , time`).WithArgs(chatID, 20, 30).WillReturnRows(sqlmock.NewRows([]string{"user_id", "username", "message", "time"}).
		AddRow(1, "user1", "message1", "2001-10-17").
		AddRow(2, "user2", "message2", "2001-10-17").
		AddRow(3, "user3", "message10", "2001-10-17"))

	history, err := repo.GetHistory(chatID, pageNumber, numberOfMessages)
	assert.NoError(t, err, "Unexpected error")
	assert.Len(t, history, 3, "Unexpected number of messages in history")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}
