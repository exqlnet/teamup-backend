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
table teamup.activity {

    [0 ]  activity_id               INT                  nullable: false len: 0
    [1 ]  activity_name             VARCHAR              nullable: false len: 0
    [2 ]  introduction              VARCHAR              nullable: false len: 0
    [3 ]  creator_id                INT                  nullable: false len: 0
    [4 ]  regulation                VARCHAR              nullable: false len: 0
    [5 ]  roles                     VARCHAR              nullable: false len: 0
    [6 ]  authority_code            INT                  nullable: false len: 0
    [7 ]  illustration              VARCHAR              nullable: false len: 0
    [8 ]  current_process_id        INT                  nullable: true len: 0
    [9 ]  status_code               INT                  nullable: false len: 0
    [10]  start_time                DATETIME             nullable: true len: 0
    [11]  end_time                  DATETIME             nullable: true len: 0
}

*/

// Activity struct is a row record of the activity table in the teamup database
type Activity struct {
	ActivityID       int           `gorm:"column:activity_id;primary_key"`
	ActivityName     string        `gorm:"column:activity_name"`
	Introduction     string        `gorm:"column:introduction"`
	CreatorID        int           `gorm:"column:creator_id"`
	Regulation       string        `gorm:"column:regulation"`
	Roles            string        `gorm:"column:roles"`
	AuthorityCode    int           `gorm:"column:authority_code"`
	Authority        CodeAuthority `gorm:"foreignkey:AuthorityCode"`
	Illustration     string        `gorm:"column:illustration"`
	CurrentProcessID sql.NullInt64 `gorm:"column:current_process_id"`
	CurrentProcess   ActivityProcess `gorm:"foreignkey:CurrentProcessID"`
	StatusCode       int           `gorm:"column:status_code"`
	Status           CodeActivityStatus `gorm:"foreignkey:StatusCOde"`
	StartTime        time.Time     `gorm:"column:start_time;default:CURRENT_TIMESTAMP"`
	EndTime          time.Time     `gorm:"column:end_time;default:CURRENT_TIMESTAMP"`
	Processes        []ActivityProcess `gorm:"foreignkey:ActivityID"`
	Teams			 []ActivityTeam			`gorm:"foreignkey:ActivityID"`
	DeletedAt *time.Time `sql:"index"`
}

// TableName sets the insert table name for this struct type
func (a *Activity) TableName() string {
	return "activity"
}
