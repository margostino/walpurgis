package shell

import (
	"github.com/margostino/griffin/pkg/griffin"
	"github.com/margostino/walpurgis/pkg/action"
)

var SimpleActionsMapping = map[string]func(){
	"ExecuteSnapshotUsers": action.ExecuteSnapshotUsers,
	"ExecuteHelp":          ExecuteHelp,
	"ExecuteExit":          griffin.ExecuteExit,
	"ExecuteShowStats":     action.ExecuteShowStats,
}

var InputStringsActionMapping = map[string]func([]string){
	"ExecuteRankUsersBy":      action.ExecuteRankUsersBy,
	"ExecuteSelectUsersWhere": action.ExecuteSelectUsersWhere,
}
