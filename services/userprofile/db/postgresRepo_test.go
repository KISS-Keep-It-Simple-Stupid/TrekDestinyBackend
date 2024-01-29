package db

import (
	"database/sql"
	"errors"
	"log"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/userprofile/pb"
	"github.com/aws/aws-sdk-go/service/s3"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func TestPostgresRepository_NewPostgresRepository(t *testing.T) {
	mockDB, err := sql.Open("sqlmock", "mock")
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer mockDB.Close()

	repo := NewPostgresRepository(mockDB)

	if _, ok := repo.(*PostgresRepository); !ok {
		t.Error("Returned repository has the wrong type")
	}

	_ = repo.(*PostgresRepository).DB
}

func TestPostgresRepository_GetUserDetails(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to initialize SQL mock: %v", err)
	}
	defer db.Close()

	repo := &PostgresRepository{DB: db}

	expectedUsername := "testuser"
	expectedRows := sqlmock.NewRows([]string{
		"email", "username", "firstname", "lastname", "birthdate", "state", "country", "gender",
		"joiningdate", "ishost", "bio", "city", "address", "phonenumber", "ispetfirendly",
		"iskidfiendly", "issmokingallowed", "roomnumber",
	}).AddRow(
		"test@example.com", expectedUsername, "John", "Doe", time.Now(), "State", "Country", 1,
		time.Now(), "true", "Bio", "City", "Address", "1234567890", "Yes", "Yes", "No", 5,
	)

	mock.ExpectQuery("select").WithArgs(expectedUsername).WillReturnRows(expectedRows)

	result, err := repo.GetUserDetails(expectedUsername)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	expectedBirthDate := time.Now().Format("2006-01-02")
	expectedJoiningDate := time.Now().Format("2006-01-02")

	expectedResult := &pb.ProfileDetailsResponse{
		Email:            "test@example.com",
		UserName:         expectedUsername,
		FirstName:        "John",
		LastName:         "Doe",
		BirthDate:        expectedBirthDate,
		State:            "State",
		Country:          "Country",
		Gender:           1,
		JoiningDate:      expectedJoiningDate,
		IsHost:           "true",
		Bio:              "Bio",
		City:             "City",
		Address:          "Address",
		PhoneNumber:      "1234567890",
		IsPetFriendly:    "Yes",
		IsKidFriendly:    "Yes",
		IsSmokingAllowed: "No",
		RoomNumber:       5,
	}

	if !compareProfileDetails(result, expectedResult) {
		t.Errorf("Unexpected result. Expected: %+v, Got: %+v", expectedResult, result)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func compareProfileDetails(a, b *pb.ProfileDetailsResponse) bool {
	return a.Email == b.Email &&
		a.UserName == b.UserName &&
		a.FirstName == b.FirstName &&
		a.LastName == b.LastName &&
		a.BirthDate == b.BirthDate &&
		a.State == b.State &&
		a.Country == b.Country &&
		a.Gender == b.Gender &&
		a.JoiningDate == b.JoiningDate &&
		a.IsHost == b.IsHost &&
		a.Bio == b.Bio &&
		a.City == b.City &&
		a.Address == b.Address &&
		a.PhoneNumber == b.PhoneNumber &&
		a.IsPetFriendly == b.IsPetFriendly &&
		a.IsKidFriendly == b.IsKidFriendly &&
		a.IsSmokingAllowed == b.IsSmokingAllowed &&
		a.RoomNumber == b.RoomNumber
}

func TestPostgresRepository_UpdateUserInformation(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to initialize SQL mock: %v", err)
	}
	defer db.Close()

	repo := &PostgresRepository{DB: db}

	expectedUsername := "testuser"
	userInfo := &pb.EditProfileRequest{
		FirstName:        "John",
		LastName:         "Doe",
		City:             "New City",
		State:            "New State",
		Country:          "New Country",
		Bio:              "New Bio",
		IsHost:           "true",
		Address:          "New Address",
		IsPetFriendly:    "Yes",
		IsKidFriendly:    "No",
		IsSmokingAllowed: "Yes",
		PhoneNumber:      "1234567890",
		RoomNumber:       42,
	}

	mock.ExpectExec("update members").WithArgs(
		sqlmock.AnyArg(),
		userInfo.FirstName,
		userInfo.LastName,
		userInfo.City,
		userInfo.State,
		userInfo.Country,
		userInfo.Bio,
		userInfo.IsHost,
		userInfo.Address,
		userInfo.IsPetFriendly,
		userInfo.IsKidFriendly,
		userInfo.IsSmokingAllowed,
		userInfo.PhoneNumber,
		userInfo.RoomNumber,
		expectedUsername,
	).WillReturnResult(sqlmock.NewResult(0, 1))

	userInfo.CurrentPassword = "current_password"
	userInfo.NewPassword = "new_password"

	mock.ExpectExec("update members").WithArgs(
		sqlmock.AnyArg(),
		userInfo.FirstName,
		userInfo.LastName,
		userInfo.City,
		userInfo.State,
		userInfo.Country,
		userInfo.Bio,
		userInfo.IsHost,
		userInfo.Address,
		userInfo.IsPetFriendly,
		userInfo.IsKidFriendly,
		userInfo.IsSmokingAllowed,
		userInfo.PhoneNumber,
		userInfo.RoomNumber,
		expectedUsername,
	).WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.UpdateUserInformation(expectedUsername, userInfo)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}
}

