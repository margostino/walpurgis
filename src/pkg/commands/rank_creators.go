package commands

import (
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/margostino/walpurgis/pkg/client"
	"github.com/margostino/walpurgis/pkg/db"
	"sort"
	"strconv"
	"time"
)

// TODO: call in parallel, improve performance

func RankCreators() {
	var profiles = make([]db.UserInfo, 0)
	users := db.LoadUsersData()
	ch := make(chan db.UserInfo)
	for _, user := range users {
		//wg.Add(1)
		go func(u db.User) {
			var lastActivity time.Time
			id, _ := strconv.ParseInt(u.ID, 10, 64)
			tweets, _, _ := client.Twitter.Timelines.UserTimeline(&twitter.UserTimelineParams{
				UserID: id,
				Count:  1,
			})

			if len(tweets) > 0 {
				lastActivity, _ = tweets[0].CreatedAtTime()
			}
			userInfo := db.UserInfo{Username: u.Username, LastActivity: lastActivity}
			ch <- userInfo
		}(user)
	}

	for i := 0; i < len(users); i++ {
		user := <-ch
		profiles = append(profiles, db.UserInfo{Username: user.Username, LastActivity: user.LastActivity})
	}

	sort.SliceStable(profiles, func(i, j int) bool {
		return profiles[i].LastActivity.Before(profiles[j].LastActivity)
	})

	for _, profile := range profiles {
		fmt.Printf("[%s] - %s]\n", profile.Username, profile.LastActivity)
	}

}
