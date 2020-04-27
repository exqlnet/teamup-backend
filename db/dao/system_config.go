package dao

import (
	"context"

	"teamup/db/model"

	"github.com/smallnest/gen/dbmeta"
)

// GetAllSystemConfig is a function to get a slice of record(s) from system_config table in the teamup database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - NotFound, db Find error
func GetAllSystemConfigs(ctx context.Context, page, pagesize int64, order string) (systemconfigs []*model.SystemConfig, totalRows int, err error) {

	systemconfigs = []*model.SystemConfig{}

	systemconfigs_orm := DB.Model(&model.SystemConfig{})
	systemconfigs_orm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		systemconfigs_orm = systemconfigs_orm.Offset(offset).Limit(pagesize)
	} else {
		systemconfigs_orm = systemconfigs_orm.Limit(pagesize)
	}

	if order != "" {
		systemconfigs_orm = systemconfigs_orm.Order(order)
	}

	if err = systemconfigs_orm.Find(&systemconfigs).Error; err != nil {
		err = NotFound
		return nil, -1, err
	}

	return systemconfigs, totalRows, nil
}

// GetSystemConfig is a function to get a single record to system_config table in the teamup database
// error - NotFound, db Find error
func GetSystemConfig(ctx context.Context, id interface{}) (record *model.SystemConfig, err error) {
	if err = DB.First(record, id).Error; err != nil {
		err = NotFound
		return nil, err
	}

	return record, nil
}

// AddSystemConfig is a function to add a single record to system_config table in the teamup database
// error - InsertFailedError, db save call failed
func AddSystemConfig(ctx context.Context, systemconfig *model.SystemConfig) (err error) {

	if err = DB.Save(systemconfig).Error; err != nil {
		err = InsertFailedError
		return
	}

	return nil
}

// UpdateSystemConfig is a function to update a single record from system_config table in the teamup database
// error - NotFound, db record for id not found
// error - UpdateFailedError, db meta data copy failed or db.Save call failed
func UpdateSystemConfig(ctx context.Context, id interface{}, updated *model.SystemConfig) (err error) {

	systemconfig := &model.SystemConfig{}
	if err = DB.First(systemconfig, id).Error; err != nil {
		err = NotFound
		return
	}

	if err = dbmeta.Copy(systemconfig, updated); err != nil {
		err = UpdateFailedError
		return
	}

	if err = DB.Save(systemconfig).Error; err != nil {
		err = UpdateFailedError
		return
	}

	return nil
}

// DeleteSystemConfig is a function to delete a single record from system_config table in the teamup database
// error - NotFound, db Find error
// error - DeleteFailedError, db Delete failed error
func DeleteSystemConfig(ctx context.Context, id interface{}) (err error) {

	systemconfig := &model.SystemConfig{}

	if DB.First(systemconfig, id).Error != nil {
		err = NotFound
		return
	}
	if err = DB.Delete(systemconfig).Error; err != nil {
		err = DeleteFailedError
		return
	}

	return nil
}