func TestPostgresRepository_CheckUserExistance(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to initialize SQL mock: %v", err)
	}
	defer db.Close()

	repo := &PostgresRepository{DB: db}

	expectedUsername := "testuser"

	mock.ExpectQuery("select id from members where username = ?").
		WithArgs(expectedUsername).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	result1, err1 := repo.CheckUserExistance(expectedUsername)
	assert.True(t, result1, "Expected user to exist")
	assert.Nil(t, err1, "Unexpected error")

	mock.ExpectQuery("select id from members where username = ?").
		WithArgs(expectedUsername).
		WillReturnError(sql.ErrNoRows)

	result2, err2 := repo.CheckUserExistance(expectedUsername)
	assert.False(t, result2, "Expected user not to exist")
	assert.Nil(t, err2, "Unexpected error")

	mock.ExpectQuery("select id from members where username = ?").
		WithArgs(expectedUsername).
		WillReturnError(sql.ErrConnDone)

	result3, err3 := repo.CheckUserExistance(expectedUsername)
	assert.False(t, result3, "Expected user not to exist due to database error")
	assert.EqualError(t, err3, sql.ErrConnDone.Error(), "Unexpected error")

	assert.Nil(t, mock.ExpectationsWereMet(), "Unfulfilled expectations")
}

func TestPostgresRepository_GetUserPassword(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to initialize SQL mock: %v", err)
	}
	defer db.Close()

	repo := &PostgresRepository{DB: db}

	existingUserName := "existing_user"
	expectedPassword := "hashed_password"
	mock.ExpectQuery("select password from members").WithArgs(existingUserName).WillReturnRows(sqlmock.NewRows([]string{"password"}).AddRow(expectedPassword))

	nonExistingUserName := "non_existing_user"
	mock.ExpectQuery("select password from members").WithArgs(nonExistingUserName).WillReturnError(sql.ErrNoRows)

	password, err := repo.GetUserPassword(existingUserName)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if password != expectedPassword {
		t.Errorf("Expected password: %s, Got: %s", expectedPassword, password)
	}

	password, err = repo.GetUserPassword(nonExistingUserName)

	if err != sql.ErrNoRows {
		t.Errorf("Expected sql.ErrNoRows, Got: %v", err)
	}

	if password != "" {
		t.Errorf("Expected empty string for password, Got: %s", password)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestPostgresRepository_GetPublicProfile(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to initialize SQL mock: %v", err)
	}
	defer db.Close()

	repo := &PostgresRepository{DB: db}

	expectedUsername := "testuser"
	expectedRows := sqlmock.NewRows([]string{
		"id", "email", "username", "firstname", "lastname", "birthdate", "state", "country", "gender", "joiningdate", "bio", "city",
	}).AddRow(
		1, "test@example.com", expectedUsername, "John", "Doe", time.Now(), "State", "Country", 1, time.Now(), "Bio", "City",
	)

	mock.ExpectQuery("select id, email, username, firstname, lastname, birthdate, state, country, gender, joiningdate, COALESCE(NULLIF(bio, NULL), '') as bio, COALESCE(NULLIF(city, NULL), '') as city from members where username = ?").
		WithArgs(expectedUsername).
		WillReturnRows(expectedRows)

	_, _, _ = repo.GetPublicProfile(expectedUsername)

	expectedBirthDate := time.Now().Format("2006-01-02")
	expectedJoiningDate := time.Now().Format("2006-01-02")

	_ = &pb.PublicProfileResponse{
		Email:       "test@example.com",
		UserName:    expectedUsername,
		FirstName:   "John",
		LastName:    "Doe",
		BirthDate:   expectedBirthDate,
		State:       "State",
		Country:     "Country",
		Gender:      1,
		JoiningDate: expectedJoiningDate,
		Bio:         "Bio",
		City:        "City",
	}
}

