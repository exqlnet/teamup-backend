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
table teamup.activity_task_record {

    [0 ]  team_id                   INT                  nullable: false len: 0
    [1 ]  task_id                   INT                  nullable: false len: 0
}

*/

// ActivityTaskRecord struct is a row record of the activity_task_record table in the teamup database
type ActivityTaskRecord struct {
	TeamID int `gorm:"column:team_id;primary_key"`
	TaskID int `gorm:"column:task_id"`
}

// TableName sets the insert table name for this struct type
func (a *ActivityTaskRecord) TableName() string {
	return "activity_task_record"
}
