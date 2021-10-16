package shell

import (
	"github.com/margostino/griffin/pkg/griffin"
)

var Mapping = map[string]func(){
	"ExecuteRankUsers":     ExecuteRankUsers,
	"ExecuteSnapshotUsers": ExecuteSnapshotUsers,
	"ExecuteHelp":          ExecuteHelp,
	"ExecuteExit":          griffin.ExecuteExit,
}
