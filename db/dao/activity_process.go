package dao

import (
	"context"

	"example.com/example/model"

	"github.com/smallnest/gen/dbmeta"
)

// GetAllActivityProcess is a function to get a slice of record(s) from activity_process table in the teamup database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - NotFound, db Find error
func GetAllActivityProcesses(ctx context.Context, page, pagesize int64, order string) (activityprocesses []*model.ActivityProcess, totalRows int, err error) {

	activityprocesses = []*model.ActivityProcess{}

	activityprocesses_orm := DB.Model(&model.ActivityProcess{})
	activityprocesses_orm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		activityprocesses_orm = activityprocesses_orm.Offset(offset).Limit(pagesize)
	} else {
		activityprocesses_orm = activityprocesses_orm.Limit(pagesize)
	}

	if order != "" {
		activityprocesses_orm = activityprocesses_orm.Order(order)
	}

	if err = activityprocesses_orm.Find(&activityprocesses).Error; err != nil {
		err = NotFound
		return nil, -1, err
	}

	return activityprocesses, totalRows, nil
}

// GetActivityProcess is a function to get a single record to activity_process table in the teamup database
// error - NotFound, db Find error
func GetActivityProcess(ctx context.Context, id interface{}) (record *model.ActivityProcess, err error) {
	if err = DB.First(record, id).Error; err != nil {
		err = NotFound
		return nil, err
	}

	return record, nil
}

// AddActivityProcess is a function to add a single record to activity_process table in the teamup database
// error - InsertFailedError, db save call failed
func AddActivityProcess(ctx context.Context, activityprocess *model.ActivityProcess) (err error) {

	if err = DB.Save(activityprocess).Error; err != nil {
		err = InsertFailedError
		return
	}

	return nil
}

// UpdateActivityProcess is a function to update a single record from activity_process table in the teamup database
// error - NotFound, db record for id not found
// error - UpdateFailedError, db meta data copy failed or db.Save call failed
func UpdateActivityProcess(ctx context.Context, id interface{}, updated *model.ActivityProcess) (err error) {

	activityprocess := &model.ActivityProcess{}
	if err = DB.First(activityprocess, id).Error; err != nil {
		err = NotFound
		return
	}

	if err = dbmeta.Copy(activityprocess, updated); err != nil {
		err = UpdateFailedError
		return
	}

	if err = DB.Save(activityprocess).Error; err != nil {
		err = UpdateFailedError
		return
	}

	return nil
}

// DeleteActivityProcess is a function to delete a single record from activity_process table in the teamup database
// error - NotFound, db Find error
// error - DeleteFailedError, db Delete failed error
func DeleteActivityProcess(ctx context.Context, id interface{}) (err error) {

	activityprocess := &model.ActivityProcess{}

	if DB.First(activityprocess, id).Error != nil {
		err = NotFound
		return
	}
	if err = DB.Delete(activityprocess).Error; err != nil {
		err = DeleteFailedError
		return
	}

	return nil
}
