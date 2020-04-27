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
table teamup.user {

    [0 ]  user_id                   INT                  nullable: false len: 0
    [1 ]  openid                    VARCHAR              nullable: false len: 0
    [2 ]  username                  VARCHAR              nullable: false len: 0
    [3 ]  avatar                    VARCHAR              nullable: false len: 0
}

*/

// User struct is a row record of the user table in the teamup database
type User struct {
	UserID   int
	Openid   string
	Username string
	Avatar   string
}

// TableName sets the insert table name for this struct type
func (u *User) TableName() string {
	return "user"
}
