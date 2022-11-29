package action

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/margostino/walpurgis/common"
	"github.com/margostino/walpurgis/config"
	"github.com/margostino/walpurgis/social"
	"strings"
)

// TODO: more stats, ML/clustering for descriptions, generalize the code, separate collections and results. Smart (cross) calculations

func ExecuteShowStats() {
	users := config.GetUsersData()
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
	total := len(users)

	for _, user := range users {
		climateChangeCounter += countIf(strings.Contains(user.Description, "climate change"))
		climateCounter += countIf(strings.Contains(user.Description, "climate"))
		geoCounter += countIf(user.GeoEnabled)
		followingLessThan100Counter += countIf(user.FriendsCount < 100)
		followingLessThan300Counter += countIf(user.FriendsCount < 300)
		followingLessThan600Counter += countIf(user.FriendsCount < 600)
		followingLessThan1000Counter += countIf(user.FriendsCount < 1000)
		followingMoreThan1000Counter += countIf(user.FriendsCount > 1000)
		followersLessThan100Counter += countIf(user.FollowersCount < 100)
		followersLessThan300Counter += countIf(user.FollowersCount < 300)
		followersLessThan600Counter += countIf(user.FollowersCount < 600)
		followersLessThan1000Counter += countIf(user.FollowersCount < 1000)
		followersMoreThan1000Counter += countIf(user.FollowersCount > 1000)
		emailCounter += countIf(user.Email != "")
	}

	fmt.Printf("Following %0.2f%% accounts related with climate change\n", (float64(climateChangeCounter))/(float64(total)))
	fmt.Printf("Following %0.2f%% accounts related with climate\n", (float64(climateCounter))/(float64(total)))
	fmt.Printf("Following %0.2f%% accounts with Geo Location enabled\n", (float64(geoCounter))/(float64(total)))
	fmt.Printf("Following %0.2f%% accounts which are following less than 100 accounts\n", (float64(followingLessThan100Counter))/(float64(total)))
	fmt.Printf("Following %0.2f%% accounts which are following less than 300 accounts\n", (float64(followingLessThan300Counter))/(float64(total)))
	fmt.Printf("Following %0.2f%% accounts which are following less than 600 accounts\n", (float64(followingLessThan600Counter))/(float64(total)))
	fmt.Printf("Following %0.2f%% accounts which are following less than 1000 accounts\n", (float64(followingLessThan1000Counter))/(float64(total)))
	fmt.Printf("Following %0.2f%% accounts which are following more than 1000 accounts\n", (float64(followingMoreThan1000Counter))/(float64(total)))
	fmt.Printf("Following %0.2f%% accounts with less than 100 followers\n", (float64(followersLessThan100Counter))/(float64(total)))
	fmt.Printf("Following %0.2f%% accounts with less than 300 followers\n", (float64(followersLessThan300Counter))/(float64(total)))
	fmt.Printf("Following %0.2f%% accounts with less than 600 followers\n", (float64(followersLessThan600Counter))/(float64(total)))
	fmt.Printf("Following %0.2f%% accounts with less than 1000 followers\n", (float64(followersLessThan1000Counter))/(float64(total)))
	fmt.Printf("Following %0.2f%% accounts with more than 1000 followers\n", (float64(followersMoreThan1000Counter))/(float64(total)))
	fmt.Printf("Following %0.2f%% accounts with email\n", (float64(emailCounter))/(float64(total)))
	fmt.Println()

}

func countIf(condition bool) int {
	if condition {
		return 1
	}
	return 0
}

func getUser() *twitter.User {
	user, _, err := social.TwitterUsers().Show(&twitter.UserShowParams{
		ScreenName: config.GetTwitterUsername(),
	})
	common.Check(err)
	return user
}
