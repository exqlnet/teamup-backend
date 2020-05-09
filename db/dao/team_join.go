package dao

import (
	"context"

	"example.com/example/model"

	"github.com/smallnest/gen/dbmeta"
)

// GetAllTeamJoin is a function to get a slice of record(s) from team_join table in the teamup database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - NotFound, db Find error
func GetAllTeamJoins(ctx context.Context, page, pagesize int64, order string) (teamjoins []*model.TeamJoin, totalRows int, err error) {

	teamjoins = []*model.TeamJoin{}

	teamjoins_orm := DB.Model(&model.TeamJoin{})
	teamjoins_orm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		teamjoins_orm = teamjoins_orm.Offset(offset).Limit(pagesize)
	} else {
		teamjoins_orm = teamjoins_orm.Limit(pagesize)
	}

	if order != "" {
		teamjoins_orm = teamjoins_orm.Order(order)
	}

	if err = teamjoins_orm.Find(&teamjoins).Error; err != nil {
		err = NotFound
		return nil, -1, err
	}

	return teamjoins, totalRows, nil
}

// GetTeamJoin is a function to get a single record to team_join table in the teamup database
// error - NotFound, db Find error
func GetTeamJoin(ctx context.Context, id interface{}) (record *model.TeamJoin, err error) {
	if err = DB.First(record, id).Error; err != nil {
		err = NotFound
		return nil, err
	}

	return record, nil
}

// AddTeamJoin is a function to add a single record to team_join table in the teamup database
// error - InsertFailedError, db save call failed
func AddTeamJoin(ctx context.Context, teamjoin *model.TeamJoin) (err error) {

	if err = DB.Save(teamjoin).Error; err != nil {
		err = InsertFailedError
		return
	}

	return nil
}

// UpdateTeamJoin is a function to update a single record from team_join table in the teamup database
// error - NotFound, db record for id not found
// error - UpdateFailedError, db meta data copy failed or db.Save call failed
func UpdateTeamJoin(ctx context.Context, id interface{}, updated *model.TeamJoin) (err error) {

	teamjoin := &model.TeamJoin{}
	if err = DB.First(teamjoin, id).Error; err != nil {
		err = NotFound
		return
	}

	if err = dbmeta.Copy(teamjoin, updated); err != nil {
		err = UpdateFailedError
		return
	}

	if err = DB.Save(teamjoin).Error; err != nil {
		err = UpdateFailedError
		return
	}

	return nil
}

// DeleteTeamJoin is a function to delete a single record from team_join table in the teamup database
// error - NotFound, db Find error
// error - DeleteFailedError, db Delete failed error
func DeleteTeamJoin(ctx context.Context, id interface{}) (err error) {

	teamjoin := &model.TeamJoin{}

	if DB.First(teamjoin, id).Error != nil {
		err = NotFound
		return
	}
	if err = DB.Delete(teamjoin).Error; err != nil {
		err = DeleteFailedError
		return
	}

	return nil
}
