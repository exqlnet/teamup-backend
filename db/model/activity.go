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
    [6 ]  authority_code            INT                  nullable: true len: 0
    [7 ]  illustration              VARCHAR              nullable: false len: 0
    [8 ]  process                   VARCHAR              nullable: false len: 0
    [9 ]  current_process           VARCHAR              nullable: false len: 0
    [10]  status_code               INT                  nullable: true len: 0
}

*/

// Activity struct is a row record of the activity table in the teamup database
type Activity struct {
	ActivityID     int
	ActivityName   string
	Introduction   string
	CreatorID      int
	Regulation     string
	Roles          string
	AuthorityCode  sql.NullInt64
	Illustration   string
	Process        string
	CurrentProcess string
	StatusCode     sql.NullInt64
}

// TableName sets the insert table name for this struct type
func (a *Activity) TableName() string {
	return "activity"
}
