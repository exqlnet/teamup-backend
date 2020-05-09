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
table teamup.team {

    [0 ]  team_id                   INT                  nullable: false len: 0
    [1 ]  activity_id               INT                  nullable: false len: 0
    [2 ]  team_name                 VARCHAR              nullable: false len: 0
}

*/

// Team struct is a row record of the team table in the teamup database
type Team struct {
	TeamID     int    `gorm:"column:team_id;primary_key"`
	ActivityID int    `gorm:"column:activity_id"`
	TeamName   string `gorm:"column:team_name"`
	Slogan     string `gorm:"column:slogan"`
}

// TableName sets the insert table name for this struct type
func (t *Team) TableName() string {
	return "team"
}
