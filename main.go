package main

import (
	"fmt"
	"flag"
	"./conf"
	"./users"
	"net/http"
	"./database"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	hostPort := flag.String("port", conf.DefaultPort, "Set custom port!")
	storage := flag.String("type", conf.DefaultStorage, "Set a type of data structure - sql/cache!")
	flag.Parse()

	if *storage == "sql" {

		database.Init()
		defer database.DB.Close()
		
		var id uint64
		err := database.DB.QueryRow("SELECT Id FROM users").Scan(&id)
		if err != nil {
			users.Init(conf.UsersFile)
		}
	} else {
		users.Init(conf.UsersFile)
	}

	conf.DefaultStorage = *storage

	http.HandleFunc("/", users.PrintOne)
	http.HandleFunc("/users", users.PrintAll)

	fmt.Println("\nListerning on port", *hostPort)
	http.ListenAndServe(*hostPort, nil)
}
