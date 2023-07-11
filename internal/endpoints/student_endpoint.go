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
}

func MakeStudentEndpoints(s services.StudentService) StudentEndpoint {
	return StudentEndpoint{
		CreateStudentEndpoint: MakeCreateStudentEndpoint(s),
	}
}

// swagger:model CreateStudentRequest
type CreateStudentRequest struct {
	Student CreateStudentRequestBody `json:"student"`
}

type CreateStudentRequestBody struct {
	Id       null.String `json:"id"`
	FullName null.String `json:"fullName"`
	Email    null.String `json:"email"`
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
		logrus.Info("Endpoint () - called the endpoint of the student")

		req, err := request.(CreateStudentRequest)
		if !err {
			logrus.Warnf("Endpoint () - nill request :", req)
		}
		student := model.Student{
			Id:       req.Student.Id,
			FullName: req.Student.FullName,
			Email:    req.Student.Email,
			Phone:    req.Student.Phone,
		}
		serviceReq, errService := s.CreateStudent(ctx, student)

		if errService != nil {
			logrus.Warnf("Endpoint () - requesting the service with ", serviceReq)
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
