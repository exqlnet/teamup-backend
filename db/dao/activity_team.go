package dao

import (
	"context"

	"example.com/example/model"

	"github.com/smallnest/gen/dbmeta"
)

// GetAllActivityTeam is a function to get a slice of record(s) from activity_team table in the teamup database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - NotFound, db Find error
func GetAllActivityTeams(ctx context.Context, page, pagesize int64, order string) (activityteams []*model.ActivityTeam, totalRows int, err error) {

	activityteams = []*model.ActivityTeam{}

	activityteams_orm := DB.Model(&model.ActivityTeam{})
	activityteams_orm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		activityteams_orm = activityteams_orm.Offset(offset).Limit(pagesize)
	} else {
		activityteams_orm = activityteams_orm.Limit(pagesize)
	}

	if order != "" {
		activityteams_orm = activityteams_orm.Order(order)
	}

	if err = activityteams_orm.Find(&activityteams).Error; err != nil {
		err = NotFound
		return nil, -1, err
	}

	return activityteams, totalRows, nil
}

// GetActivityTeam is a function to get a single record to activity_team table in the teamup database
// error - NotFound, db Find error
func GetActivityTeam(ctx context.Context, id interface{}) (record *model.ActivityTeam, err error) {
	if err = DB.First(record, id).Error; err != nil {
		err = NotFound
		return nil, err
	}

	return record, nil
}

// AddActivityTeam is a function to add a single record to activity_team table in the teamup database
// error - InsertFailedError, db save call failed
func AddActivityTeam(ctx context.Context, activityteam *model.ActivityTeam) (err error) {

	if err = DB.Save(activityteam).Error; err != nil {
		err = InsertFailedError
		return
	}

	return nil
}

// UpdateActivityTeam is a function to update a single record from activity_team table in the teamup database
// error - NotFound, db record for id not found
// error - UpdateFailedError, db meta data copy failed or db.Save call failed
func UpdateActivityTeam(ctx context.Context, id interface{}, updated *model.ActivityTeam) (err error) {

	activityteam := &model.ActivityTeam{}
	if err = DB.First(activityteam, id).Error; err != nil {
		err = NotFound
		return
	}

	if err = dbmeta.Copy(activityteam, updated); err != nil {
		err = UpdateFailedError
		return
	}

	if err = DB.Save(activityteam).Error; err != nil {
		err = UpdateFailedError
		return
	}

	return nil
}

// DeleteActivityTeam is a function to delete a single record from activity_team table in the teamup database
// error - NotFound, db Find error
// error - DeleteFailedError, db Delete failed error
func DeleteActivityTeam(ctx context.Context, id interface{}) (err error) {

	activityteam := &model.ActivityTeam{}

	if DB.First(activityteam, id).Error != nil {
		err = NotFound
		return
	}
	if err = DB.Delete(activityteam).Error; err != nil {
		err = DeleteFailedError
		return
	}

	return nil
}
