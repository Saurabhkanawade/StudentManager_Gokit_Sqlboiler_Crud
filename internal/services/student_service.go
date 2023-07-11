package services

import (
	"context"
	"github.com/saurabhkanawade/studentmanager/internal/dao"
	"github.com/saurabhkanawade/studentmanager/model"
	"github.com/sirupsen/logrus"
)

type StudentService interface {
	CreateStudent(ctx context.Context, student model.Student) (*model.Student, error)
	GetStudentById(ctx context.Context, studentId string) (model.Student, error)
}
type StudentServiceImpl struct {
	dao dao.Student
}

func (s StudentServiceImpl) GetStudentById(ctx context.Context, studentId string) (model.Student, error) {
	logrus.Debugf("Servive() - adding new organization -- service with studentId: %v", studentId)

	studentModel, err := s.dao.GetStudentById(ctx, studentId)
	if err != nil {
		logrus.Warnf("Service() - GetStudentByID nil %v", err)
		return model.Student{}, err
	}
	logrus.Debugf("Service() - get all data of the id %v from the dao is : %v", studentId, studentModel)

	return model.DbToModel(*studentModel), nil
}

func (s StudentServiceImpl) CreateStudent(ctx context.Context, student model.Student) (*model.Student, error) {
	logrus.Debugf("Servive() - adding new organization -- service dbModel: %v", student)
	dbModel := student.ModelToDb()
	req, err := s.dao.CreateStudent(ctx, &dbModel)

	if err != nil {
		return nil, err
	}

	logrus.Debugf("adding new organization -- db model: %v : %v", dbModel, req)
	return &student, nil
}

func NewStudentService(studentDao dao.Student) StudentService {
	return &StudentServiceImpl{
		dao: studentDao,
	}
}