func TestPostgresRepository_GetPublicProfileHost(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to initialize SQL mock: %v", err)
	}
	defer db.Close()

	repo := &PostgresRepository{DB: db}

	expectedGuestID := 1
	expectedHostUsername := "testhost"
	expectedHostID := 2
	expectedRows := sqlmock.NewRows([]string{
		"email", "username", "firstname", "lastname", "birthdate", "state", "country", "gender", "joiningdate",
		"ishost", "bio", "city", "address", "phonenumber", "ispetfirendly", "iskidfiendly", "issmokingallowed", "roomnumber",
	}).AddRow(
		"host@example.com", expectedHostUsername, "Host", "User", time.Now(), "State", "Country", 1, time.Now(),
		"true", "Host Bio", "Host City", "Host Address", "1234567890", "Yes", "Yes", "No", 5,
	)

	mock.ExpectQuery("select id from members where username = ?").
		WithArgs(expectedHostUsername).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(expectedHostID))

	mock.ExpectQuery(".*").
		WillReturnRows(sqlmock.NewRows([]string{"announcement_id"}).AddRow(1))

	mock.ExpectQuery("select email, username, firstname, lastname, birthdate, state, country, gender, joiningdate, ishost, COALESCE(NULLIF(bio, NULL), '') as bio, COALESCE(NULLIF(city, NULL), '') as city, COALESCE(NULLIF(address, NULL), '') as address, COALESCE(NULLIF(phonenumber, NULL), '') as phonenumber, COALESCE(NULLIF(ispetfirendly::text, ''), '') as ispetfirendly, COALESCE(NULLIF(iskidfiendly::text, ''), '') as iskidfiendly, COALESCE(NULLIF(issmokingallowed::text, ''), '') as issmokingallowed, COALESCE(NULLIF(roomnumber, NULL), 0) as roomnumber from members where username = ?").
		WithArgs(expectedHostUsername).
		WillReturnRows(expectedRows)

	_, _, _ = repo.GetPublicProfileHost(expectedGuestID, expectedHostUsername)

	expectedBirthDate := time.Now().Format("2006-01-02")
	expectedJoiningDate := time.Now().Format("2006-01-02")

	_ = &pb.PublicProfileHostResponse{
		Email:            "host@example.com",
		UserName:         expectedHostUsername,
		FirstName:        "Host",
		LastName:         "User",
		BirthDate:        expectedBirthDate,
		State:            "State",
		Country:          "Country",
		Gender:           1,
		JoiningDate:      expectedJoiningDate,
		IsHost:           "true",
		Bio:              "Host Bio",
		City:             "Host City",
		Address:          "Host Address",
		PhoneNumber:      "1234567890",
		IsPetFriendly:    "Yes",
		IsKidFriendly:    "Yes",
		IsSmokingAllowed: "No",
		RoomNumber:       5,
	}
}

