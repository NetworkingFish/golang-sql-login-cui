package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter username: ")
	scanner.Scan()
	username := scanner.Text() // here
	fmt.Print("Enter Password: ")
	scanner.Scan()
	password := scanner.Text() // here
	check1 := (usernameExists(username))
	check2 := (passwordExist(password))

	if check1 == true && check2 == true {
		fmt.Println("Access Granted.")
	} else {
		fmt.Println("Username or password does not match.")
		fmt.Println(check1)
		fmt.Println(check2)
	}

}

func usernameExists(username string) bool {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/select")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	query := "SELECT username FROM users WHERE username=?"
	row, error := db.Query(query, username)
	if error != nil {
		return false
	}

	if !row.Next() {
		return false
	}
	return true

}

func passwordExist(password string) bool {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/select")
	if err != nil {
		fmt.Println(err.Error())
	}
	defer db.Close()

	query := "SELECT password FROM users WHERE password=?"
	row, error := db.Query(query, password)
	if error != nil {
		return false
	}

	if !row.Next() {
		return false
	}
	return true

}
