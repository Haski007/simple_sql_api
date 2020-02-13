package users

import (
	conf "../conf"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"../database"
)

// PrintAll prints a short list of all users
func PrintAll(w http.ResponseWriter, r *http.Request) {
	if conf.DefaultStorage == "cache" {
		for _, user := range AllUsers {
			fmt.Printf("%4d\t%s\n", user.ID, user.UserName)
		}
	} else {
		var users []user


		rows, err := database.DB.Query("SELECT id, user_name FROM users")
		if err != nil {
			log.Println(err)
			os.Exit(-1)
		}
		defer rows.Close()

		for rows.Next() {
			var u user
			err := rows.Scan(&u.ID,
				&u.UserName)
			if err != nil {
				log.Println(err)
				os.Exit(-1)
			}
			users = append(users, u)
		}
		for _, user := range users {
			fmt.Printf("%4d\t%s\n", user.ID, user.UserName)
		}
	}
}

// PrintOne prints one user by GET request users/{id}
func PrintOne(w http.ResponseWriter, r *http.Request) {
	var id uint64
	var url = r.URL.Path

	match, _ := regexp.Match(`^\/users\/\d*$`, []byte(url))
	if match != true {
		return
	}
	
	id = getID(url)
	if conf.DefaultStorage == "cache" {
		found := false
		for _, user := range AllUsers {
			if user.ID == id {
				found = true
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
		if found == false {
			log.Println("User with such id has been not found!")
		}
	} else {
		row := database.DB.QueryRow("SELECT * FROM users WHERE id=?", id)

		var resUser user
		err := row.Scan(&resUser.ID,
			&resUser.UserName,
			&resUser.FullName,
			&resUser.City,
			&resUser.BirthDate,
			&resUser.Departament,
			&resUser.Gender,
			&resUser.ExperienceYears)
		if err != nil {
			log.Println(err)
			os.Exit(-1)
		}

		fmt.Printf("%16s %v\n", "Id:", resUser.ID)
		fmt.Printf("%16s %v\n", "UserName:", resUser.UserName)
		fmt.Printf("%16s %v\n", "FullName:", resUser.FullName)
		fmt.Printf("%16s %v\n", "City:", resUser.City)
		fmt.Printf("%16s %v\n", "BirthDate:", resUser.BirthDate)
		fmt.Printf("%16s %v\n", "Departament:", resUser.Departament)
		fmt.Printf("%16s %v\n", "Gender:", resUser.Gender)
		fmt.Printf("%16s %v\n", "ExperienceYears", resUser.ExperienceYears)
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