func TestPostgresRepository_CheckIfUserCanViewProfile(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to initialize SQL mock: %v", err)
	}
	defer db.Close()

	repo := &PostgresRepository{DB: db}

	expectedGuestID := 1
	expectedHostID := 2
	expectedRows := sqlmock.NewRows([]string{"announcement_id"}).AddRow(1)

	mock.ExpectQuery("select announcement_id from announcement_offer where host_id = ?").
		WithArgs(expectedHostID).
		WillReturnRows(expectedRows)

	mock.ExpectQuery("select user_id from announcement where id = ? and user_id = ?").
		WithArgs(1, expectedGuestID).
		WillReturnRows(sqlmock.NewRows([]string{"user_id"}).AddRow(expectedGuestID))

	_, _ = repo.CheckIfUserCanViewProfile(expectedGuestID, expectedHostID)
}

func TestPostgresRepository_GetIdFromUsername(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to initialize SQL mock: %v", err)
	}
	defer db.Close()

	repo := &PostgresRepository{DB: db}

	expectedUsername := "testuser"

	mock.ExpectQuery("select id from members where username = ?").
		WithArgs(expectedUsername).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	result1, err1 := repo.GetIdFromUsername(expectedUsername)
	assert.Equal(t, 1, result1, "Expected user ID to be 1")
	assert.Nil(t, err1, "Unexpected error")

	mock.ExpectQuery("select id from members where username = ?").
		WithArgs(expectedUsername).
		WillReturnError(sql.ErrNoRows)

	result2, err2 := repo.GetIdFromUsername(expectedUsername)
	assert.Equal(t, -1, result2, "Expected user ID to be -1")
	assert.EqualError(t, err2, sql.ErrNoRows.Error(), "Expected error to be sql.ErrNoRows")

	mock.ExpectQuery("select id from members where username = ?").
		WithArgs(expectedUsername).
		WillReturnError(sql.ErrConnDone)

	result3, err3 := repo.GetIdFromUsername(expectedUsername)
	assert.Equal(t, -1, result3, "Expected user ID to be -1")
	assert.EqualError(t, err3, sql.ErrConnDone.Error(), "Expected error to be sql.ErrConnDone")

	assert.Nil(t, mock.ExpectationsWereMet(), "Unfulfilled expectations")
}

func TestPostgresRepository_InsertChatList(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to initialize SQL mock: %v", err)
	}
	defer db.Close()

	repo := &PostgresRepository{DB: db}

	expectedHostID := 1
	expectedGuestID := 2
	expectedAnnouncementID := 3

	mock.ExpectExec("insert into chatlist").
		WithArgs(expectedHostID, expectedGuestID, expectedAnnouncementID).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.InsertChatList(expectedHostID, expectedGuestID, expectedAnnouncementID)

	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("Unfulfilled expectations: %s", err)
	}
}

func TestPostgresRepository_GetChatList(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to initialize SQL mock: %v", err)
	}
	defer db.Close()

	repo := &PostgresRepository{DB: db}

	expectedGuestID := 1
	expectedRows := sqlmock.NewRows([]string{"id", "host_id", "guest_id", "announcement_id", "chat_status"}).
		AddRow(1, 2, expectedGuestID, 3, 1).
		AddRow(2, expectedGuestID, 4, 5, 2)

	mock.ExpectQuery("select id ,host_id , guest_id, announcement_id , chat_status from chatlist where guest_id = $1 or host_id = $2").
		WithArgs(expectedGuestID, expectedGuestID).
		WillReturnRows(expectedRows)

	mockObj := new(s3.S3)

	_, err = repo.GetChatList(expectedGuestID, mockObj)
	log.Printf("%s", err)
}

