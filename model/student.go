package model

import (
	"github.com/saurabhkanawade/studentmanager/internal/dbmodels"
	"github.com/volatiletech/null/v8"
)

type Student struct {
	Id       null.String `json:"id"`
	FullName null.String `json:"fullName"`
	Gmail    null.String `json:"email"`
	Phone    null.String `json:"phone"`
}

// ModelToDb it is like ModelToEntity use in update and create function / PUT & POST
func (o Student) ModelToDb() dbmodels.Student {
	dbmodel := dbmodels.Student{}
	dbmodel.ID = o.Id.String
	dbmodel.Fullname = o.FullName
	dbmodel.Gmail = o.Gmail
	dbmodel.Phone = o.Phone
	return dbmodel
}

// DbToModel it is like EntityToModel use in Get and GetAll function / GET
func DbToModel(student dbmodels.Student) Student {
	model := Student{}
	model.Id = null.StringFrom(student.ID)
	model.FullName = student.Fullname
	model.Gmail = student.Gmail
	model.Phone = student.Phone
	return model
}

type StudentUpdate struct {
	FullName null.String `json:"fullName"`
	Gmail    null.String `json:"email"`
	Phone    null.String `json:"phone"`
}
