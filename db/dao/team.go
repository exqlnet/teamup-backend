package dao

import (
	"context"

	"teamup/db/model"

	"github.com/smallnest/gen/dbmeta"
)

// GetAllTeam is a function to get a slice of record(s) from team table in the teamup database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - NotFound, db Find error
func GetAllTeams(ctx context.Context, page, pagesize int64, order string) (teams []*model.Team, totalRows int, err error) {

	teams = []*model.Team{}

	teams_orm := DB.Model(&model.Team{})
	teams_orm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		teams_orm = teams_orm.Offset(offset).Limit(pagesize)
	} else {
		teams_orm = teams_orm.Limit(pagesize)
	}

	if order != "" {
		teams_orm = teams_orm.Order(order)
	}

	if err = teams_orm.Find(&teams).Error; err != nil {
		err = NotFound
		return nil, -1, err
	}

	return teams, totalRows, nil
}

// GetTeam is a function to get a single record to team table in the teamup database
// error - NotFound, db Find error
func GetTeam(ctx context.Context, id interface{}) (record *model.Team, err error) {
	if err = DB.First(record, id).Error; err != nil {
		err = NotFound
		return nil, err
	}

	return record, nil
}

// AddTeam is a function to add a single record to team table in the teamup database
// error - InsertFailedError, db save call failed
func AddTeam(ctx context.Context, team *model.Team) (err error) {

	if err = DB.Save(team).Error; err != nil {
		err = InsertFailedError
		return
	}

	return nil
}

// UpdateTeam is a function to update a single record from team table in the teamup database
// error - NotFound, db record for id not found
// error - UpdateFailedError, db meta data copy failed or db.Save call failed
func UpdateTeam(ctx context.Context, id interface{}, updated *model.Team) (err error) {

	team := &model.Team{}
	if err = DB.First(team, id).Error; err != nil {
		err = NotFound
		return
	}

	if err = dbmeta.Copy(team, updated); err != nil {
		err = UpdateFailedError
		return
	}

	if err = DB.Save(team).Error; err != nil {
		err = UpdateFailedError
		return
	}

	return nil
}

// DeleteTeam is a function to delete a single record from team table in the teamup database
// error - NotFound, db Find error
// error - DeleteFailedError, db Delete failed error
func DeleteTeam(ctx context.Context, id interface{}) (err error) {

	team := &model.Team{}

	if DB.First(team, id).Error != nil {
		err = NotFound
		return
	}
	if err = DB.Delete(team).Error; err != nil {
		err = DeleteFailedError
		return
	}

	return nil
}
