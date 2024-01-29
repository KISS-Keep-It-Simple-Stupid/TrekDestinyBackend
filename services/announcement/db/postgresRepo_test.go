package db

import (
	"database/sql"
	"errors"
	"testing"
	"time"

	pb "github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/announcement/pb"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestGetIdFromUsername(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	repo := &PostgresRepository{DB: mockDB}

	expectedUsername := "testuser"
	expectedID := 123
	mock.ExpectQuery(`^select id from members where username = \$1$`).
		WithArgs(expectedUsername).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedID))

	id, err := repo.GetIdFromUsername(expectedUsername)

	assert.NoError(t, err)
	assert.Equal(t, expectedID, id)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetUsernameFromId(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	repo := &PostgresRepository{DB: mockDB}

	expectedID := 123
	expectedUsername := "testuser"
	mock.ExpectQuery(`^select username from members where id = \$1$`).
		WithArgs(expectedID).
		WillReturnRows(sqlmock.NewRows([]string{"username"}).AddRow(expectedUsername))

	username, err := repo.GetUsernameFromId(expectedID)

	assert.NoError(t, err)
	assert.Equal(t, expectedUsername, username)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestInsertAnnouncement(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock: %v", err)
	}
	defer mockDB.Close()

	repo := &PostgresRepository{DB: mockDB}

	announcementInfo := &pb.CreateCardRequest{
		Description:        "Test Announcement",
		StartDate:          "2023-01-01",
		EndDate:            "2023-01-10",
		DestinationCity:    "Test City",
		DestinationState:   "Test State",
		DestinationCountry: "Test Country",
		NumberOfTravelers:  5,
	}
	userID := 1

	expectedID := 123

	mock.ExpectExec("insert into announcement").WithArgs(
		userID,
		announcementInfo.Description,
		time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC),
		time.Date(2023, 1, 10, 0, 0, 0, 0, time.UTC),
		announcementInfo.DestinationCity,
		announcementInfo.DestinationState,
		announcementInfo.DestinationCountry,
		announcementInfo.NumberOfTravelers,
	).WillReturnResult(sqlmock.NewResult(1, 1))

	rows := sqlmock.NewRows([]string{"id"}).AddRow(expectedID)
	mock.ExpectQuery("select id from announcement").WillReturnRows(rows)

	resultID, err := repo.InsertAnnouncement(announcementInfo, userID)

	assert.NoError(t, err, "Expected no error")
	assert.Equal(t, expectedID, resultID, "Expected announcement ID to match")

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestInsertAnnouncementLanguage(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer mockDB.Close()

	repo := &PostgresRepository{DB: mockDB}

	announcementID := 1
	language := "en"
	mock.ExpectExec(`insert into announcement_language \(announcement_id, language\) values \(\$1, \$2\)`).
		WithArgs(announcementID, language).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.InsertAnnouncementLanguage(announcementID, language)

	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestCheckAnnouncementTimeValidation(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer mockDB.Close()

	repo := &PostgresRepository{DB: mockDB}

	expectedUserID := 123
	expectedStartDate := "2023-01-01"
	expectedEndDate := "2023-01-05"
	parsedStartDate, err := time.Parse("2006-01-02", expectedStartDate)
	require.NoError(t, err)
	parsedEndDate, err := time.Parse("2006-01-02", expectedEndDate)
	require.NoError(t, err)

	mock.ExpectQuery(`^select id from announcement where user_id = \$1 and enddate >= \$2 and startdate <= \$3$`).
		WithArgs(expectedUserID, parsedStartDate, parsedEndDate).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	isValid, err := repo.CheckAnnouncementTimeValidation(expectedStartDate, expectedEndDate, expectedUserID)
	assert.NoError(t, err)
	assert.False(t, isValid)

	require.NoError(t, mock.ExpectationsWereMet())

	mock.ExpectationsWereMet()

	expectedStartDateInvalid := "2023-02-01"
	expectedEndDateInvalid := "2023-02-05"
	parsedStartDateInvalid, err := time.Parse("2006-01-02", expectedStartDateInvalid)
	require.NoError(t, err)
	parsedEndDateInvalid, err := time.Parse("2006-01-02", expectedEndDateInvalid)
	require.NoError(t, err)

	mock.ExpectQuery(`^select id from announcement where user_id = \$1 and enddate >= \$2 and startdate <= \$3$`).
		WithArgs(expectedUserID, parsedStartDateInvalid, parsedEndDateInvalid).
		WillReturnRows(sqlmock.NewRows([]string{"id"}))

	isValid, err = repo.CheckAnnouncementTimeValidation(expectedStartDateInvalid, expectedEndDateInvalid, expectedUserID)
	assert.NoError(t, err)
	assert.True(t, isValid)

	require.NoError(t, mock.ExpectationsWereMet())
}

func TestGetAnnouncementDetails(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer mockDB.Close()

	repo := &PostgresRepository{DB: mockDB}

	mock.ExpectQuery(`^select id, user_id, description, startdate, enddate, city, state, country, numberoftravelers from announcement where announcement_status = \$1(.+)$`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "description", "startdate", "enddate", "city", "state", "country", "numberoftravelers"}).
			AddRow(1, 2, "Description 1", time.Now(), time.Now().AddDate(0, 1, 0), "City 1", "State 1", "Country 1", 3).
			AddRow(2, 3, "Description 2", time.Now(), time.Now().AddDate(0, 2, 0), "City 2", "State 2", "Country 2", 4))

	mock.ExpectQuery(`^select id, user_id, description, startdate, enddate, city, state, country, numberoftravelers from announcement where announcement_status = \$1(.+)$`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "description", "startdate", "enddate", "city", "state", "country", "numberoftravelers"}).
			AddRow(1, 2, "Description 1", time.Now(), time.Now().AddDate(0, 1, 0), "City 1", "State 1", "Country 1", 3).
			AddRow(2, 3, "Description 2", time.Now(), time.Now().AddDate(0, 2, 0), "City 2", "State 2", "Country 2", 4))

	mock.ExpectQuery(`^select id, user_id, description, startdate, enddate, city, state, country, numberoftravelers from announcement where announcement_status = \$1(.+)$`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "description", "startdate", "enddate", "city", "state", "country", "numberoftravelers"}).
			AddRow(1, 2, "Description 1", time.Now(), time.Now().AddDate(0, 1, 0), "City 1", "State 1", "Country 1", 3).
			AddRow(2, 3, "Description 2", time.Now(), time.Now().AddDate(0, 2, 0), "City 2", "State 2", "Country 2", 4))

	_, _ = repo.GetAnnouncementDetails([]string{}, "", 10, 1)

	mock.ExpectQuery(`^select id, user_id, description, startdate, enddate, city, state, country, numberoftravelers from announcement where announcement_status = \$1(.+)$`).
		WithArgs(2).
		WillReturnError(err)

	_, err = repo.GetAnnouncementDetails([]string{}, "", 10, 1)
	if err == nil {
		t.Error("Expected error, but got nil")
	}

	mock.ExpectQuery(`^select id, user_id, description, startdate, enddate, city, state, country, numberoftravelers from announcement where announcement_status = \$1(.+)$`).
		WithArgs(3).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "description", "startdate", "enddate", "city", "state", "country", "numberoftravelers"}).
			AddRow("invalid", 2, "Description 1", time.Now(), time.Now().AddDate(0, 1, 0), "City 1", "State 1", "Country 1", 3))

	_, err = repo.GetAnnouncementDetails([]string{}, "", 10, 1)
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}

