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
table teamup.activity_process {

    [0 ]  process_id                INT                  nullable: false len: 0
    [1 ]  process_name              VARCHAR              nullable: false len: 0
    [2 ]  activity_id               INT                  nullable: false len: 0
    [3 ]  start_time                DATETIME             nullable: false len: 0
}

*/

// ActivityProcess struct is a row record of the activity_process table in the teamup database
type ActivityProcess struct {
	ProcessID   int       `gorm:"column:process_id;primary_key"`
	ProcessName string    `gorm:"column:process_name"`
	ActivityID  int       `gorm:"column:activity_id"`
	StartTime   time.Time `gorm:"column:start_time"`
	Tasks	[]ActivityTask `gorm:"foreignkey:ProcessID"`
}

// TableName sets the insert table name for this struct type
func (a *ActivityProcess) TableName() string {
	return "activity_process"
}
