package main

import (
	"log"
	"flag"
	"fmt"
	"net/http"
	"os"
	conf "./conf"
	"./users"
	"./database"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	hostPort := flag.String("port", conf.DefaultPort, "Set custom port!")
	storage := flag.String("type", conf.DefaultStorage, "Set a type of data structure - sql/cache!")
	flag.Parse()

	//////////////////////////////////////

	var err error
	database.DB, err = sql.Open("mysql", "demian:Zak-Efron123@tcp(" + conf.Host + ":3306)/quests")
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	defer database.DB.Close()

	var id uint64
	err = database.DB.QueryRow("SELECT Id FROM users").Scan(&id)
	if err != nil || conf.DefaultStorage == "cache" {
		users.Init(conf.UsersFile)
	}

	//////////////////////////////////

	conf.DefaultStorage = *storage

	http.HandleFunc("/", users.PrintOne)
	http.HandleFunc("/users", users.PrintAll)

	fmt.Println("\nListerning on port", *hostPort)
	http.ListenAndServe(*hostPort, nil)
}
