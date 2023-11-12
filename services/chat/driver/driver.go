package driver

import (
	"database/sql"
	"fmt"

	"github.com/spf13/viper"
	_ "github.com/lib/pq"
)

func NewDBConnection() (*sql.DB, error) {
	var (
		host = viper.Get("DBHOST").(string)
		port = viper.Get("DBPORT").(string)
		user = viper.Get("DBUSER").(string)
		password = viper.Get("DBPASSWORD").(string)
		dbname = viper.Get("DBNAME").(string)
	)

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