func TestGetLanguagesOfAnnouncement(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer mockDB.Close()

	repo := &PostgresRepository{DB: mockDB}

	announcementID := 1
	expectedLanguages := []string{"en", "fr", "es"}
	rows := sqlmock.NewRows([]string{"language"}).
		AddRow("en").
		AddRow("fr").
		AddRow("es")
	mock.ExpectQuery(`select language from announcement_language where announcement_id = \$1`).
		WithArgs(announcementID).
		WillReturnRows(rows)

	languages, err := repo.GetLanguagesOfAnnouncement(announcementID)

	assert.NoError(t, err)

	assert.Equal(t, expectedLanguages, languages)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestInsertOffer(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer mockDB.Close()

	repo := &PostgresRepository{DB: mockDB}

	offerInfo := &pb.CreateOfferRequest{
		AnnouncementId: 123,
	}
	userID := 456
	mock.ExpectExec(`insert into announcement_offer \(announcement_id, host_id\) values \(\$1, \$2\)`).
		WithArgs(offerInfo.AnnouncementId, userID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.InsertOffer(offerInfo, userID)

	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestGetGuestID(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer mockDB.Close()

	repo := &PostgresRepository{DB: mockDB}

	announcementID := 1
	expectedGuestID := 123
	rows := sqlmock.NewRows([]string{"user_id"}).AddRow(expectedGuestID)
	mock.ExpectQuery(`select user_id from announcement where id = \$1`).
		WithArgs(announcementID).
		WillReturnRows(rows)

	guestID, err := repo.GetGuestID(announcementID)

	assert.NoError(t, err)

	assert.Equal(t, expectedGuestID, guestID)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestGetOfferDetails(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer mockDB.Close()

	repo := &PostgresRepository{DB: mockDB}

	mock.ExpectQuery(`^select host_id , offer_status from announcement_offer where announcement_id = \$1$`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"host_id", "offer_status"}).
			AddRow(1, 2))

	mock.ExpectQuery(`^select firstname, lastname, username from members where id = \$1$`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"firstname", "lastname", "username"}).
			AddRow("John", "Doe", "johndoe"))

	response, err := repo.GetOfferDetails(1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(response.Offers) != 1 {
		t.Errorf("Expected 1 offer, but got %d", len(response.Offers))
	}

	mock.ExpectQuery(`^select host_id , offer_status from announcement_offer where announcement_id = \$1$`).
		WithArgs(2).
		WillReturnError(err)

	_, err = repo.GetOfferDetails(2)
	if err == nil {
		t.Error("Expected error, but got nil")
	}

	mock.ExpectQuery(`^select host_id , offer_status from announcement_offer where announcement_id = \$1$`).
		WithArgs(3).
		WillReturnRows(sqlmock.NewRows([]string{"host_id", "offer_status"}).
			AddRow("invalid", 2))

	_, err = repo.GetOfferDetails(3)
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}

func TestGetProfileAnnouncementDetails(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer mockDB.Close()

	repo := &PostgresRepository{DB: mockDB}

	mock.ExpectQuery(`^select id, user_id, description, startdate, enddate, city, state, country, numberoftravelers , announcement_status , main_host from announcement where user_id = \$1$`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "description", "startdate", "enddate", "city", "state", "country", "numberoftravelers", "announcement_status", "main_host"}).
			AddRow(1, 1, "Announcement 1", time.Now(), time.Now().Add(24*time.Hour), "City1", "State1", "Country1", 3, "active", true))

	_, _ = repo.GetProfileAnnouncementDetails(1)

	mock.ExpectQuery(`^select id, user_id, description, startdate, enddate, city, state, country, numberoftravelers , announcement_status , main_host from announcement where user_id = \$1$`).
		WithArgs(2).
		WillReturnError(err)

	_, err = repo.GetProfileAnnouncementDetails(2)
	if err == nil {
		t.Error("Expected error, but got nil")
	}

	mock.ExpectQuery(`^select id, user_id, description, startdate, enddate, city, state, country, numberoftravelers , announcement_status , main_host from announcement where user_id = \$1$`).
		WithArgs(3).
		WillReturnRows(sqlmock.NewRows([]string{"id", "user_id", "description", "startdate", "enddate", "city", "state", "country", "numberoftravelers", "announcement_status", "main_host"}).
			AddRow(3, "invalid", "Announcement 3", time.Now(), time.Now().Add(24*time.Hour), "City3", "State3", "Country3", 2, "active", false))

	_, err = repo.GetProfileAnnouncementDetails(3)
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}

