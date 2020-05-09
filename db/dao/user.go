package dao

import (
	"context"

	"example.com/example/model"

	"github.com/smallnest/gen/dbmeta"
)

// GetAllUser is a function to get a slice of record(s) from user table in the teamup database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - NotFound, db Find error
func GetAllUsers(ctx context.Context, page, pagesize int64, order string) (users []*model.User, totalRows int, err error) {

	users = []*model.User{}

	users_orm := DB.Model(&model.User{})
	users_orm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		users_orm = users_orm.Offset(offset).Limit(pagesize)
	} else {
		users_orm = users_orm.Limit(pagesize)
	}

	if order != "" {
		users_orm = users_orm.Order(order)
	}

	if err = users_orm.Find(&users).Error; err != nil {
		err = NotFound
		return nil, -1, err
	}

	return users, totalRows, nil
}

// GetUser is a function to get a single record to user table in the teamup database
// error - NotFound, db Find error
func GetUser(ctx context.Context, id interface{}) (record *model.User, err error) {
	if err = DB.First(record, id).Error; err != nil {
		err = NotFound
		return nil, err
	}

	return record, nil
}

// AddUser is a function to add a single record to user table in the teamup database
// error - InsertFailedError, db save call failed
func AddUser(ctx context.Context, user *model.User) (err error) {

	if err = DB.Save(user).Error; err != nil {
		err = InsertFailedError
		return
	}

	return nil
}

// UpdateUser is a function to update a single record from user table in the teamup database
// error - NotFound, db record for id not found
// error - UpdateFailedError, db meta data copy failed or db.Save call failed
func UpdateUser(ctx context.Context, id interface{}, updated *model.User) (err error) {

	user := &model.User{}
	if err = DB.First(user, id).Error; err != nil {
		err = NotFound
		return
	}

	if err = dbmeta.Copy(user, updated); err != nil {
		err = UpdateFailedError
		return
	}

	if err = DB.Save(user).Error; err != nil {
		err = UpdateFailedError
		return
	}

	return nil
}

// DeleteUser is a function to delete a single record from user table in the teamup database
// error - NotFound, db Find error
// error - DeleteFailedError, db Delete failed error
func DeleteUser(ctx context.Context, id interface{}) (err error) {

	user := &model.User{}

	if DB.First(user, id).Error != nil {
		err = NotFound
		return
	}
	if err = DB.Delete(user).Error; err != nil {
		err = DeleteFailedError
		return
	}

	return nil
}
