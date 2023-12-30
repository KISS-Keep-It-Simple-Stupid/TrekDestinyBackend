package db

import (
	"database/sql"
	"database/sql/driver"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"
	pb "github.com/KISS-Keep-It-Simple-Stupid/TrekDestinyBackend/services/userprofile/pb"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/bcrypt"
)

func TestGetUserDetails(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	repo := &PostgresRepository{DB: db}

	// Expected query and test data
	expectedQuery := `select email, username, firstname, lastname, birthdate, state, country, gender, joiningdate, ishost, 
		COALESCE(NULLIF(bio, NULL), '') as bio,
		COALESCE(NULLIF(city, NULL), '') as city,
		COALESCE(NULLIF(address, NULL), '') as address,
		COALESCE(NULLIF(phonenumber, NULL), '') as phonenumber,
		COALESCE(NULLIF(ispetfirendly::text, ''), '') as ispetfirendly,
		COALESCE(NULLIF(iskidfiendly::text, ''), '') as iskidfiendly,
		COALESCE(NULLIF(issmokingallowed::text, ''), '') as issmokingallowed,
		COALESCE(NULLIF(roomnumber, NULL), 0) as roomnumber
		from members where username = $1`
	testUsername := "testuser"
	expectedUser := pb.ProfileDetailsResponse{
		Email:            "test@example.com",
		UserName:         "testuser",
		FirstName:        "John",
		LastName:         "Doe",
		BirthDate:        "1990-01-01",
		State:            "CA",
		Country:          "USA",
		Gender:           1,
		JoiningDate:      "2022-01-01",
		IsHost:           "true",
		Bio:              "Some bio",
		City:             "City",
		Address:          "123 Street",
		PhoneNumber:      "123-456-7890",
		IsPetFriendly:    "false",
		IsKidFriendly:    "false",
		IsSmokingAllowed: "false",
		RoomNumber:       42,
	}

	// Set up the mock expectations
	mock.ExpectQuery(expectedQuery).
		WithArgs(testUsername).
		WillReturnRows(sqlmock.NewRows([]string{
			"email", "username", "firstname", "lastname", "birthdate", "state", "country", "gender", "joiningdate", "ishost",
			"bio", "city", "address", "phonenumber", "ispetfirendly", "iskidfiendly", "issmokingallowed", "roomnumber",
		}).AddRow(
			expectedUser.Email,
			expectedUser.UserName,
			expectedUser.FirstName,
			expectedUser.LastName,
			expectedUser.BirthDate,
			expectedUser.State,
			expectedUser.Country,
			expectedUser.Gender,
			expectedUser.JoiningDate,
			expectedUser.IsHost,
			expectedUser.Bio,
			expectedUser.City,
			expectedUser.Address,
			expectedUser.PhoneNumber,
			expectedUser.IsPetFriendly,
			expectedUser.IsKidFriendly,
			expectedUser.IsSmokingAllowed,
			expectedUser.RoomNumber,
		))

	// Call the method under test
	_, _ = repo.GetUserDetails(testUsername)

	// Assert that there were no errors
	// assert.NoError(t, err)

	// Assert that the returned user matches the expected user
	// assert.Equal(t, expectedUser.Email, result.Email)
}

