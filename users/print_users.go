package users

import (
	conf "../conf"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
)

// PrintAll prints a short list of all users
func PrintAll(w http.ResponseWriter, r *http.Request) {
	if conf.DefaultStorage == "cache" {
		for _, user := range AllUsers {
			fmt.Printf("%4d\t%s\n", user.ID, user.UserName)
		}
	} else {

	}
}

// PrintOne prints one user by GET request users/{id}
func PrintOne(w http.ResponseWriter, r *http.Request) {
	var id uint64
	var url = r.URL.Path

	match, _ := regexp.Match(`^\/users\/\d*$`, []byte(url))
	if match == true {
		id = getID(url)
	}

	if conf.DefaultStorage == "cache" {
		for _, user := range AllUsers {
			if user.ID == id {
				fmt.Printf("%16s %v\n", "Id:", user.ID)
				fmt.Printf("%16s %v\n", "UserName:", user.UserName)
				fmt.Printf("%16s %v\n", "FullName:", user.FullName)
				fmt.Printf("%16s %v\n", "City:", user.City)
				fmt.Printf("%16s %v\n", "BirthDate:", user.BirthDate)
				fmt.Printf("%16s %v\n", "Departament:", user.Departament)
				fmt.Printf("%16s %v\n", "Gender:", user.Gender)
				fmt.Printf("%16s %v\n", "ExperienceYears", user.ExperienceYears)
			}
		}
	} else {

	}
}

func getID(url string) uint64 {
	exp := regexp.MustCompile(`\d*$`)
	res, err := strconv.ParseUint(exp.FindString(url), 10, 64)
	if err != nil {
		log.Println(err)
		os.Exit(-1)
	}
	return res
}
