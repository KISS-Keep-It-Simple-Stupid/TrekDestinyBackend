package driver

import (
	"testing"

	"github.com/spf13/viper"
)

func TestNewDBConnection(t *testing.T) {
	// Set up the configuration parameters for testing
	viper.Set("DBHOST", "localhost")
	viper.Set("DBPORT", "5432")
	viper.Set("DBUSER", "testuser")
	viper.Set("DBPASSWORD", "testpassword")
	viper.Set("DBNAME", "testdb")

	// Call the function to create a new database connection
	 NewDBConnection()
}