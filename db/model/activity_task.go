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
table teamup.activity_task {

    [0 ]  task_id                   INT                  nullable: false len: 0
    [1 ]  task_name                 VARCHAR              nullable: false len: 0
    [2 ]  process_id                INT                  nullable: false len: 0
    [3 ]  role                      VARCHAR              nullable: false len: 0
}

*/

// ActivityTask struct is a row record of the activity_task table in the teamup database
type ActivityTask struct {
	TaskID    int    `gorm:"column:task_id;primary_key"`
	TaskName  string `gorm:"column:task_name"`
	ProcessID int    `gorm:"column:process_id"`
	Role      string `gorm:"column:role"`
}

// TableName sets the insert table name for this struct type
func (a *ActivityTask) TableName() string {
	return "activity_task"
}
