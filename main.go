package main

import (
	"flag"
	"fmt"
	"net/http"
	// // "os"
	// // "./usage"
	conf "./conf"
	"./users"
)

func main() {
	hostPort := flag.String("port", conf.DefaultPort, "Set custom port!")
	storage := flag.String("type", conf.DefaultStorage, "Set a type of data structure - sql/cache!")
	flag.Parse()
	conf.DefaultStorage = *storage

	users.Init(conf.UsersFile)

	http.HandleFunc("/", users.PrintOne)
	http.HandleFunc("/users", users.PrintAll)

	fmt.Println("\nListerning on port", *hostPort)
	http.ListenAndServe(*hostPort, nil)
}
