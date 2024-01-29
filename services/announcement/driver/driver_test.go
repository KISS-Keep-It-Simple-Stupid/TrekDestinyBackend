package driver

import (
	"testing"

	"github.com/spf13/viper"
)

func TestNewDBConnection(t *testing.T) {
	viper.Set("DBHOST", "localhost")
	viper.Set("DBPORT", "5432")
	viper.Set("DBUSER", "testuser")
	viper.Set("DBPASSWORD", "testpassword")
	viper.Set("DBNAME", "testdb")

	_, _ = NewDBConnection()
}
