package action

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/margostino/walpurgis/pkg/context"
	"github.com/margostino/walpurgis/pkg/db"
	"github.com/margostino/walpurgis/pkg/helper"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

func ExecuteRankUsersBy(args []string) {
	users := db.LoadUsersData()

	if len(args) == 0 {
		// TODO: pre-defined fallback/default
		rankByLastStatus(users)
	} else if len(args) == 2 && (args[1] == "asc" || args[1] == "desc") {
		rankByAttribute(users, args[0], args[1])
	} else {
		rankByAttribute(users, args[0], "asc")
	}
}

func rankByAttribute(users []*db.User, attribute string, order string) {
	rankBy(users, attribute, order)
	print(users)
}

func rankBy(users []*db.User, attribute string, order string) {
	sort.SliceStable(users, func(i, j int) bool {
		switch strings.ToLower(attribute) {
		case "status":
			return rankByTime((users)[i].LastStatusCreatedAt, (users)[j].LastStatusCreatedAt, order)
		case "age":
			return rankByTime((users)[i].CreatedAt, (users)[j].CreatedAt, order)
		case "fav":
			return rankByNumber((users)[i].FavouritesCount, (users)[j].FavouritesCount, order)
		case "followers":
			return rankByNumber((users)[i].FollowersCount, (users)[j].FollowersCount, order)
		case "following":
			return rankByNumber((users)[i].FriendsCount, (users)[j].FriendsCount, order)
		default:
			return false
		}
	})
}

func rankByTime(first time.Time, second time.Time, order string) bool {
	if order == "asc" {
		return first.Before(second)
	}
	return first.After(second)
}

func rankByNumber(first int, second int, order string) bool {
	if order == "asc" {
		return first < second
	}
	return first > second
}

func rankByLastStatus(users []*db.User) {
	sort.SliceStable(users, func(i, j int) bool {
		return (users)[i].LastStatusCreatedAt.Before((users)[j].LastStatusCreatedAt)
	})
}

//func rankByLastTimelineActivity(users []db.User) {
//	ch := make(chan db.UserInfo)
//	wg := sync.WaitGroup{}
//	wg.Add(len(users))
//	wgAppender := sync.WaitGroup{}
//	wgAppender.Add(len(users))
//
//	getLastGetLastActivityOf(&users, ch, &wg)
//
//	wg.Wait()
//	profiles := getUserProfile(users, ch, &wgAppender)
//	wgAppender.Wait()
//	sortByLastActivity(&profiles)
//	print(&profiles)
//	close(ch)
//}

func getLastGetLastActivityOf(users *[]db.User, ch chan<- db.UserInfo, wg *sync.WaitGroup) {
	for _, user := range *users {
		go getLastActivityBy(user, ch)
		wg.Done()
	}
}

func getLastActivityBy(user db.User, ch chan<- db.UserInfo) {
	var lastActivity time.Time
	id, _ := strconv.ParseInt(user.ID, 10, 64)
	tweets, _, err := context.TwitterTimelines().UserTimeline(&twitter.UserTimelineParams{
		UserID: id,
		Count:  1,
	})
	helper.Check(err)
	if len(tweets) > 0 {
		lastActivity, _ = tweets[0].CreatedAtTime()
	}
	userInfo := db.UserInfo{Username: user.Username, LastActivity: lastActivity}
	ch <- userInfo
}

func getUserProfile(users []db.User, ch <-chan db.UserInfo, wg *sync.WaitGroup) []db.UserInfo {
	var profiles = make([]db.UserInfo, 0)
	for i := 0; i < len(users); i++ {
		user := <-ch
		profiles = append(profiles, db.UserInfo{Username: user.Username, LastActivity: user.LastActivity})
		wg.Done()
	}
	return profiles
}

func sortByLastActivity(profiles *[]db.UserInfo) {
	sort.SliceStable(*profiles, func(i, j int) bool {
		return (*profiles)[i].LastActivity.Before((*profiles)[j].LastActivity)
	})
}

// TODO: generalize all prints and the requested information
func print(users []*db.User) {
	for _, user := range users {
		fmt.Printf("User: %s - Created At: %s - Last Activity: %s]\n", user.Username, user.CreatedAt, user.LastStatusCreatedAt)
	}
}
