package shell

import (
	"fmt"
	"github.com/margostino/walpurgis/pkg/db"
	"strings"
)

func ExecuteSelectUsersWhere(args []string) {
	users := db.LoadUsersData()

	for _, user := range users {
		if strings.Contains(user.Description, args[0]) {
			fmt.Printf("[%s] - %s]\n", user.Username, user.Description)
		}
	}
}
