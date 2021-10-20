package db

import (
	"bufio"
	"github.com/margostino/walpurgis/pkg/helper"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type UserInfo struct {
	Username     string
	LastActivity time.Time
}

type User struct {
	ID                  string
	Username            string
	Name                string
	CreatedAt           time.Time
	Description         string
	Email               string
	FavouritesCount     int
	FollowRequestSent   bool
	FollowersCount      int
	FriendsCount        int
	GeoEnabled          bool
	Lang                string
	Location            string
	LastStatusCreatedAt time.Time
	Retweeted           bool
	IsRetweet           bool
	StatusRetweetCount  int
	StatusReplyCount    int
	StatusQuoteCount    int
	StatusText          string
	NextCursorStr       string
}

const (
	FILENAME string = "../data/users"
)

func TruncateFile() *os.File {
	file, err := os.OpenFile(FILENAME, os.O_TRUNC|os.O_RDWR|os.O_CREATE, 0644)
	helper.Check(err)
	return file
}

func OpenFile() *os.File {
	file, err := os.OpenFile(FILENAME, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	helper.Check(err)
	return file
}

func LoadUsersData() []*User {
	var allUsers = make([]*User, 0)
	file := OpenFile()
	defer file.Close()
	scanner := bufio.NewScanner(file)
	// optionally, resize scanner's capacity for lines over 64K
	for scanner.Scan() {
		// TODO: validations for time/date
		values := strings.Split(scanner.Text(), ",")
		createdAt, _ := time.Parse("Mon Jan 02 15:04:05 -0700 2006", values[3])
		statusCreatedAt, _ := time.Parse("Wed Jan 09 20:56:37 +0000 2019", values[12])
		favouritesCount, _ := strconv.Atoi(values[5])
		followRequestSent, _ := strconv.ParseBool(values[6])
		followersCount, _ := strconv.Atoi(values[7])
		friendsCount, _ := strconv.Atoi(values[8])
		geoEnabled, _ := strconv.ParseBool(values[9])
		retweeted, _ := strconv.ParseBool(values[13])
		isRetweet, _ := strconv.ParseBool(values[14])
		statusRetweetCount, _ := strconv.Atoi(values[15])
		statusReplyCount, _ := strconv.Atoi(values[16])
		statusQuoteCount, _ := strconv.Atoi(values[17])

		allUsers = append(allUsers, &User{
			ID:                  values[0],
			Username:            values[1],
			Name:                values[2],
			CreatedAt:           createdAt,
			Description:         values[19],
			Email:               values[4],
			FavouritesCount:     favouritesCount,
			FollowRequestSent:   followRequestSent,
			FollowersCount:      followersCount,
			FriendsCount:        friendsCount,
			GeoEnabled:          geoEnabled,
			Lang:                values[10],
			Location:            values[11],
			LastStatusCreatedAt: statusCreatedAt,
			Retweeted:           retweeted,
			IsRetweet:           isRetweet,
			StatusRetweetCount:  statusRetweetCount,
			StatusReplyCount:    statusReplyCount,
			StatusQuoteCount:    statusQuoteCount,
			StatusText:          values[18],
		})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return allUsers
}
