package dao

import (
	"context"

	"example.com/example/model"

	"github.com/smallnest/gen/dbmeta"
)

// GetAllActivityTask is a function to get a slice of record(s) from activity_task table in the teamup database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - NotFound, db Find error
func GetAllActivityTasks(ctx context.Context, page, pagesize int64, order string) (activitytasks []*model.ActivityTask, totalRows int, err error) {

	activitytasks = []*model.ActivityTask{}

	activitytasks_orm := DB.Model(&model.ActivityTask{})
	activitytasks_orm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		activitytasks_orm = activitytasks_orm.Offset(offset).Limit(pagesize)
	} else {
		activitytasks_orm = activitytasks_orm.Limit(pagesize)
	}

	if order != "" {
		activitytasks_orm = activitytasks_orm.Order(order)
	}

	if err = activitytasks_orm.Find(&activitytasks).Error; err != nil {
		err = NotFound
		return nil, -1, err
	}

	return activitytasks, totalRows, nil
}

// GetActivityTask is a function to get a single record to activity_task table in the teamup database
// error - NotFound, db Find error
func GetActivityTask(ctx context.Context, id interface{}) (record *model.ActivityTask, err error) {
	if err = DB.First(record, id).Error; err != nil {
		err = NotFound
		return nil, err
	}

	return record, nil
}

// AddActivityTask is a function to add a single record to activity_task table in the teamup database
// error - InsertFailedError, db save call failed
func AddActivityTask(ctx context.Context, activitytask *model.ActivityTask) (err error) {

	if err = DB.Save(activitytask).Error; err != nil {
		err = InsertFailedError
		return
	}

	return nil
}

// UpdateActivityTask is a function to update a single record from activity_task table in the teamup database
// error - NotFound, db record for id not found
// error - UpdateFailedError, db meta data copy failed or db.Save call failed
func UpdateActivityTask(ctx context.Context, id interface{}, updated *model.ActivityTask) (err error) {

	activitytask := &model.ActivityTask{}
	if err = DB.First(activitytask, id).Error; err != nil {
		err = NotFound
		return
	}

	if err = dbmeta.Copy(activitytask, updated); err != nil {
		err = UpdateFailedError
		return
	}

	if err = DB.Save(activitytask).Error; err != nil {
		err = UpdateFailedError
		return
	}

	return nil
}

// DeleteActivityTask is a function to delete a single record from activity_task table in the teamup database
// error - NotFound, db Find error
// error - DeleteFailedError, db Delete failed error
func DeleteActivityTask(ctx context.Context, id interface{}) (err error) {

	activitytask := &model.ActivityTask{}

	if DB.First(activitytask, id).Error != nil {
		err = NotFound
		return
	}
	if err = DB.Delete(activitytask).Error; err != nil {
		err = DeleteFailedError
		return
	}

	return nil
}
