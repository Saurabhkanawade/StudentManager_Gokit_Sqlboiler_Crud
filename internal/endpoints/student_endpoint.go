package endpoints

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/saurabhkanawade/studentmanager/internal/services"
	"github.com/saurabhkanawade/studentmanager/model"
	"github.com/sirupsen/logrus"
	"github.com/volatiletech/null/v8"
)

type StudentEndpoint struct {
	CreateStudentEndpoint endpoint.Endpoint
	GetStudentEndpoint    endpoint.Endpoint
	GetAllStudentEndpoint endpoint.Endpoint
	UpdateStudentEndpoint endpoint.Endpoint
	DeleteStudentEndpoint endpoint.Endpoint
}

func MakeStudentEndpoints(s services.StudentService) StudentEndpoint {
	return StudentEndpoint{
		CreateStudentEndpoint: MakeCreateStudentEndpoint(s),
		GetStudentEndpoint:    MakeGetStudentEndpoint(s),
		GetAllStudentEndpoint: MakeGetAllStudentEndpoint(s),
		UpdateStudentEndpoint: MakeUpdateStudentEndpoint(s),
		DeleteStudentEndpoint: MakeDeleteStudentEndpoint(s),
	}
}

type StudentDeleteRequest struct {
	StudentId string
}

// DeleteStudentResponse
// swagger:response DeleteStudentResponse
type DeleteStudentResponse struct {
	StudentId string
}

func MakeDeleteStudentEndpoint(s services.StudentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		logrus.Info("Endpoint () - called the endpoint of the DeleteStudent.")
		reqStudent, err := request.(StudentDeleteRequest)

		if !err {
			logrus.Warnf("Endpoint() - getStudent nill %v ", reqStudent)
		}

		deleteStudent, ok := s.DeleteStudentById(ctx, reqStudent.StudentId)
		logrus.Infof("Endpoint () - deleted the student with studentId %v response is %v and return is %v", reqStudent.StudentId, deleteStudent, ok)

		res := DeleteStudentResponse{
			StudentId: deleteStudent,
		}

		return res, nil

	}
}

type UpdateStudentRequest struct {
	StudentId string
	Student   model.StudentUpdate
}

// swagger:model StudentRequest
type StudentRequest struct {
	Id       string `json:"id"`
	FullName string `json:"fullName"`
	Gmail    string `json:"email"`
	Phone    string `json:"phone"`
}

// UpdateStudentResponse
// swagger:response UpdateStudentResponse
type UpdateStudentResponse struct {
	//in:body
	Student model.Student `json:"student"`
}

func MakeUpdateStudentEndpoint(s services.StudentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		updateRequest, err := request.(StudentRequest)

		if !err {
			logrus.Errorf("Endpoint () - updating the request of update")
		}

		logrus.Infof("Endpoint() - update request %v", updateRequest)

		student := model.Student{
			Id:       null.StringFrom(updateRequest.Id),
			FullName: null.StringFrom(updateRequest.FullName),
			Gmail:    null.StringFrom(updateRequest.Gmail),
			Phone:    null.StringFrom(updateRequest.Phone),
		}

		UpdateStudentRes, updateErr := s.UpdateStudentById(ctx, updateRequest.Id, student)

		if updateErr != nil {
			logrus.Errorf("Endpoint () - error in updateStudentByID %v", err)
		}

		logrus.Debugf("Endpoint () - updated student payload %v", student)
		logrus.Debugf("Endpoint () - updated student response %v", UpdateStudentRes)

		updateStudentResponse := UpdateStudentResponse{
			Student: *UpdateStudentRes,
		}
		return updateStudentResponse, nil
	}
}

type GetAllStudentResponse struct {
	Students []model.Student
}

// GetAllStudentResponseBody
//
//swagger:response GetAllStudentResponseBody
type GetAllStudentResponseBody struct {
	//in:body
	Body GetAllStudentResponse `json:",inline"`
}

func MakeGetAllStudentEndpoint(s services.StudentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		allStudents, err := s.GetAllStudent(ctx)
		logrus.Infof("Endpoint () - getall student %v", allStudents)

		res := GetAllStudentResponse{
			Students: allStudents,
		}
		return res, nil
	}
}

type GetStudentByIdRequest struct {
	StudentId string
}

// GetStudentByIdResponseBody
// swagger:response GetStudentByIdResponseBody
type GetStudentByIdResponseBody struct {
	//in:body
	Student model.Student `json:"student"`
}

func MakeGetStudentEndpoint(s services.StudentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		logrus.Info("Endpoint () - called the endpoint of the GetStudent")

		getStudent, err := request.(GetStudentByIdRequest)
		if !err {
			logrus.Warnf("Endpoint() - getStudent nill %v ", getStudent)
		}

		student, err1 := s.GetStudentById(ctx, getStudent.StudentId)

		if err1 != nil {
			return nil, err1
		}

		logrus.Debugf("Endpoint() - get all data of the id %v from the dao is : %v", getStudent, student)

		Body := GetStudentByIdResponseBody{
			Student: student,
		}
		return Body.Student, nil
	}
}

// swagger:model CreateStudentRequest
type CreateStudentRequest struct {
	Student CreateStudentRequestBody `json:"student"`
}

type CreateStudentRequestBody struct {
	Id       string `json:"id"`
	FullName string `json:"fullName"`
	Gmail    string `json:"gmail"`
	Phone    string `json:"phone"`
}

type createResponseBody struct {
	Student model.Student `json:"student"`
}

// createStudentResponse
// swagger:response createStudentResponse
type createStudentResponse struct {
	//in:body
	Body createResponseBody `json:",inline"`
}

func MakeCreateStudentEndpoint(s services.StudentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		logrus.Info("Endpoint () - called the endpoint of the CreateStudent")

		req, err := request.(CreateStudentRequest)
		if !err {
			logrus.Warnf("Endpoint () - nill request %v:", req)
		}
		student := model.Student{
			Id:       null.StringFrom(req.Student.Id),
			FullName: null.StringFrom(req.Student.FullName),
			Gmail:    null.StringFrom(req.Student.Gmail),
			Phone:    null.StringFrom(req.Student.Phone),
		}
		serviceReq, errService := s.CreateStudent(ctx, student)

		if errService != nil {
			logrus.Warnf("Endpoint () - requesting the service with %v", serviceReq)
		}

		res := createStudentResponse{
			Body: createResponseBody{
				Student: *serviceReq,
			},
		}
		logrus.Debugf("Endpoint () - student response body %v", res.Body)
		return res.Body, errService
	}
}