func TestValidateOffer(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer mockDB.Close()

	repo := &PostgresRepository{DB: mockDB}

	mock.ExpectQuery(`^select user_id from announcement where id = \$1$`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"user_id"}).AddRow(2))

	mock.ExpectQuery(`^select host_id from announcement_offer where announcement_id = \$1$`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"host_id"}))

	validation, message, err := repo.ValidateOffer(1, 3)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !validation {
		t.Error("Expected validation to be true, but got false")
	}
	if message != "" {
		t.Errorf("Expected message to be empty, but got: %s", message)
	}

	mock.ExpectQuery(`^select user_id from announcement where id = \$1$`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"user_id"}).AddRow(3))

	validation, _, err = repo.ValidateOffer(1, 3)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if validation {
		t.Error("Expected validation to be false, but got true")
	}

	mock.ExpectQuery(`^select user_id from announcement where id = \$1$`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"user_id"}).AddRow(2))

	mock.ExpectQuery(`^select host_id from announcement_offer where announcement_id = \$1$`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"host_id"}).AddRow(3))

	validation, message, err = repo.ValidateOffer(1, 3)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if validation {
		t.Error("Expected validation to be false, but got true")
	}
	if message != "you have already offered to this announcement" {
		t.Errorf("Expected message to be 'you have already offered to this announcement', but got: %s", message)
	}
}

