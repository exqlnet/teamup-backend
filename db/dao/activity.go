package dao

import (
	"context"

	"teamup/db/model"

	"github.com/smallnest/gen/dbmeta"
)

// GetAllActivity is a function to get a slice of record(s) from activity table in the teamup database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - NotFound, db Find error
func GetAllActivities(ctx context.Context, page, pagesize int64, order string) (activities []*model.Activity, totalRows int, err error) {

	activities = []*model.Activity{}

	activities_orm := DB.Model(&model.Activity{})
	activities_orm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		activities_orm = activities_orm.Offset(offset).Limit(pagesize)
	} else {
		activities_orm = activities_orm.Limit(pagesize)
	}

	if order != "" {
		activities_orm = activities_orm.Order(order)
	}

	if err = activities_orm.Find(&activities).Error; err != nil {
		err = NotFound
		return nil, -1, err
	}

	return activities, totalRows, nil
}

// GetActivity is a function to get a single record to activity table in the teamup database
// error - NotFound, db Find error
func GetActivity(ctx context.Context, id interface{}) (record *model.Activity, err error) {
	if err = DB.First(record, id).Error; err != nil {
		err = NotFound
		return nil, err
	}

	return record, nil
}

// AddActivity is a function to add a single record to activity table in the teamup database
// error - InsertFailedError, db save call failed
func AddActivity(ctx context.Context, activity *model.Activity) (err error) {

	if err = DB.Save(activity).Error; err != nil {
		err = InsertFailedError
		return
	}

	return nil
}

// UpdateActivity is a function to update a single record from activity table in the teamup database
// error - NotFound, db record for id not found
// error - UpdateFailedError, db meta data copy failed or db.Save call failed
func UpdateActivity(ctx context.Context, id interface{}, updated *model.Activity) (err error) {

	activity := &model.Activity{}
	if err = DB.First(activity, id).Error; err != nil {
		err = NotFound
		return
	}

	if err = dbmeta.Copy(activity, updated); err != nil {
		err = UpdateFailedError
		return
	}

	if err = DB.Save(activity).Error; err != nil {
		err = UpdateFailedError
		return
	}

	return nil
}

// DeleteActivity is a function to delete a single record from activity table in the teamup database
// error - NotFound, db Find error
// error - DeleteFailedError, db Delete failed error
func DeleteActivity(ctx context.Context, id interface{}) (err error) {

	activity := &model.Activity{}

	if DB.First(activity, id).Error != nil {
		err = NotFound
		return
	}
	if err = DB.Delete(activity).Error; err != nil {
		err = DeleteFailedError
		return
	}

	return nil
}
