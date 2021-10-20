package shell

import (
	"github.com/margostino/griffin/pkg/griffin"
)

var SimpleActionsMapping = map[string]func(){
	"ExecuteSnapshotUsers": ExecuteSnapshotUsers,
	"ExecuteHelp":          ExecuteHelp,
	"ExecuteExit":          griffin.ExecuteExit,
}

var InputStringsActionMapping = map[string]func([]string){
	"ExecuteRankUsersBy":      ExecuteRankUsersBy,
	"ExecuteSelectUsersWhere": ExecuteSelectUsersWhere,
}
