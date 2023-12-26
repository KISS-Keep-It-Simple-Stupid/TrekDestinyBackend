package dbrepo

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/notification/models"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	db, _, _ := sqlmock.New()
	repo := New(db)
	_ , ok :=  repo.(*postgresRepo)
	assert.True(t, ok, "Expected a *PostgresRepository")
}


func TestInsertNotification(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to open mock database: %v", err)
	}
	defer mockDB.Close()
	repo := &postgresRepo{db: mockDB}
	message := models.NotifMessage{
		UserID:  1,
		Message: "Test notification",
	}
	expectedQuery := `insert into notifications \(user_id,message\) values\(\$1,\$2\)`
	mock.ExpectExec(expectedQuery).
		WithArgs(message.UserID, message.Message).
		WillReturnResult(sqlmock.NewResult(1, 1)) 
	err = repo.InsertNotification(message)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
	assert.NoError(t, err, "Expected no error from InsertNotification")
}

func TestGetNotificationByID(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to open mock database: %v", err)
	}
	defer mockDB.Close()
	repo := &postgresRepo{db: mockDB}
	userID := 1
	expectedQuery := `select id , message from notifications where user_id = \$1`
	rows := sqlmock.NewRows([]string{"id", "message"}).
		AddRow(1, "Notification 1").
		AddRow(2, "Notification 2")

	mock.ExpectQuery(expectedQuery).
		WithArgs(userID).
		WillReturnRows(rows) 
	notifs, err := repo.GetNotificationByID(userID)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
	assert.NoError(t, err, "Expected no error from GetNotificationByID")
	expectedNotifs := []*models.NotifResponse{
		{ID: 1, Message: "Notification 1"},
		{ID: 2, Message: "Notification 2"},
	}
	assert.Equal(t, expectedNotifs, notifs, "Unexpected result from GetNotificationByID")

	repo.GetNotificationByID(10)
}

func TestDeleteNotification(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to open mock database: %v", err)
	}
	defer mockDB.Close()
	repo := &postgresRepo{db: mockDB}
	notificationID := 1
	expectedQuery := `delete from notifications where id = \$1`
	mock.ExpectExec(expectedQuery).
		WithArgs(notificationID).
		WillReturnResult(sqlmock.NewResult(0, 1)) 
	err = repo.DeleteNotification(notificationID)
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
	assert.NoError(t, err, "Expected no error from DeleteNotification")
}