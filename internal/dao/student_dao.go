package dao

import (
	"context"
	"github.com/saurabhkanawade/studentmanager/internal/database"
	"github.com/saurabhkanawade/studentmanager/internal/dbmodels"
	"github.com/sirupsen/logrus"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"log"
)

type Student interface {
	CreateStudent(ctx context.Context, student *dbmodels.Student) (dbmodels.Student, error)
	GetStudentById(ctx context.Context, studentId string) (*dbmodels.Student, error)
	GetStudents(ctx context.Context) (dbmodels.StudentSlice, error)
	UpdateStudent(ctx context.Context, student dbmodels.Student) (*dbmodels.Student, error)
	DeleteStudent(ctx context.Context, studentId string) (string, studentIds string)
}

type studentDaoImpl struct {
	con    database.DbConnection
	Logger logrus.Logger
}

func (s studentDaoImpl) GetStudentById(ctx context.Context, studentId string) (*dbmodels.Student, error) {
	s.Logger.Infof("Dao() - Getting the student from the database. with %v", studentId)
	student, err := dbmodels.FindStudent(ctx, s.con.Conn, studentId)
	if err != nil {
		logrus.Errorf("Dao () - Get nil student in database query %v", err)
	}
	logrus.Info("Dao() - get all data of the id %v from the database is : %v", studentId, student)
	return student, nil
}

func (s studentDaoImpl) GetStudents(ctx context.Context) (dbmodels.StudentSlice, error) {
	s.Logger.Info("Dao() - Getting All the student from the database.")
	allStudent, err := dbmodels.Students().All(ctx, s.con.Conn)
	if err != nil {
		s.Logger.Errorf("Dao() - nil student found from the database %v", err)
	}
	return allStudent, nil
}

func (s studentDaoImpl) UpdateStudent(ctx context.Context, student dbmodels.Student) (*dbmodels.Student, error) {
	s.Logger.Infof("Dao() - Updating the student with with new data. %v", student)

	studentUpdate, err := student.Update(ctx, s.con.Conn, boil.Infer())

	if err != nil {
		s.Logger.Errorf("Dao () - error while updating the student with payload %v & error is %v", studentUpdate, err)
	}

	return &student, nil
}

func (s studentDaoImpl) DeleteStudent(ctx context.Context, studentId string) (string, studentIds string) {
	s.Logger.Info("Dao() - Deleting the student with %v with new data.", studentId)

	getStudent, _ := s.GetStudentById(ctx, studentId)

	_, _ = getStudent.Delete(ctx, s.con.Conn)

	return "Deleted the student with ID :", studentId
}

func (s studentDaoImpl) CreateStudent(ctx context.Context, student *dbmodels.Student) (dbmodels.Student, error) {
	s.Logger.Info("Dao() - adding new employee into the database with payload %v", student)
	err := student.Insert(ctx, s.con.Conn, boil.Infer())

	if err != nil {
		s.Logger.Debugf("Dao() - Insert - error while inserting the employee into the database")
	}
	log.Println("Dao () - Successfully added new employee to database ", student)

	return *student, nil
}

func NewStudentDao(conn database.DbConnection, log *logrus.Logger) Student {
	return &studentDaoImpl{
		con:    conn,
		Logger: *log,
	}
}
