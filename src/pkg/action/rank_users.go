package action

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/margostino/walpurgis/pkg/client"
	"github.com/margostino/walpurgis/pkg/db"
	"github.com/margostino/walpurgis/pkg/helper"
	"sort"
	"strconv"
	"sync"
	"time"
)

// TODO: call in parallel, improve performance

func ExecuteRankUsers() {
	users := db.LoadUsersData()

	ch := make(chan db.UserInfo)
	wg := sync.WaitGroup{}
	wg.Add(len(users))
	wgAppender := sync.WaitGroup{}
	wgAppender.Add(len(users))

	getLastGetLastActivityOf(&users, ch, &wg)

	wg.Wait()
	profiles := getUserProfile(users, ch, &wgAppender)
	wgAppender.Wait()
	sortByLastActivity(&profiles)
	print(&profiles)
	close(ch)
}

func getLastGetLastActivityOf(users *[]db.User, ch chan<- db.UserInfo, wg *sync.WaitGroup) {
	for _, user := range *users {
		go getLastActivityBy(user, ch)
		wg.Done()
	}
}

func getLastActivityBy(user db.User, ch chan<- db.UserInfo) {
	var lastActivity time.Time
	id, _ := strconv.ParseInt(user.ID, 10, 64)
	tweets, _, err := client.Twitter.Timelines.UserTimeline(&twitter.UserTimelineParams{
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

func print(profiles *[]db.UserInfo) {
	for _, profile := range *profiles {
		fmt.Printf("[%s] - %s]\n", profile.Username, profile.LastActivity)
	}
}
