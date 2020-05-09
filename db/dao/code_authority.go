package dao

import (
	"context"

	"example.com/example/model"

	"github.com/smallnest/gen/dbmeta"
)

// GetAllCodeAuthority is a function to get a slice of record(s) from code_authority table in the teamup database
// params - page     - page requested (defaults to 0)
// params - pagesize - number of records in a page  (defaults to 20)
// params - order    - db sort order column
// error - NotFound, db Find error
func GetAllCodeAuthorities(ctx context.Context, page, pagesize int64, order string) (codeauthorities []*model.CodeAuthority, totalRows int, err error) {

	codeauthorities = []*model.CodeAuthority{}

	codeauthorities_orm := DB.Model(&model.CodeAuthority{})
	codeauthorities_orm.Count(&totalRows)

	if page > 0 {
		offset := (page - 1) * pagesize
		codeauthorities_orm = codeauthorities_orm.Offset(offset).Limit(pagesize)
	} else {
		codeauthorities_orm = codeauthorities_orm.Limit(pagesize)
	}

	if order != "" {
		codeauthorities_orm = codeauthorities_orm.Order(order)
	}

	if err = codeauthorities_orm.Find(&codeauthorities).Error; err != nil {
		err = NotFound
		return nil, -1, err
	}

	return codeauthorities, totalRows, nil
}

// GetCodeAuthority is a function to get a single record to code_authority table in the teamup database
// error - NotFound, db Find error
func GetCodeAuthority(ctx context.Context, id interface{}) (record *model.CodeAuthority, err error) {
	if err = DB.First(record, id).Error; err != nil {
		err = NotFound
		return nil, err
	}

	return record, nil
}

// AddCodeAuthority is a function to add a single record to code_authority table in the teamup database
// error - InsertFailedError, db save call failed
func AddCodeAuthority(ctx context.Context, codeauthority *model.CodeAuthority) (err error) {

	if err = DB.Save(codeauthority).Error; err != nil {
		err = InsertFailedError
		return
	}

	return nil
}

// UpdateCodeAuthority is a function to update a single record from code_authority table in the teamup database
// error - NotFound, db record for id not found
// error - UpdateFailedError, db meta data copy failed or db.Save call failed
func UpdateCodeAuthority(ctx context.Context, id interface{}, updated *model.CodeAuthority) (err error) {

	codeauthority := &model.CodeAuthority{}
	if err = DB.First(codeauthority, id).Error; err != nil {
		err = NotFound
		return
	}

	if err = dbmeta.Copy(codeauthority, updated); err != nil {
		err = UpdateFailedError
		return
	}

	if err = DB.Save(codeauthority).Error; err != nil {
		err = UpdateFailedError
		return
	}

	return nil
}

// DeleteCodeAuthority is a function to delete a single record from code_authority table in the teamup database
// error - NotFound, db Find error
// error - DeleteFailedError, db Delete failed error
func DeleteCodeAuthority(ctx context.Context, id interface{}) (err error) {

	codeauthority := &model.CodeAuthority{}

	if DB.First(codeauthority, id).Error != nil {
		err = NotFound
		return
	}
	if err = DB.Delete(codeauthority).Error; err != nil {
		err = DeleteFailedError
		return
	}

	return nil
}
