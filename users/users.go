package users

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	conf "../conf"
)

type user struct {
	ID              uint64 `json:"Id"`
	UserName        string `json:"Username"`
	FullName        string `json:"FullName"`
	City            string `json:"City"`
	BirthDate       string `json:"BirthDate"`
	Departament     string `json:"Department"`
	Gender          string `json:"Gender"`
	ExperienceYears int    `json:"ExperienceYears"`
}

// AllUsers - cache of the program.
var AllUsers = []*user{}

// Init takes name of .json file and type of storage
// and fills DataBase by file's data
func Init(fileName string) {
	if conf.DefaultStorage == "sql" {
		initSQL(fileName)
	} else if conf.DefaultStorage == "cache" {
		initCache(fileName)
	} else {
		log.Println("Unknown data structure!")
		os.Exit(-1)
	}
}

// InitSQL takes name of .json file with users
// and fills DataBase by that data
func initSQL(fileName string) {

}

// InitCache takes name of .json file with users
// and fills cache by that data
func initCache(fileName string) {
	fmt.Println("Starting encode...")

	jsonFile, err := os.Open(fileName)
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	defer jsonFile.Close()

	bytes, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}

	json.Unmarshal(bytes, &AllUsers)

	fmt.Println("Table has been filled!")
}
