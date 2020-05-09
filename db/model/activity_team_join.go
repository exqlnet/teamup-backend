package model

import (
	"database/sql"
	"time"

	"github.com/guregu/null"
)

var (
	_ = time.Second
	_ = sql.LevelDefault
	_ = null.Bool{}
)

/*
DB Table Details
table teamup.activity_team_join {

    [0 ]  user_id                   INT                  nullable: false len: 0
    [1 ]  team_id                   INT                  nullable: false len: 0
    [2 ]  join_time                 DATETIME             nullable: true len: 0
    [3 ]  role                      INT                  nullable: true len: 0
}

*/

// ActivityTeamJoin struct is a row record of the activity_team_join table in the teamup database
type ActivityTeamJoin struct {
	UserID   int           `gorm:"column:user_id;primary_key"`
	TeamID   int           `gorm:"column:team_id"`
	JoinTime time.Time     `gorm:"column:join_time"`
	Role     sql.NullInt64 `gorm:"column:role"`
}

// TableName sets the insert table name for this struct type
func (a *ActivityTeamJoin) TableName() string {
	return "activity_team_join"
}
