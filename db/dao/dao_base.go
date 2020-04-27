package dao

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

// BuildInfo is used to define the application build info, and inject values into via the build process.
type BuildInfo struct {
	BuildDate    string
	LatestCommit string
	BuildNumber  string
	BuiltOnIp    string
	BuiltOnOs    string
	RuntimeVer   string
}

var (
	NotFound            = fmt.Errorf("record Not Found")
	UnableToMarshalJson = fmt.Errorf("json payload corrupt")
	UpdateFailedError   = fmt.Errorf("db update error")
	InsertFailedError   = fmt.Errorf("db insert error")
	DeleteFailedError   = fmt.Errorf("db delete error")
	BadParamsError      = fmt.Errorf("bad params error")

	DB           *gorm.DB
	AppBuildInfo *BuildInfo
)
