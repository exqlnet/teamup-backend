package dao

import (
	"context"

	"example.com/example/model"

	"github.com/smallnest/gen/dbmeta"
)

// GetAllActivityTeamJoin is a function to get a slice of record(s) from activity_team_join table in the teamup database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - NotFound, db Find error
func GetAllActivityTeamJoins(ctx context.Context, page, pagesize int64, order string) (activityteamjoins []*model.ActivityTeamJoin, totalRows int, err error) {

	activityteamjoins = []*model.ActivityTeamJoin{}

	activityteamjoins_orm := DB.Model(&model.ActivityTeamJoin{})
	activityteamjoins_orm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		activityteamjoins_orm = activityteamjoins_orm.Offset(offset).Limit(pagesize)
	} else {
		activityteamjoins_orm = activityteamjoins_orm.Limit(pagesize)
	}

	if order != "" {
		activityteamjoins_orm = activityteamjoins_orm.Order(order)
	}

	if err = activityteamjoins_orm.Find(&activityteamjoins).Error; err != nil {
		err = NotFound
		return nil, -1, err
	}

	return activityteamjoins, totalRows, nil
}

// GetActivityTeamJoin is a function to get a single record to activity_team_join table in the teamup database
// error - NotFound, db Find error
func GetActivityTeamJoin(ctx context.Context, id interface{}) (record *model.ActivityTeamJoin, err error) {
	if err = DB.First(record, id).Error; err != nil {
		err = NotFound
		return nil, err
	}

	return record, nil
}

// AddActivityTeamJoin is a function to add a single record to activity_team_join table in the teamup database
// error - InsertFailedError, db save call failed
func AddActivityTeamJoin(ctx context.Context, activityteamjoin *model.ActivityTeamJoin) (err error) {

	if err = DB.Save(activityteamjoin).Error; err != nil {
		err = InsertFailedError
		return
	}

	return nil
}

// UpdateActivityTeamJoin is a function to update a single record from activity_team_join table in the teamup database
// error - NotFound, db record for id not found
// error - UpdateFailedError, db meta data copy failed or db.Save call failed
func UpdateActivityTeamJoin(ctx context.Context, id interface{}, updated *model.ActivityTeamJoin) (err error) {

	activityteamjoin := &model.ActivityTeamJoin{}
	if err = DB.First(activityteamjoin, id).Error; err != nil {
		err = NotFound
		return
	}

	if err = dbmeta.Copy(activityteamjoin, updated); err != nil {
		err = UpdateFailedError
		return
	}

	if err = DB.Save(activityteamjoin).Error; err != nil {
		err = UpdateFailedError
		return
	}

	return nil
}

// DeleteActivityTeamJoin is a function to delete a single record from activity_team_join table in the teamup database
// error - NotFound, db Find error
// error - DeleteFailedError, db Delete failed error
func DeleteActivityTeamJoin(ctx context.Context, id interface{}) (err error) {

	activityteamjoin := &model.ActivityTeamJoin{}

	if DB.First(activityteamjoin, id).Error != nil {
		err = NotFound
		return
	}
	if err = DB.Delete(activityteamjoin).Error; err != nil {
		err = DeleteFailedError
		return
	}

	return nil
}
