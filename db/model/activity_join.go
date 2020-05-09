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
table teamup.activity_join {

    [0 ]  user_id                   INT                  nullable: false len: 0
    [1 ]  activity_id               INT                  nullable: false len: 0
}

*/

// ActivityJoin struct is a row record of the activity_join table in the teamup database
type ActivityJoin struct {
	UserID     int `gorm:"column:user_id;primary_key"`
	ActivityID int `gorm:"column:activity_id"`
}

// TableName sets the insert table name for this struct type
func (a *ActivityJoin) TableName() string {
	return "activity_join"
}
