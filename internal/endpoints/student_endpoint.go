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

func MakeUpdateStudentEndpoint(s services.StudentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {

		//updateRequest, err := request.(UpdateStudentRequest)

		//if !err {
		//	logrus.Errorf("Endpoint () - updating the request of updae")
		//}

		//UpdateStudent, err := s.UpdateStudentById(ctx, updateRequest.StudentId, updateRequest.Student)

		return nil, nil
	}
}

//swagger:model GetAllStudentResponse
type GetAllStudentResponse struct {
}

func MakeGetAllStudentEndpoint(s services.StudentService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		return nil, err
	}
}

// swagger:model GetStudentRequest
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
	Id       null.String `json:"id"`
	FullName null.String `json:"fullName"`
	Gmail    null.String `json:"gmail"`
	Phone    null.String `json:"phone"`
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
			Id:       req.Student.Id,
			FullName: req.Student.FullName,
			Gmail:    req.Student.Gmail,
			Phone:    req.Student.Phone,
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
