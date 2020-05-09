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
table teamup.team_join {

    [0 ]  team_join_id              INT                  nullable: false len: 0
    [1 ]  user_id                   INT                  nullable: false len: 0
    [2 ]  team_id                   INT                  nullable: false len: 0
    [3 ]  role                      VARCHAR              nullable: false len: 0
}

*/

// TeamJoin struct is a row record of the team_join table in the teamup database
type TeamJoin struct {
	TeamJoinID int    `gorm:"column:team_join_id;primary_key"`
	UserID     int    `gorm:"column:user_id"`
	TeamID     int    `gorm:"column:team_id"`
	Role       string `gorm:"column:role"`
}

// TableName sets the insert table name for this struct type
func (t *TeamJoin) TableName() string {
	return "team_join"
}
