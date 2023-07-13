package transport

import (
	"context"
	"encoding/json"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/saurabhkanawade/studentmanager/internal/endpoints"
	"github.com/saurabhkanawade/studentmanager/model"
	"github.com/sirupsen/logrus"
	"io"
	"net/http"
)

func CreateStudentHttpHandler(endpoint endpoints.StudentEndpoint, router *mux.Router) {

	router.Handle("/student",
		httptransport.NewServer(
			endpoint.CreateStudentEndpoint,
			decodeCreateStudent,
			encodeStudent,
		)).Methods(http.MethodPost)

	router.Handle("/students/{studentId}",
		httptransport.NewServer(
			endpoint.GetStudentEndpoint,
			decodeGetByIdStudent,
			encodeStudent,
		)).Methods(http.MethodGet)

	router.Handle("/students/{studentId}",
		httptransport.NewServer(
			endpoint.DeleteStudentEndpoint,
			decodeDeleteStudent,
			encodeStudent,
		)).Methods(http.MethodDelete)

	router.Handle("/students/{studentId}",
		httptransport.NewServer(
			endpoint.UpdateStudentEndpoint,
			decodeUpdateStudent,
			encodeStudent,
		))

	router.Handle("/students",
		httptransport.NewServer(
			endpoint.GetAllStudentEndpoint,
			decodeGetStudents,
			encodeStudent,
		))
}

func decodeUpdateStudent(ctx context.Context, request2 *http.Request) (request interface{}, err error) {
	logrus.Info("Decode with update request")
	var student endpoints.StudentRequest

	vars := mux.Vars(request2)

	studentId, ok := vars["studentId"]
	logrus.Infof("Decode () - update the request with studentId %v", studentId)
	if !ok {
		logrus.Warnf("DecodeDeleteStudent() - error while getting vars %v", ok)
	}
	body, err := io.ReadAll(request2.Body)

	err = json.Unmarshal(body, &student)

	//student.Id = null.StringFrom(studentId)

	if err != nil {
		logrus.Warnf("Decode () - Error while unmarshaling %v ", err)
	}
	logrus.Debugf("Decode () -transport add organization incoming object: %v")

	return endpoints.StudentRequest{
		Id:       student.Id,
		FullName: student.FullName,
		Gmail:    student.Gmail,
		Phone:    student.Phone,
	}, nil
}

func decodeGetStudents(ctx context.Context, request2 *http.Request) (request interface{}, err error) {
	var student model.Student
	return student, err
}

func decodeDeleteStudent(ctx context.Context, request2 *http.Request) (request interface{}, err error) {
	vars := mux.Vars(request2)

	studentId, ok := vars["studentId"]

	if !ok {
		logrus.Warnf("DecodeDeleteStudent() - error while getting vars %v", ok)
	}

	res := endpoints.StudentDeleteRequest{
		StudentId: studentId,
	}
	return res, err
}

func decodeGetByIdStudent(ctx context.Context, request2 *http.Request) (request interface{}, err error) {
	vars := mux.Vars(request2)

	studentId, ok := vars["studentId"]
	if !ok {
		logrus.Warnf("DecodeGetByStudent() - error while getting vars %v", ok)
	}

	res := endpoints.GetStudentByIdRequest{
		StudentId: studentId,
	}
	return res, nil
}

func decodeCreateStudent(ctx context.Context, request2 *http.Request) (request interface{}, err error) {

	var student endpoints.CreateStudentRequest

	body, err := io.ReadAll(request2.Body)

	if err != nil {
		logrus.Warnf("Decode () - nill error in body %v", err)
	}
	logrus.Debugf("Decode () - transport add organization - request body: %s", string(body))

	err = json.Unmarshal(body, &student)

	if err != nil {
		logrus.Warnf("Decode () - Error while unmarshaling %v ", err)
	}
	logrus.Debugf("Decode () -transport add organization incoming object: %v", student.Student)

	request = student
	err = nil

	return
}

func encodeStudent(ctx context.Context, writer http.ResponseWriter, response interface{}) error {
	if err, ok := response.(error); ok && err != nil {
		logrus.Warnf("Encode ()  - error %v", ok)
	}
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	return json.NewEncoder(writer).Encode(response)
}
