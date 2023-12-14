package driver

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

// func TestNewDBConnection(t *testing.T) {
// 	viper.Set("DBHOST", "localhost")
// 	viper.Set("DBPORT", "5432")
// 	viper.Set("DBUSER", "testuser")
// 	viper.Set("DBPASSWORD", "testpassword")
// 	viper.Set("DBNAME", "testdb")

// 	NewDBConnection()
// }

func TestNewDBConnection(t *testing.T) {
	// Set up your test configuration, for example using a test database
	viper.Set("DBHOST", "localhost")
	viper.Set("DBPORT", "5432")
	viper.Set("DBUSER", "testuser")
	viper.Set("DBPASSWORD", "testpassword")
	viper.Set("DBNAME", "testdb")

	// Test case: successful database connection
	db, err := NewDBConnection()
	assert.NoError(t, err)
	assert.NotNil(t, db)

	// 	// Test case: incorrect DBHOST, should return an error
	// 	viper.Set("DBHOST", "nonexistenthost")
	// 	db, err = NewDBConnection()
	// 	assert.Error(t, err)
	// 	assert.Nil(t, db)

	// 	// Test case: incorrect DBPORT, should return an error
	// 	viper.Set("DBPORT", "invalidport")
	// 	db, err = NewDBConnection()
	// 	assert.Error(t, err)
	// 	assert.Nil(t, db)

	// 	// Test case: incorrect DBUSER, should return an error
	// 	viper.Set("DBUSER", "")
	// 	db, err = NewDBConnection()
	// 	assert.Error(t, err)
	// 	assert.Nil(t, db)

	// 	// Test case: incorrect DBPASSWORD, should return an error
	// 	viper.Set("DBPASSWORD", "")
	// 	db, err = NewDBConnection()
	// 	assert.Error(t, err)
	// 	assert.Nil(t, db)

	// // Test case: incorrect DBNAME, should return an error
	// viper.Set("DBNAME", "")
	// db, err = NewDBConnection()
	// assert.Error(t, err)
	// assert.Nil(t, db)
}
