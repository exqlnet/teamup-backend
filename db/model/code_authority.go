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
table teamup.code_authority {

    [0 ]  authority_code            INT                  nullable: false len: 0
    [1 ]  name                      VARCHAR              nullable: false len: 0
}

*/

// CodeAuthority struct is a row record of the code_authority table in the teamup database
type CodeAuthority struct {
	AuthorityCode int    `gorm:"column:authority_code;primary_key"`
	Name          string `gorm:"column:name"`
}

// TableName sets the insert table name for this struct type
func (c *CodeAuthority) TableName() string {
	return "code_authority"
}
