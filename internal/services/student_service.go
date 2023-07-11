package services

import (
	"context"
	"github.com/saurabhkanawade/studentmanager/internal/dao"
	"github.com/saurabhkanawade/studentmanager/model"
	"github.com/sirupsen/logrus"
)

type StudentService interface {
	CreateStudent(ctx context.Context, student model.Student) (*model.Student, error)
}
type StudentServiceImpl struct {
	dao dao.Student
}

func (s StudentServiceImpl) CreateStudent(ctx context.Context, student model.Student) (*model.Student, error) {
	logrus.Debugf("Servive() - adding new organization -- service dbModel: %v", student)
	dbModel := student.MakeDbModel()
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
