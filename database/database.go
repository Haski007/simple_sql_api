package database

import (
	"os"
	"log"
	"../conf"
	"io/ioutil"
	"database/sql"
)

// DB main database variable
var DB *sql.DB

// Init Create DB variable to use SQL
func Init() {
	var err error
	DB, err = sql.Open("mysql", conf.DatabaseUser + ":" + conf.DatabasePsw + "@tcp(" + conf.Host + conf.DatabasePort + ")/")
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}

	err = createDatabase()
    if err != nil {
         log.Println(err)
         os.Exit(-1)
	}

	DB, err = sql.Open("mysql", conf.DatabaseUser + ":" + conf.DatabasePsw + "@tcp(" + conf.Host + conf.DatabasePort + ")/" + conf.DatabaseName)
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
    err = createTable()
    if err != nil {
         log.Println(err)
    }
}

func createDatabase() error {
	if _, err := DB.Exec("USE " + conf.DatabaseName); err != nil {
			_, err = DB.Exec("CREATE DATABASE " + conf.DatabaseName)
			return err
	}
	return nil
}

func createTable() error {
	file, err := ioutil.ReadFile("./database/create_users.up.sql")
	if err != nil {
			return err
	}

	_, err = DB.Exec(string(file))
	if err != nil {
			return err
	}
	return nil
}