func TestInsertPost(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer mockDB.Close()

	repo := &PostgresRepository{DB: mockDB}

	mock.ExpectQuery(`^select user_id, main_host from announcement where id = \$1$`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"user_id", "main_host"}).AddRow(2, 3))

	mock.ExpectExec(`^insert into post \(announcement_id, host_id, guest_id, title, rating, body\) values \(\$1, \$2, \$3, \$4, \$5, \$6\)$`).
		WithArgs(1, 3, 2, "Post Title", 4, "Post Body").
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.InsertPost(&pb.CreatePostRequest{AnnouncementId: 1, PostTitle: "Post Title", HostRating: 4, PostBody: "Post Body"})
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	mock.ExpectQuery(`^select user_id, main_host from announcement where id = \$1$`).
		WithArgs(1).
		WillReturnError(errors.New("database error"))

	err = repo.InsertPost(&pb.CreatePostRequest{AnnouncementId: 1, PostTitle: "Post Title", HostRating: 4, PostBody: "Post Body"})
	if err == nil {
		t.Error("Expected error, but got nil")
	}

	mock.ExpectQuery(`^select user_id, main_host from announcement where id = \$1$`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"user_id", "main_host"}).AddRow(2, 3))

	mock.ExpectExec(`^insert into post \(announcement_id, host_id, guest_id, title, rating, body\) values \(\$1, \$2, \$3, \$4, \$5, \$6\)$`).
		WithArgs(1, 3, 2, "Post Title", 4, "Post Body").
		WillReturnError(errors.New("database error"))

	err = repo.InsertPost(&pb.CreatePostRequest{AnnouncementId: 1, PostTitle: "Post Title", HostRating: 4, PostBody: "Post Body"})
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}

func TestGetLastPostId(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer mockDB.Close()

	repo := &PostgresRepository{DB: mockDB}

	mock.ExpectQuery(`^select id from post order by id desc limit 1$`).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(42))

	lastPostID, err := repo.GetLastPostId()
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if lastPostID != 42 {
		t.Errorf("Expected lastPostID to be 42, but got %d", lastPostID)
	}

	mock.ExpectQuery(`^select id from post order by id desc limit 1$`).
		WillReturnError(errors.New("database error"))

	_, err = repo.GetLastPostId()
	if err == nil {
		t.Error("Expected error, but got nil")
	}

	mock.ExpectQuery(`^select id from post order by id desc limit 1$`).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow("invalid"))

	_, err = repo.GetLastPostId()
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}

func TestGetMyPostDetails(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer mockDB.Close()

	repo := &PostgresRepository{DB: mockDB}

	guestID := 123
	rows := sqlmock.NewRows([]string{"id", "announcement_id", "host_id", "guest_id", "title", "rating", "body"}).
		AddRow(1, 2, 3, 4, "Test Post", 4, "This is a test post body")
	mock.ExpectQuery(`select id, announcement_id, host_id, guest_id, title, rating, body from post where guest_id = \$1`).
		WithArgs(guestID).
		WillReturnRows(rows)

	hostID := 3
	guestIDDB := 4
	hostUsername := "host_user"
	guestUsername := "guest_user"
	rows = sqlmock.NewRows([]string{"username"}).AddRow(hostUsername)
	mock.ExpectQuery(`select username from members where id = \$1`).
		WithArgs(hostID).
		WillReturnRows(rows)

	rows = sqlmock.NewRows([]string{"username"}).AddRow(guestUsername)
	mock.ExpectQuery(`select username from members where id = \$1`).
		WithArgs(guestIDDB).
		WillReturnRows(rows)

	response, err := repo.GetMyPostDetails(guestID)

	assert.NoError(t, err)

	expectedResponse := &pb.GetMyPostResponse{
		Posts: []*pb.PostRecord{
			{
				PostId:         1,
				AnnouncementId: 2,
				HostId:         3,
				GuestId:        4,
				PostTitle:      "Test Post",
				HostRating:     4,
				PostBody:       "This is a test post body",
				HostUsername:   hostUsername,
				GuestUsername:  guestUsername,
			},
		},
	}
	assert.Equal(t, expectedResponse, response)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestGetPostHostDetails(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock DB: %v", err)
	}
	defer db.Close()

	repo := &PostgresRepository{DB: db}

	username := "testuser"
	hostID := 1
	guestID := 2

	mock.ExpectQuery("SELECT id FROM members WHERE username = ?").
		WithArgs(username).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(123))

	mock.ExpectQuery("SELECT id, announcement_id, host_id, guest_id, title, rating, body FROM post WHERE host_id = ? OR guest_id= ?").
		WithArgs(123).
		WillReturnRows(sqlmock.NewRows([]string{"id", "announcement_id", "host_id", "guest_id", "title", "rating", "body"}).
			AddRow(1, 1, hostID, guestID, "Title1", 4.5, "Body1").
			AddRow(2, 2, hostID, guestID, "Title2", 3.5, "Body2"))

	mock.ExpectQuery("SELECT username FROM members WHERE id = ?").
		WithArgs(hostID).
		WillReturnRows(sqlmock.NewRows([]string{"username"}).AddRow("HostUsername"))
	mock.ExpectQuery("SELECT username FROM members WHERE id = ?").
		WithArgs(guestID).
		WillReturnRows(sqlmock.NewRows([]string{"username"}).AddRow("GuestUsername"))

	_, _ = repo.GetPostHostDetails(username)
}

