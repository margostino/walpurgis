package db

import (
	"bufio"
	"github.com/margostino/walpurgis/pkg/helper"
	"log"
	"os"
	"strings"
	"time"
)

type UserInfo struct {
	Username     string
	LastActivity time.Time
}

type User struct {
	ID        string
	Username  string
	CreatedAt time.Time
}

func OpenFile() *os.File {
	filename := "../data/users"
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	helper.Check(err)
	return file
}

func LoadUsersData() []User {
	var allUsers = make([]User, 0)
	file := OpenFile()
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K, see next example
	for scanner.Scan() {
		values := strings.Split(scanner.Text(), ",")
		createdAt, _ := time.Parse("Wed Jan 09 20:56:37 +0000 2019", values[2])
		allUsers = append(allUsers, User{
			ID:        values[1],
			Username:  values[0],
			CreatedAt: createdAt,
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return allUsers
}
