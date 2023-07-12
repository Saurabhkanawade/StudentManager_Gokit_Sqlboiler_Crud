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
	GetAllStudent(ctx context.Context) ([]model.Student, error)
	UpdateStudentById(ctx context.Context, studentId string, student model.StudentUpdate) (*model.Student, error)
	DeleteStudentById(ctx context.Context, studentId string) (string, studentIds string)
}
type StudentServiceImpl struct {
	dao    dao.Student
	Logger logrus.Logger
}

func (s StudentServiceImpl) GetAllStudent(ctx context.Context) ([]model.Student, error) {
	s.Logger.Debugf("Service() - fetching all the students.")
	studentsResponse, err := s.dao.GetStudents(ctx)

	if err != nil {
		s.Logger.Errorf("Service")
	}

	students := make([]model.Student, len(studentsResponse))

	for i := range studentsResponse {
		students[i] = model.DbToModel(*studentsResponse[i])
	}

	return students, nil
}

func (s StudentServiceImpl) UpdateStudentById(ctx context.Context, studentId string, updateStudent model.StudentUpdate) (*model.Student, error) {
	s.Logger.Debugf("Service() - Updating the student with studentId :%v", studentId)
	//getstudent

	existingStudent, err := s.GetStudentById(ctx, studentId)

	if err != nil {
		s.Logger.Errorf("Service () - nill student found with the id %v", studentId)
	}
	//setting up the updated student to payload

	student := existingStudent
	student.FullName = updateStudent.FullName
	student.Gmail = updateStudent.Gmail
	student.Phone = updateStudent.Phone

	//convertion of the model

	//updatePayload := student.ModelToDb()

	//updateStudent, _ := s.dao.UpdateStudent(ctx, updatePayload)
	//
	//resPayload := model.DbToModel(*updateStudent)

	return nil, nil
}

func (s StudentServiceImpl) DeleteStudentById(ctx context.Context, studentId string) (string, studentIds string) {
	s.Logger.Debugf("Service() - Deleting the student with studentId:%v", studentId)
	_, _ = s.dao.DeleteStudent(ctx, studentId)
	return "Deleted the student with ID :", studentIds
}

func (s StudentServiceImpl) GetStudentById(ctx context.Context, studentId string) (model.Student, error) {
	s.Logger.Debugf("Service() - fetching the student with studentId :%v", studentId)

	studentModel, err := s.dao.GetStudentById(ctx, studentId)
	if err != nil {
		s.Logger.Warnf("Service() - GetStudentByID nil %v", err)
		return model.Student{}, err
	}
	s.Logger.Debugf("Service() - get all data of the id %v from the dao is : %v", studentId, studentModel)

	return model.DbToModel(*studentModel), nil
}

func (s StudentServiceImpl) CreateStudent(ctx context.Context, student model.Student) (*model.Student, error) {
	s.Logger.Debugf("Service() - adding new student dbModel : %v", student)
	dbModel := student.ModelToDb()

	req, err := s.dao.CreateStudent(ctx, &dbModel)

	if err != nil {
		return nil, err
	}

	s.Logger.Debugf("adding new organization -- db model: %v : %v", dbModel, req)
	return &student, nil
}

func NewStudentService(studentDao dao.Student) StudentService {
	return &StudentServiceImpl{
		dao: studentDao,
	}
}
