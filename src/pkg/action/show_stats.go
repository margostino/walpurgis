package action

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/margostino/walpurgis/pkg/context"
	"github.com/margostino/walpurgis/pkg/db"
	"github.com/margostino/walpurgis/pkg/helper"
	"strings"
)

// TODO: more stats, ML/clustering for descriptions, generalize the code, separate collections and results. Smart (cross) calculations

func ExecuteShowStats() {
	// TODO: load data users in memory once shell is started, rather than in different commands
	users := db.LoadUsersData()
	rootUser := getUser()
	if rootUser != nil {
		fmt.Println()
		fmt.Printf("Username: %s\n", rootUser.ScreenName)
		fmt.Printf("Created at: %s\n", rootUser.CreatedAt)
		fmt.Printf("Description: %s\n", rootUser.Description)
		fmt.Printf("Location: %s\n", rootUser.Location)
		fmt.Printf("Followers: %d\n", rootUser.FollowersCount)
		fmt.Printf("Following: %d\n", rootUser.FriendsCount)
		if rootUser.Status != nil {
			fmt.Printf("Last Activity at: %s\n", rootUser.Status.CreatedAt)
		}
	}

	climateCounter := 0
	climateChangeCounter := 0
	followingLessThan100Counter := 0
	followingLessThan300Counter := 0
	followingLessThan600Counter := 0
	followingLessThan1000Counter := 0
	followingMoreThan1000Counter := 0
	followersLessThan100Counter := 0
	followersLessThan300Counter := 0
	followersLessThan600Counter := 0
	followersLessThan1000Counter := 0
	followersMoreThan1000Counter := 0
	emailCounter := 0
	geoCounter := 0

	for _, user := range users {
		if strings.Contains(user.Description, "climate change") {
			climateChangeCounter += 1
		}
		if strings.Contains(user.Description, "climate") {
			climateCounter += 1
		}
		if user.GeoEnabled {
			geoCounter += 1
		}
		if user.FriendsCount < 100 {
			followingLessThan100Counter += 1
		}
		if user.FriendsCount < 300 {
			followingLessThan300Counter += 1
		}
		if user.FriendsCount < 600 {
			followingLessThan600Counter += 1
		}
		if user.FriendsCount < 1000 {
			followingLessThan1000Counter += 1
		}
		if user.FriendsCount > 1000 {
			followingMoreThan1000Counter += 1
		}
		if user.FollowersCount < 100 {
			followersLessThan100Counter += 1
		}
		if user.FollowersCount < 300 {
			followersLessThan300Counter += 1
		}
		if user.FollowersCount < 600 {
			followersLessThan600Counter += 1
		}
		if user.FollowersCount < 1000 {
			followersLessThan1000Counter += 1
		}
		if user.FollowersCount > 1000 {
			followersMoreThan1000Counter += 1
		}
		if user.Email != "" {
			emailCounter += 1
		}

	}

	fmt.Printf("Following %d accounts related with climate change\n", climateChangeCounter)
	fmt.Printf("Following %d accounts related with climate\n", climateCounter)
	fmt.Printf("Following %d accounts with Geo Location enabled\n", geoCounter)
	fmt.Printf("Following %d accounts which are following less than 100 accounts\n", followingLessThan100Counter)
	fmt.Printf("Following %d accounts which are following less than 300 accounts\n", followingLessThan300Counter)
	fmt.Printf("Following %d accounts which are following less than 600 accounts\n", followingLessThan600Counter)
	fmt.Printf("Following %d accounts which are following less than 1000 accounts\n", followingLessThan1000Counter)
	fmt.Printf("Following %d accounts which are following more than 1000 accounts\n", followingMoreThan1000Counter)
	fmt.Printf("Following %d accounts with less than 100 followers\n", followersLessThan100Counter)
	fmt.Printf("Following %d accounts with less than 300 followers\n", followersLessThan300Counter)
	fmt.Printf("Following %d accounts with less than 600 followers\n", followersLessThan600Counter)
	fmt.Printf("Following %d accounts with less than 1000 followers\n", followersLessThan1000Counter)
	fmt.Printf("Following %d accounts with more than 1000 followers\n", followersMoreThan1000Counter)
	fmt.Printf("Following %d accounts with email\n", emailCounter)
	fmt.Println()

}

func getUser() *twitter.User {
	user, _, err := context.TwitterUsers().Show(&twitter.UserShowParams{
		ScreenName: context.GetTwitterUsername(),
	})
	helper.Check(err)
	return user
}