func TestAcceptUserAsHost(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer mockDB.Close()

	repo := &PostgresRepository{DB: mockDB}

	offerInfo := &pb.AcceptOfferRequest{
		HostId:         123,
		AnnouncementId: 456,
	}
	mock.ExpectExec(`update announcement set main_host = \$1, status = 'Accepted' where id = \$2`).
		WithArgs(offerInfo.HostId, offerInfo.AnnouncementId).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.AcceptUserAsHost(offerInfo)

	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestRejectUserAsHost(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer mockDB.Close()

	repo := &PostgresRepository{DB: mockDB}

	offerInfo := &pb.RejectOfferRequest{
		HostId:         123,
		AnnouncementId: 456,
	}
	mock.ExpectExec(`delete from announcement_offer where announcement_id = \$1 and host_id = \$2`).
		WithArgs(offerInfo.AnnouncementId, offerInfo.HostId).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.RejectUserOffer(offerInfo)

	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestUpdateAnnouncementInformation(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer mockDB.Close()

	repo := &PostgresRepository{DB: mockDB}

	announcementInfo := &pb.EditAnnouncementRequest{
		Description:         "Updated Description",
		StartDate:           "2022-01-01",
		EndDate:             "2022-02-01",
		DestinationCity:     "Updated City",
		DestinationState:    "Updated State",
		DestinationCountry:  "Updated Country",
		NumberOfTravelers:   5,
		CardId:              1,
	}

	mock.ExpectExec(`^update announcement set description = COALESCE\(NULLIF\(\$1, ''\), description\), startdate = \$2, enddate = \$3, city = COALESCE\(NULLIF\(\$4, ''\), city\), state = COALESCE\(NULLIF\(\$5, ''\), state\), country = COALESCE\(NULLIF\(\$6, ''\), country\), numberoftravelers = COALESCE\(NULLIF\(\$7, 0\), numberoftravelers\) where id = \$8$`).
		WithArgs(announcementInfo.Description, "2022-01-01", "2022-02-01", announcementInfo.DestinationCity, announcementInfo.DestinationState, announcementInfo.DestinationCountry, announcementInfo.NumberOfTravelers, announcementInfo.CardId).
		WillReturnResult(sqlmock.NewResult(0, 1))

	mock.ExpectExec(`^delete from announcement_language where announcement_id = \$1$`).
		WithArgs(announcementInfo.CardId).
		WillReturnResult(sqlmock.NewResult(0, 1))

	_ = repo.UpdateAnnouncementInformation(announcementInfo)

	mock.ExpectExec(`^update announcement set description = COALESCE\(NULLIF\(\$1, ''\), description\), startdate = \$2, enddate = \$3, city = COALESCE\(NULLIF\(\$4, ''\), city\), state = COALESCE\(NULLIF\(\$5, ''\), state\), country = COALESCE\(NULLIF\(\$6, ''\), country\), numberoftravelers = COALESCE\(NULLIF\(\$7, 0\), numberoftravelers\) where id = \$8$`).
		WithArgs(announcementInfo.Description, "2022-01-01", "2022-02-01", announcementInfo.DestinationCity, announcementInfo.DestinationState, announcementInfo.DestinationCountry, announcementInfo.NumberOfTravelers, announcementInfo.CardId).
		WillReturnError(errors.New("database error"))

	err = repo.UpdateAnnouncementInformation(announcementInfo)
	if err == nil {
		t.Error("Expected error, but got nil")
	}

	mock.ExpectExec(`^update announcement set description = COALESCE\(NULLIF\(\$1, ''\), description\), startdate = \$2, enddate = \$3, city = COALESCE\(NULLIF\(\$4, ''\), city\), state = COALESCE\(NULLIF\(\$5, ''\), state\), country = COALESCE\(NULLIF\(\$6, ''\), country\), numberoftravelers = COALESCE\(NULLIF\(\$7, 0\), numberoftravelers\) where id = \$8$`).
		WithArgs(announcementInfo.Description, "2022-01-01", "2022-02-01", announcementInfo.DestinationCity, announcementInfo.DestinationState, announcementInfo.DestinationCountry, announcementInfo.NumberOfTravelers, announcementInfo.CardId).
		WillReturnResult(sqlmock.NewResult(0, 1))

	mock.ExpectExec(`^delete from announcement_language where announcement_id = \$1$`).
		WithArgs(announcementInfo.CardId).
		WillReturnError(errors.New("database error"))

	err = repo.UpdateAnnouncementInformation(announcementInfo)
	if err == nil {
		t.Error("Expected error, but got nil")
	}

	announcementInfo.StartDate = "invalid date"
	err = repo.UpdateAnnouncementInformation(announcementInfo)
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}

func TestDeleteAnnouncement(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer mockDB.Close()

	repo := &PostgresRepository{DB: mockDB}

	announcementID := 1

	mock.ExpectExec(`^delete from announcement_offer where announcement_id = \$1$`).
		WithArgs(announcementID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(`^delete from chatlist where announcement_id = \$1$`).
		WithArgs(announcementID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(`^delete from announcement where id = \$1$`).
		WithArgs(announcementID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.DeleteAnnouncement(announcementID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	mock.ExpectExec(`^delete from announcement_offer where announcement_id = \$1$`).
		WithArgs(announcementID).
		WillReturnError(errors.New("database error"))

	err = repo.DeleteAnnouncement(announcementID)
	if err == nil {
		t.Error("Expected error, but got nil")
	}

	mock.ExpectExec(`^delete from announcement_offer where announcement_id = \$1$`).
		WithArgs(announcementID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(`^delete from chatlist where announcement_id = \$1$`).
		WithArgs(announcementID).
		WillReturnError(errors.New("database error"))

	err = repo.DeleteAnnouncement(announcementID)
	if err == nil {
		t.Error("Expected error, but got nil")
	}

	mock.ExpectExec(`^delete from announcement_offer where announcement_id = \$1$`).
		WithArgs(announcementID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(`^delete from chatlist where announcement_id = \$1$`).
		WithArgs(announcementID).
		WillReturnResult(sqlmock.NewResult(0, 1))
	mock.ExpectExec(`^delete from announcement where id = \$1$`).
		WithArgs(announcementID).
		WillReturnError(errors.New("database error"))

	err = repo.DeleteAnnouncement(announcementID)
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}

func TestUpdatePostInformation(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer mockDB.Close()

	repo := &PostgresRepository{DB: mockDB}

	postInfo := &pb.EditPostRequest{
		PostTitle: "New Title",
		HostRating: 5,
		PostBody:   "Updated Body",
		PostId:     1,
	}
	mock.ExpectExec(`^update post set title = COALESCE\(NULLIF\(\$1, ''\), title\), rating = COALESCE\(NULLIF\(\$2, 0\), rating\), body = COALESCE\(NULLIF\(\$3, ''\), body\) where id = \$4$`).
		WithArgs(postInfo.PostTitle, postInfo.HostRating, postInfo.PostBody, postInfo.PostId).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.UpdatePostInformation(postInfo)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	mock.ExpectExec(`^update post set title = COALESCE\(NULLIF\(\$1, ''\), title\), rating = COALESCE\(NULLIF\(\$2, 0\), rating\), body = COALESCE\(NULLIF\(\$3, ''\), body\) where id = \$4$`).
		WithArgs(postInfo.PostTitle, postInfo.HostRating, postInfo.PostBody, postInfo.PostId).
		WillReturnError(errors.New("database error"))

	err = repo.UpdatePostInformation(postInfo)
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}

func TestGetHostId(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer mockDB.Close()

	repo := &PostgresRepository{DB: mockDB}

	announcementID := 1
	hostID := 42
	rows := sqlmock.NewRows([]string{"main_host"}).
		AddRow(hostID)
	mock.ExpectQuery(`^select main_host from announcement where id = \$1$`).
		WithArgs(announcementID).
		WillReturnRows(rows)

	result, err := repo.GetHostId(announcementID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if result != hostID {
		t.Errorf("Unexpected host ID. Expected: %d, Got: %d", hostID, result)
	}

	mock.ExpectQuery(`^select main_host from announcement where id = \$1$`).
		WithArgs(announcementID).
		WillReturnError(sql.ErrNoRows)

	_, err = repo.GetHostId(announcementID)
	if err != sql.ErrNoRows {
		t.Errorf("Expected sql.ErrNoRows, but got: %v", err)
	}

	mock.ExpectQuery(`^select main_host from announcement where id = \$1$`).
		WithArgs(announcementID).
		WillReturnError(errors.New("database error"))

	_, err = repo.GetHostId(announcementID)
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}

func TestUpdateHostImagesCount(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer mockDB.Close()

	repo := &PostgresRepository{DB: mockDB}

	userID := 1
	imageCount := 5
	mock.ExpectExec(`^update members set hostImageCount=\$1 where id=\$2$`).
		WithArgs(imageCount, userID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.UpdateHostImagesCount(userID, imageCount)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	mock.ExpectExec(`^update members set hostImageCount=\$1 where id=\$2$`).
		WithArgs(imageCount, userID).
		WillReturnError(errors.New("database error"))

	err = repo.UpdateHostImagesCount(userID, imageCount)
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}

func TestDeleteUserChatList(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer mockDB.Close()

	repo := &PostgresRepository{DB: mockDB}

	announcementID := 1
	hostID := 2
	mock.ExpectExec(`^delete from chatlist where announcement_id = \$1 and host_id = \$2$`).
		WithArgs(announcementID, hostID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.DeleteUserChatList(announcementID, hostID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	mock.ExpectExec(`^delete from chatlist where announcement_id = \$1 and host_id = \$2$`).
		WithArgs(announcementID, hostID).
		WillReturnError(errors.New("database error"))

	err = repo.DeleteUserChatList(announcementID, hostID)
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}

func TestUpdateChatListStatus(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer mockDB.Close()

	repo := &PostgresRepository{DB: mockDB}

	announcementID := 1
	hostID := 2
	mock.ExpectExec(`^update chatlist set chat_status = \$1 where announcement_id = \$2 and host_id != \$3$`).
		WithArgs(2, announcementID, hostID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	repo.UpdateChatListStatus(announcementID, hostID)
	mock.ExpectExec(`^update chatlist set chat_status = \$1 where announcement_id = \$2 and host_id != \$3$`).
		WithArgs(2, announcementID, hostID).
		WillReturnError(errors.New("database error"))

	repo.UpdateChatListStatus(announcementID, hostID)
}

func TestUpdateAnnouncementStatus(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer mockDB.Close()

	repo := &PostgresRepository{DB: mockDB}

	announcementID := 1
	hostID := 2
	mock.ExpectExec(`^update announcement set announcement_status=\$1, main_host = \$2 where id = \$3$`).
		WithArgs(2, hostID, announcementID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	_ = repo.UpdateAnnouncementStatus(announcementID, hostID)

	mock.ExpectExec(`^update announcement set announcement_status=\$1, main_host = \$2 where id = \$3$`).
		WithArgs(2, hostID, announcementID).
		WillReturnError(errors.New("database error"))

	err = repo.UpdateAnnouncementStatus(announcementID, hostID)
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}

func TestUpdateMainHostStatusInChatList(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock DB: %v", err)
	}
	defer db.Close()

	repo := &PostgresRepository{DB: db}

	announcementID := 1
	hostID := 2

	mock.ExpectExec("UPDATE chatlist SET chat_status = ? WHERE announcement_id = ? AND host_id = ?").
		WithArgs(2, announcementID, hostID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	_ = repo.UpdateMainHostStatusInChatList(announcementID, hostID)
}