func TestPostgresRepository_GetUserNameByID(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer mockDB.Close()

	repo := &PostgresRepository{DB: mockDB}

	userID := 1
	expectedUsername := "JohnDoe"
	rows := sqlmock.NewRows([]string{"username"}).AddRow(expectedUsername)
	mock.ExpectQuery("select username from members where id = \\$1").
		WithArgs(userID).
		WillReturnRows(rows)

	username, err := repo.GetUserNameByID(userID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if username != expectedUsername {
		t.Errorf("Unexpected username. Expected: %s, Got: %s", expectedUsername, username)
	}

	mock.ExpectQuery("select username from members where id = \\$1").
		WithArgs(userID).
		WillReturnError(errors.New("database error"))

	_, err = repo.GetUserNameByID(userID)
	if err == nil {
		t.Error("Expected error, but got nil")
	}

	mock.ExpectQuery("select username from members where id = \\$1").
		WithArgs(userID).
		WillReturnRows(sqlmock.NewRows([]string{"username"}).AddRow(123)) // invalid data type

	_, _ = repo.GetUserNameByID(userID)
}

func TestPostgresRepository_GetHouseImagesCount(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer mockDB.Close()

	repo := &PostgresRepository{DB: mockDB}

	userID := 1
	expectedCount := 42
	rows := sqlmock.NewRows([]string{"hostImageCount"}).AddRow(expectedCount)
	mock.ExpectQuery("select hostImageCount from members where id = ?").
		WithArgs(userID).
		WillReturnRows(rows)

	count, err := repo.GetHouseImagesCount(userID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if count != expectedCount {
		t.Errorf("Unexpected count. Expected: %d, Got: %d", expectedCount, count)
	}

	mock.ExpectQuery("select hostImageCount from members where id = ?").
		WithArgs(userID).
		WillReturnError(errors.New("database error"))

	_, err = repo.GetHouseImagesCount(userID)
	if err == nil {
		t.Error("Expected error, but got nil")
	}

	mock.ExpectQuery("select hostImageCount from members where id = ?").
		WithArgs(userID).
		WillReturnRows(sqlmock.NewRows([]string{"hostImageCount"}).AddRow("invalid"))

	_, err = repo.GetHouseImagesCount(userID)
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}

func TestPostgresRepository_UpdateOfferStatus(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer mockDB.Close()

	repo := &PostgresRepository{DB: mockDB}

	announcementID := 1
	hostID := 2
	mock.ExpectExec("update announcement_offer set offer_status=\\$1 where announcement_id = \\$2 and host_id = \\$3").
		WithArgs(2, announcementID, hostID).
		WillReturnResult(sqlmock.NewResult(0, 1))

	err = repo.UpdateOfferStatus(announcementID, hostID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	mock.ExpectExec("update announcement_offer set offer_status=\\$1 where announcement_id = \\$2 and host_id = \\$3").
		WithArgs(2, announcementID, hostID).
		WillReturnResult(sqlmock.NewResult(0, 0))

	_ = repo.UpdateOfferStatus(announcementID, hostID)

	mock.ExpectExec("update announcement_offer set offer_status=\\$1 where announcement_id = \\$2 and host_id = \\$3").
		WithArgs(2, announcementID, hostID).
		WillReturnError(errors.New("database error"))

	err = repo.UpdateOfferStatus(announcementID, hostID)
	if err == nil {
		t.Error("Expected error, but got nil")
	}
}

func TestPostgresRepository_InsertUserLanguage(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer mockDB.Close()

	repo := &PostgresRepository{DB: mockDB}

	userID := 1
	language := "English"
	mock.ExpectExec("insert into member_language \\(user_id, language\\) values \\(\\$1, \\$2\\)").
		WithArgs(userID, language).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.InsertUserLanguage(userID, language)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	mock.ExpectExec("insert into member_language \\(user_id, language\\) values \\(\\$1, \\$2\\)").
		WithArgs(userID, language).
		WillReturnError(errors.New("database error"))

	err = repo.InsertUserLanguage(userID, language)
	if err == nil {
		t.Error("Expected error, but got nil")
	}

	mock.ExpectExec("insert into member_language \\(user_id, language\\) values \\(\\$1, \\$2\\)").
		WithArgs(userID, language).
		WillReturnResult(sqlmock.NewResult(0, 0))

	err = repo.InsertUserLanguage(userID, language)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestPostgresRepository_InsertUserInterest(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer mockDB.Close()

	repo := &PostgresRepository{DB: mockDB}

	userID := 1
	interest := "Programming"
	mock.ExpectExec("insert into member_interest \\(user_id, interest\\) values \\(\\$1, \\$2\\)").
		WithArgs(userID, interest).
		WillReturnResult(sqlmock.NewResult(1, 1))

	err = repo.InsertUserInterest(userID, interest)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	mock.ExpectExec("insert into member_interest \\(user_id, interest\\) values \\(\\$1, \\$2\\)").
		WithArgs(userID, interest).
		WillReturnError(errors.New("database error"))

	err = repo.InsertUserInterest(userID, interest)
	if err == nil {
		t.Error("Expected error, but got nil")
	}

	mock.ExpectExec("insert into member_interest \\(user_id, interest\\) values \\(\\$1, \\$2\\)").
		WithArgs(userID, interest).
		WillReturnResult(sqlmock.NewResult(0, 0))

	err = repo.InsertUserInterest(userID, interest)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
}

func TestPostgresRepository_GetLanguagesOfUser(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer mockDB.Close()

	repo := &PostgresRepository{DB: mockDB}

	mock.ExpectQuery(`^select language from member_language where user_id = \$1$`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"language"}).
			AddRow("English").
			AddRow("Spanish"))

	languages, err := repo.GetLanguagesOfUser(1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(languages) != 2 {
		t.Errorf("Expected 2 languages, but got %d", len(languages))
	}

	mock.ExpectQuery(`^select language from member_language where user_id = \$1$`).
		WithArgs(2).
		WillReturnError(err)

	_, err = repo.GetLanguagesOfUser(2)
	if err == nil {
		t.Error("Expected error, but got nil")
	}

	mock.ExpectQuery(`^select language from member_language where user_id = \$1$`).
		WithArgs(3).
		WillReturnRows(sqlmock.NewRows([]string{"language"}).
			AddRow("English").
			AddRow("Invalid Language"))

	_, _ = repo.GetLanguagesOfUser(3)
}

func TestPostgresRepository_GetInterestsOfUser(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer mockDB.Close()

	repo := &PostgresRepository{DB: mockDB}

	mock.ExpectQuery(`^select interest from member_interest where user_id = \$1$`).
		WithArgs(1).
		WillReturnRows(sqlmock.NewRows([]string{"interest"}).
			AddRow("Music").
			AddRow("Travel"))

	interests, err := repo.GetInterestsOfUser(1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if len(interests) != 2 {
		t.Errorf("Expected 2 interests, but got %d", len(interests))
	}

	mock.ExpectQuery(`^select interest from member_interest where user_id = \$1$`).
		WithArgs(2).
		WillReturnError(err)

	_, err = repo.GetInterestsOfUser(2)
	if err == nil {
		t.Error("Expected error, but got nil")
	}

	mock.ExpectQuery(`^select interest from member_interest where user_id = \$1$`).
		WithArgs(3).
		WillReturnRows(sqlmock.NewRows([]string{"interest"}).
			AddRow("Music").
			AddRow("Invalid Interest"))

	_, _ = repo.GetInterestsOfUser(3)
}
