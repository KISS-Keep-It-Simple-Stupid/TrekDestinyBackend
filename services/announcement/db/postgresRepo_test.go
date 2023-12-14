package db

import (
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

	announcementID := 1
	hostID := 123
	firstname := "John"
	lastname := "Doe"
	username := "johndoe"
	rows := sqlmock.NewRows([]string{"host_id"}).AddRow(hostID)
	mock.ExpectQuery(`select host_id from announcement_offer where announcement_id = \$1`).
		WithArgs(announcementID).
		WillReturnRows(rows)

	rows = sqlmock.NewRows([]string{"firstname", "lastname", "username"}).
		AddRow(firstname, lastname, username)
	mock.ExpectQuery(`select firstname, lastname, username from members where id = \$1`).
		WithArgs(hostID).
		WillReturnRows(rows)

	response, err := repo.GetOfferDetails(announcementID)

	assert.NoError(t, err)

	expectedResponse := &pb.GetOfferResponse{
		Offers: []*pb.OfferRecord{
			{
				HostId:        int32(hostID),
				HostFirstName: firstname,
				HostLastName:  lastname,
				HostUsername:  username,
			},
		},
	}
	assert.Equal(t, expectedResponse, response)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestGetProfileAnnouncementDetails(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer mockDB.Close()

	repo := &PostgresRepository{DB: mockDB}

	userID := 1
	announcementID := 123
	description := "Test Announcement"
	startdate := time.Now()
	enddate := startdate.Add(7 * 24 * time.Hour)
	city := "Test City"
	state := "Test State"
	country := "Test Country"
	numTravelers := 2

	rows := sqlmock.NewRows([]string{
		"id", "user_id", "description", "startdate", "enddate", "city", "state", "country", "numberoftravelers",
	}).AddRow(
		announcementID, userID, description, startdate, enddate, city, state, country, numTravelers,
	)

	mock.ExpectQuery(`select id, user_id, description, startdate, enddate, city, state, country, numberoftravelers from announcement where user_id = \$1`).
		WithArgs(userID).
		WillReturnRows(rows)

	response, err := repo.GetProfileAnnouncementDetails(userID)

	assert.NoError(t, err)

	expectedResponse := &pb.GetCardProfileResponse{
		Cards: []*pb.CardRecord{
			{
				CardId:             int32(announcementID),
				UserId:             int32(userID),
				Description:        description,
				StartDate:          startdate.Format("2006-01-02"),
				EndDate:            enddate.Format("2006-01-02"),
				DestinationCity:    city,
				DestinationState:   state,
				DestinationCountry: country,
				NumberOfTravelers:  int32(numTravelers),
			},
		},
	}
	assert.Equal(t, expectedResponse, response)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestValidateOffer(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer mockDB.Close()

	repo := &PostgresRepository{DB: mockDB}

	announcementID := 1
	userID := 123
	ownAnnouncementRows := sqlmock.NewRows([]string{"user_id"}).AddRow(userID)
	mock.ExpectQuery(`select user_id from announcement where id = \$1`).
		WithArgs(announcementID).
		WillReturnRows(ownAnnouncementRows)

	canOffer, errorMessage, err := repo.ValidateOffer(announcementID, userID)

	assert.NoError(t, err)

	assert.False(t, canOffer)
	assert.Equal(t, "you can not offer to your own announcement", errorMessage)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
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

	err = repo.RejectUserAsHost(offerInfo)

	assert.NoError(t, err)

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
