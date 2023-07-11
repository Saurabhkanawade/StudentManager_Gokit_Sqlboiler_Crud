package dao

import (
	"context"
	"github.com/saurabhkanawade/studentmanager/internal/database"
	"github.com/saurabhkanawade/studentmanager/internal/dbmodels"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"log"
)

type Student interface {
	CreateStudent(ctx context.Context, student *dbmodels.Student) (dbmodels.Student, error)
	GetStudentById(ctx context.Context, studentId string) (dbmodels.Student, error)
	GetStudents(ctx context.Context) ([]dbmodels.Student, error)
	UpdateStudent(ctx context.Context, studentId string, student dbmodels.Student) (dbmodels.Student, error)
	DeleteStudent(ctx context.Context, studentId string) error
}

type studentDaoImpl struct {
	con database.DbConnection
}

func (s studentDaoImpl) GetStudentById(ctx context.Context, studentId string) (dbmodels.Student, error) {
	panic("implement me")
}

func (s studentDaoImpl) GetStudents(ctx context.Context) ([]dbmodels.Student, error) {
	panic("implement me")
}

func (s studentDaoImpl) UpdateStudent(ctx context.Context, studentId string, student dbmodels.Student) (dbmodels.Student, error) {
	panic("implement me")
}

func (s studentDaoImpl) DeleteStudent(ctx context.Context, studentId string) error {
	panic("implement me")
}

func (s studentDaoImpl) CreateStudent(ctx context.Context, student *dbmodels.Student) (dbmodels.Student, error) {
	log.Println("Dao() - adding new employee into the database")
	err := student.Insert(ctx, s.con.Conn, boil.Infer())
	log.Println("dao()-after the insert method")

	if err != nil {
		log.Println("Dao() - Insert - error while inserting the employee into the database")
	}
	log.Println("Dao () - Successfully added new employee to database ", student)

	return *student, nil
}

func NewStudentDao(conn database.DbConnection) Student {
	return &studentDaoImpl{
		con: conn,
	}
}
