package configs

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/spf13/viper"
	"gopkg.in/mgo.v2"
)

func init() {
	viper.SetConfigFile("config.json")
	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}
}

func ConnectDB() (*sql.DB, error) {
	dbHost := viper.GetString(`database.host`)
	dbPort := viper.GetString(`database.port`)
	dbUser := viper.GetString(`database.user`)
	dbName := viper.GetString(`database.dbname`)

	dsn := fmt.Sprintf("%s:@tcp(%s:%s)/%s", dbUser, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func ConnectMO() (*mgo.Session, error) {
	var session, err = mgo.Dial("localhost")
	if err != nil {
		return nil, err
	}
	return session, nil
}
