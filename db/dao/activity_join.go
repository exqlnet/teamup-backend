package dao

import (
	"context"

	"example.com/example/model"

	"github.com/smallnest/gen/dbmeta"
)

// GetAllActivityJoin is a function to get a slice of record(s) from activity_join table in the teamup database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - NotFound, db Find error
func GetAllActivityJoins(ctx context.Context, page, pagesize int64, order string) (activityjoins []*model.ActivityJoin, totalRows int, err error) {

	activityjoins = []*model.ActivityJoin{}

	activityjoins_orm := DB.Model(&model.ActivityJoin{})
	activityjoins_orm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		activityjoins_orm = activityjoins_orm.Offset(offset).Limit(pagesize)
	} else {
		activityjoins_orm = activityjoins_orm.Limit(pagesize)
	}

	if order != "" {
		activityjoins_orm = activityjoins_orm.Order(order)
	}

	if err = activityjoins_orm.Find(&activityjoins).Error; err != nil {
		err = NotFound
		return nil, -1, err
	}

	return activityjoins, totalRows, nil
}

// GetActivityJoin is a function to get a single record to activity_join table in the teamup database
// error - NotFound, db Find error
func GetActivityJoin(ctx context.Context, id interface{}) (record *model.ActivityJoin, err error) {
	if err = DB.First(record, id).Error; err != nil {
		err = NotFound
		return nil, err
	}

	return record, nil
}

// AddActivityJoin is a function to add a single record to activity_join table in the teamup database
// error - InsertFailedError, db save call failed
func AddActivityJoin(ctx context.Context, activityjoin *model.ActivityJoin) (err error) {

	if err = DB.Save(activityjoin).Error; err != nil {
		err = InsertFailedError
		return
	}

	return nil
}

// UpdateActivityJoin is a function to update a single record from activity_join table in the teamup database
// error - NotFound, db record for id not found
// error - UpdateFailedError, db meta data copy failed or db.Save call failed
func UpdateActivityJoin(ctx context.Context, id interface{}, updated *model.ActivityJoin) (err error) {

	activityjoin := &model.ActivityJoin{}
	if err = DB.First(activityjoin, id).Error; err != nil {
		err = NotFound
		return
	}

	if err = dbmeta.Copy(activityjoin, updated); err != nil {
		err = UpdateFailedError
		return
	}

	if err = DB.Save(activityjoin).Error; err != nil {
		err = UpdateFailedError
		return
	}

	return nil
}

// DeleteActivityJoin is a function to delete a single record from activity_join table in the teamup database
// error - NotFound, db Find error
// error - DeleteFailedError, db Delete failed error
func DeleteActivityJoin(ctx context.Context, id interface{}) (err error) {

	activityjoin := &model.ActivityJoin{}

	if DB.First(activityjoin, id).Error != nil {
		err = NotFound
		return
	}
	if err = DB.Delete(activityjoin).Error; err != nil {
		err = DeleteFailedError
		return
	}

	return nil
}
