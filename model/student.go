package model

import (
	"github.com/saurabhkanawade/studentmanager/internal/dbmodels"
	"github.com/volatiletech/null/v8"
)

type Student struct {
	Id       null.String `json:"id"`
	FullName null.String `json:"fullName"`
	Email    null.String `json:"email"`
	Phone    null.String `json:"phone"`
}

func (o Student) MakeDbModel() dbmodels.Student {
	m := dbmodels.Student{}

	m.ID = o.Id.String
	m.Fullname = o.FullName
	m.Gmail = o.Email
	m.Phone = o.Phone

	return m
}
