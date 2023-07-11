package transport

import (
	"context"
	"encoding/json"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/saurabhkanawade/studentmanager/internal/endpoints"
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
	logrus.Debugf("Decode () -transport add organization incoming object: %v", student)

	request = student
	err = nil

	return
}

func encodeStudent(ctx context.Context, writer http.ResponseWriter, response interface{}) error {
	if err, ok := response.(error); ok && err != nil {
		logrus.Warnf("Encode ()  - error %v", ok)
	}
	writer.Header().Set("Content-Type", "application/json; charset=utf-8")
	writer.WriteHeader(http.StatusCreated)
	return json.NewEncoder(writer).Encode(response)
}