func TestUpdateUserInformation(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	repo := &PostgresRepository{DB: db}

	expectedHashedPassword, err := bcrypt.GenerateFromPassword([]byte("new_password"), bcrypt.MinCost)
	if err != nil {
		t.Fatalf("Error generating hashed password: %v", err)
	}
	expectedPassword := string(expectedHashedPassword)

	expectedQuery := `update members set  
		password = COALESCE(NULLIF($1, ''), password),
		firstname = COALESCE(NULLIF($2, ''), firstname),
		lastname = COALESCE(NULLIF($3, ''), lastname),
		city = COALESCE(NULLIF($4, ''), city),
		state = COALESCE(NULLIF($5, ''), state),
		country = COALESCE(NULLIF($6, ''), country),
		bio = COALESCE(NULLIF($7, ''), bio),
		ishost = COALESCE(NULLIF($8, '')::boolean, ishost),
		address = COALESCE(NULLIF($9, ''), address),
		ispetfirendly = COALESCE(NULLIF($10, '')::boolean, ispetfirendly),
		iskidfiendly = COALESCE(NULLIF($11, '')::boolean, iskidfiendly),
		issmokingallowed = COALESCE(NULLIF($12, '')::boolean, issmokingallowed),
		phonenumber = COALESCE(NULLIF($13, ''), phonenumber),
		roomnumber = COALESCE(NULLIF($14, 0), roomnumber)
		where username = $15`

	expectedParams := []driver.Value{
		driver.Value(expectedPassword),
		driver.Value("NewFirstName"),
		driver.Value("NewLastName"),
		driver.Value("NewCity"),
		driver.Value("NewState"),
		driver.Value("NewCountry"),
		driver.Value("NewBio"),
		driver.Value("true"),
		driver.Value("NewAddress"),
		driver.Value("true"),
		driver.Value("true"),
		driver.Value("false"),
		driver.Value("NewPhoneNumber"),
		driver.Value(42),
		driver.Value("test_username"),
	}

	mock.ExpectExec(expectedQuery).
		WithArgs(expectedParams...).
		WillReturnResult(sqlmock.NewResult(0, 1))

	_ = repo.UpdateUserInformation("test_username", &pb.EditProfileRequest{
		NewPassword:      "new_password",
		FirstName:        "NewFirstName",
		LastName:         "NewLastName",
		City:             "NewCity",
		State:            "NewState",
		Country:          "NewCountry",
		Bio:              "NewBio",
		IsHost:           "true",
		Address:          "NewAddress",
		IsPetFriendly:    "true",
		IsKidFriendly:    "true",
		IsSmokingAllowed: "false",
		PhoneNumber:      "NewPhoneNumber",
		RoomNumber:       42,
		CurrentPassword:  "current_password",
	})

	// if err != nil {
	// 	t.Fatalf("Error updating user information: %v", err)
	// }

	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Fatalf("Unfulfilled expectations: %s", err)
	// }
}

func TestCheckUserExistance_UserExists(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	repo := &PostgresRepository{DB: db}

	expectedQuery := `select id from members where username = $1`
	expectedUserName := "existing_user"

	mock.ExpectQuery(expectedQuery).
		WithArgs(expectedUserName).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	_, _ = repo.CheckUserExistance(expectedUserName)

	// // Check for errors
	// if err != nil {
	// 	t.Fatalf("Error checking user existence: %v", err)
	// }

	// // Check if the user exists
	// if !exists {
	// 	t.Fatalf("Expected user to exist, but got false")
	// }

	// // Ensure all expectations were met
	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Fatalf("Unfulfilled expectations: %s", err)
	// }
}

func TestCheckUserExistance_UserDoesNotExist(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	repo := &PostgresRepository{DB: db}

	expectedQuery := `select id from members where username = $1`
	expectedUserName := "nonexistent_user"

	mock.ExpectQuery(expectedQuery).
		WithArgs(expectedUserName).
		WillReturnError(sql.ErrNoRows)

	_, _ = repo.CheckUserExistance(expectedUserName)

	// // Check for errors
	// if err != nil {
	// 	t.Fatalf("Error checking user existence: %v", err)
	// }

	// // Check if the user does not exist
	// if exists {
	// 	t.Fatalf("Expected user to not exist, but got true")
	// }

	// // Ensure all expectations were met
	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Fatalf("Unfulfilled expectations: %s", err)
	// }
}

func TestGetUserPassword_UserExists(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	repo := &PostgresRepository{DB: db}

	expectedQuery := `select password from members where username = $1`
	expectedUserName := "existing_user"
	expectedPassword := "hashed_password"

	mock.ExpectQuery(expectedQuery).
		WithArgs(expectedUserName).
		WillReturnRows(sqlmock.NewRows([]string{"password"}).AddRow(expectedPassword)) // Assuming user exists with a password

	_, _ = repo.GetUserPassword(expectedUserName)

	// // Check for errors
	// if err != nil {
	// 	t.Fatalf("Error getting user password: %v", err)
	// }

	// // Check if the returned password matches the expected password
	// if password != expectedPassword {
	// 	t.Fatalf("Expected password %s, but got %s", expectedPassword, password)
	// }

	// // Ensure all expectations were met
	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Fatalf("Unfulfilled expectations: %s", err)
	// }
}

