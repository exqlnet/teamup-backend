package dao

import (
	"context"

	"example.com/example/model"

	"github.com/smallnest/gen/dbmeta"
)

// GetAllActivityTaskRecord is a function to get a slice of record(s) from activity_task_record table in the teamup database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - NotFound, db Find error
func GetAllActivityTaskRecords(ctx context.Context, page, pagesize int64, order string) (activitytaskrecords []*model.ActivityTaskRecord, totalRows int, err error) {

	activitytaskrecords = []*model.ActivityTaskRecord{}

	activitytaskrecords_orm := DB.Model(&model.ActivityTaskRecord{})
	activitytaskrecords_orm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		activitytaskrecords_orm = activitytaskrecords_orm.Offset(offset).Limit(pagesize)
	} else {
		activitytaskrecords_orm = activitytaskrecords_orm.Limit(pagesize)
	}

	if order != "" {
		activitytaskrecords_orm = activitytaskrecords_orm.Order(order)
	}

	if err = activitytaskrecords_orm.Find(&activitytaskrecords).Error; err != nil {
		err = NotFound
		return nil, -1, err
	}

	return activitytaskrecords, totalRows, nil
}

// GetActivityTaskRecord is a function to get a single record to activity_task_record table in the teamup database
// error - NotFound, db Find error
func GetActivityTaskRecord(ctx context.Context, id interface{}) (record *model.ActivityTaskRecord, err error) {
	if err = DB.First(record, id).Error; err != nil {
		err = NotFound
		return nil, err
	}

	return record, nil
}

// AddActivityTaskRecord is a function to add a single record to activity_task_record table in the teamup database
// error - InsertFailedError, db save call failed
func AddActivityTaskRecord(ctx context.Context, activitytaskrecord *model.ActivityTaskRecord) (err error) {

	if err = DB.Save(activitytaskrecord).Error; err != nil {
		err = InsertFailedError
		return
	}

	return nil
}

// UpdateActivityTaskRecord is a function to update a single record from activity_task_record table in the teamup database
// error - NotFound, db record for id not found
// error - UpdateFailedError, db meta data copy failed or db.Save call failed
func UpdateActivityTaskRecord(ctx context.Context, id interface{}, updated *model.ActivityTaskRecord) (err error) {

	activitytaskrecord := &model.ActivityTaskRecord{}
	if err = DB.First(activitytaskrecord, id).Error; err != nil {
		err = NotFound
		return
	}

	if err = dbmeta.Copy(activitytaskrecord, updated); err != nil {
		err = UpdateFailedError
		return
	}

	if err = DB.Save(activitytaskrecord).Error; err != nil {
		err = UpdateFailedError
		return
	}

	return nil
}

// DeleteActivityTaskRecord is a function to delete a single record from activity_task_record table in the teamup database
// error - NotFound, db Find error
// error - DeleteFailedError, db Delete failed error
func DeleteActivityTaskRecord(ctx context.Context, id interface{}) (err error) {

	activitytaskrecord := &model.ActivityTaskRecord{}

	if DB.First(activitytaskrecord, id).Error != nil {
		err = NotFound
		return
	}
	if err = DB.Delete(activitytaskrecord).Error; err != nil {
		err = DeleteFailedError
		return
	}

	return nil
}
