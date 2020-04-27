package dao

import (
	"context"

	"teamup/db/model"

	"github.com/smallnest/gen/dbmeta"
)

// GetAllCodeActivityStatus is a function to get a slice of record(s) from code_activity_status table in the teamup database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - NotFound, db Find error
func GetAllCodeActivityStatuses(ctx context.Context, page, pagesize int64, order string) (codeactivitystatuses []*model.CodeActivityStatus, totalRows int, err error) {

	codeactivitystatuses = []*model.CodeActivityStatus{}

	codeactivitystatuses_orm := DB.Model(&model.CodeActivityStatus{})
	codeactivitystatuses_orm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		codeactivitystatuses_orm = codeactivitystatuses_orm.Offset(offset).Limit(pagesize)
	} else {
		codeactivitystatuses_orm = codeactivitystatuses_orm.Limit(pagesize)
	}

	if order != "" {
		codeactivitystatuses_orm = codeactivitystatuses_orm.Order(order)
	}

	if err = codeactivitystatuses_orm.Find(&codeactivitystatuses).Error; err != nil {
		err = NotFound
		return nil, -1, err
	}

	return codeactivitystatuses, totalRows, nil
}

// GetCodeActivityStatus is a function to get a single record to code_activity_status table in the teamup database
// error - NotFound, db Find error
func GetCodeActivityStatus(ctx context.Context, id interface{}) (record *model.CodeActivityStatus, err error) {
	if err = DB.First(record, id).Error; err != nil {
		err = NotFound
		return nil, err
	}

	return record, nil
}

// AddCodeActivityStatus is a function to add a single record to code_activity_status table in the teamup database
// error - InsertFailedError, db save call failed
func AddCodeActivityStatus(ctx context.Context, codeactivitystatus *model.CodeActivityStatus) (err error) {

	if err = DB.Save(codeactivitystatus).Error; err != nil {
		err = InsertFailedError
		return
	}

	return nil
}

// UpdateCodeActivityStatus is a function to update a single record from code_activity_status table in the teamup database
// error - NotFound, db record for id not found
// error - UpdateFailedError, db meta data copy failed or db.Save call failed
func UpdateCodeActivityStatus(ctx context.Context, id interface{}, updated *model.CodeActivityStatus) (err error) {

	codeactivitystatus := &model.CodeActivityStatus{}
	if err = DB.First(codeactivitystatus, id).Error; err != nil {
		err = NotFound
		return
	}

	if err = dbmeta.Copy(codeactivitystatus, updated); err != nil {
		err = UpdateFailedError
		return
	}

	if err = DB.Save(codeactivitystatus).Error; err != nil {
		err = UpdateFailedError
		return
	}

	return nil
}

// DeleteCodeActivityStatus is a function to delete a single record from code_activity_status table in the teamup database
// error - NotFound, db Find error
// error - DeleteFailedError, db Delete failed error
func DeleteCodeActivityStatus(ctx context.Context, id interface{}) (err error) {

	codeactivitystatus := &model.CodeActivityStatus{}

	if DB.First(codeactivitystatus, id).Error != nil {
		err = NotFound
		return
	}
	if err = DB.Delete(codeactivitystatus).Error; err != nil {
		err = DeleteFailedError
		return
	}

	return nil
}