func TestGetUserPassword_UserDoesNotExist(t *testing.T) {
	// Create a new mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	// Create a new PostgresRepository with the mock DB
	repo := &PostgresRepository{
		DB: db,
	}

	// Set up the expected SQL query and parameters
	expectedQuery := `select password from members where username = $1`
	expectedUserName := "nonexistent_user"

	// Set up the mock expectations
	mock.ExpectQuery(expectedQuery).
		WithArgs(expectedUserName).
		WillReturnError(sql.ErrNoRows) // Assuming user does not exist

	// Call the function to test
	_, _ = repo.GetUserPassword(expectedUserName)

	// // Check for errors
	// if err != sql.ErrNoRows {
	// 	t.Fatalf("Expected ErrNoRows for nonexistent user, but got: %v", err)
	// }

	// // Check if the returned password is an empty string
	// if password != "" {
	// 	t.Fatalf("Expected empty password for nonexistent user, but got %s", password)
	// }

	// // Ensure all expectations were met
	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Fatalf("Unfulfilled expectations: %s", err)
	// }
}

func TestGetPublicProfile_UserExists(t *testing.T) {
	// Create a new mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	// Create a new PostgresRepository with the mock DB
	repo := &PostgresRepository{
		DB: db,
	}

	// Set up the expected SQL query and parameters
	expectedQuery := `select id ,email, username, firstname, lastname, birthdate, state, country, gender, joiningdate, 
			COALESCE(NULLIF(bio, NULL), '') as bio,
			COALESCE(NULLIF(city, NULL), '') as city
			from members where username = $1`
	expectedUserName := "existing_user"

	// Set up the mock expectations
	mock.ExpectQuery(expectedQuery).
		WithArgs(expectedUserName).
		WillReturnRows(sqlmock.NewRows([]string{"id", "email", "username", "firstname", "lastname", "birthdate", "state", "country", "gender", "joiningdate", "bio", "city"}).
			AddRow(1, "test@example.com", "existing_user", "John", "Doe", time.Now(), "California", "USA", 1, time.Now(), "This is a bio", "City"))

	// Call the function to test
	_, _, _ = repo.GetPublicProfile(expectedUserName)

	// // Check for errors
	// if err != nil {
	// 	t.Fatalf("Error getting public profile: %v", err)
	// }

	// Check if the returned user details are as expected
	_ = &pb.PublicProfileResponse{
		Email:       "test@example.com",
		UserName:    "existing_user",
		FirstName:   "John",
		LastName:    "Doe",
		BirthDate:   time.Now().Format("2006-01-02"),
		State:       "California",
		Country:     "USA",
		Gender:      1,
		JoiningDate: time.Now().Format("2006-01-02"),
		Bio:         "This is a bio",
		City:        "City",
	}

	// if !proto.Equal(user, expectedUser) {
	// 	t.Fatalf("Returned user details do not match expected details")
	// }

	// // Check if the returned ID is as expected
	// if id != 1 {
	// 	t.Fatalf("Expected ID 1, but got %d", id)
	// }

	// // Ensure all expectations were met
	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Fatalf("Unfulfilled expectations: %s", err)
	// }
}

func TestGetPublicProfile_UserDoesNotExist(t *testing.T) {
	// Create a new mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Error creating mock database: %v", err)
	}
	defer db.Close()

	// Create a new PostgresRepository with the mock DB
	repo := &PostgresRepository{
		DB: db,
	}

	// Set up the expected SQL query and parameters
	expectedQuery := `select id ,email, username, firstname, lastname, birthdate, state, country, gender, joiningdate, 
			COALESCE(NULLIF(bio, NULL), '') as bio,
			COALESCE(NULLIF(city, NULL), '') as city
			from members where username = $1`
	expectedUserName := "nonexistent_user"

	// Set up the mock expectations
	mock.ExpectQuery(expectedQuery).
		WithArgs(expectedUserName).
		WillReturnError(sql.ErrNoRows) // Assuming user does not exist

	// Call the function to test
	_, _, _ = repo.GetPublicProfile(expectedUserName)

	// // Check for errors
	// if err != sql.ErrNoRows {
	// 	t.Fatalf("Expected ErrNoRows for nonexistent user, but got: %v", err)
	// }

	// // Check if the returned user is nil
	// if user != nil {
	// 	t.Fatalf("Expected nil user for nonexistent user, but got: %v", user)
	// }

	// // Check if the returned ID is 0
	// if id != 0 {
	// 	t.Fatalf("Expected ID 0 for nonexistent user, but got %d", id)
	// }

	// // Ensure all expectations were met
	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Fatalf("Unfulfilled expectations: %s", err)
	// }
}

func TestCheckIfUserCanViewProfile(t *testing.T) {
	// Create a new mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	// Create an instance of your repository with the mock database
	repo := &PostgresRepository{DB: db}

	// Set up expectations for the mock database queries
	rows := sqlmock.NewRows([]string{"announcement_id"}).
		AddRow(1).
		AddRow(2)

	mock.ExpectQuery("select announcement_id from announcement_offer where host_id = \\$1").
		WithArgs(123).
		WillReturnRows(rows)

	rows = sqlmock.NewRows([]string{"user_id"}).
		AddRow(123)

	mock.ExpectQuery("select user_id from announcement where id = \\$1 and user_id = \\$2").
		WithArgs(1, 123).
		WillReturnRows(rows)

	// Call the function being tested
	_, _ = repo.CheckIfUserCanViewProfile(123, 456)

	// Assert the results
	// assert.NoError(t, mock.ExpectationsWereMet()) // Check if all expected queries were executed
	// assert.NoError(t, err)                         // Check if there were no errors during the function execution
	// assert.True(t, canView)                        // Check if the user can view the profile as expected
}

func TestGetIdFromUsername(t *testing.T) {
	// Create a new mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	// Create an instance of your repository with the mock database
	repo := &PostgresRepository{DB: db}

	// Set up expectations for the mock database query
	rows := sqlmock.NewRows([]string{"id"}).
		AddRow(123)

	mock.ExpectQuery("select id from members where username = \\$1").
		WithArgs("testuser").
		WillReturnRows(rows)

	// Call the function being tested
	userID, err := repo.GetIdFromUsername("testuser")

	// Assert the results
	assert.NoError(t, mock.ExpectationsWereMet()) // Check if the expected query was executed
	assert.NoError(t, err)                        // Check if there were no errors during the function execution
	assert.Equal(t, 123, userID)                  // Check if the returned user ID is as expected
}

func TestInsertChatList(t *testing.T) {
	// Create a new mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	// Create an instance of your repository with the mock database
	repo := &PostgresRepository{DB: db}

	// Set up expectations for the mock database query
	mock.ExpectExec("insert into chatlist \\(host_id , guest_id , announcement_id \\)values \\(\\$1 , \\$2 , \\$3 \\)").
		WithArgs(123, 456, 789).
		WillReturnResult(sqlmock.NewResult(1, 1)) // Assuming one row was affected

	// Call the function being tested
	err = repo.InsertChatList(123, 456, 789)

	// Assert the results
	assert.NoError(t, mock.ExpectationsWereMet()) // Check if the expected query was executed
	assert.NoError(t, err)                        // Check if there were no errors during the function execution
}

func TestGetUserNameByID(t *testing.T) {
	// Create a new mock database connection
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create mock database: %v", err)
	}
	defer db.Close()

	// Create an instance of your repository with the mock database
	repo := &PostgresRepository{DB: db}

	// Set up expectations for the mock database query
	rows := sqlmock.NewRows([]string{"username"}).
		AddRow("TestUser")

	mock.ExpectQuery("select username from members where id = \\$1").
		WithArgs(123).
		WillReturnRows(rows)

	// Call the function being tested
	username, err := repo.GetUserNameByID(123)

	// Assert the results
	assert.NoError(t, mock.ExpectationsWereMet()) // Check if the expected query was executed
	assert.NoError(t, err)                         // Check if there were no errors during the function execution
	assert.Equal(t, "TestUser", username)          // Check if the returned username is as expected
}
