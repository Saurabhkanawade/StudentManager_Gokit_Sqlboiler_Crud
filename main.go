// Package student service API.
//
// # Endpoints for student
//
//	 StudentManager
//
//		Schemes: http , https
//		Host: localhost:8001
//		BasePath: /
//		Version: 0.0.1
//
// Consumes:
// - application/json
//
// Produces:
// - application/json
//
// swagger:meta
package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/saurabhkanawade/studentmanager/internal/config"
	"github.com/saurabhkanawade/studentmanager/internal/dao"
	"github.com/saurabhkanawade/studentmanager/internal/database"
	"github.com/saurabhkanawade/studentmanager/internal/endpoints"
	"github.com/saurabhkanawade/studentmanager/internal/services"
	"github.com/saurabhkanawade/studentmanager/internal/transport"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"time"
)

func main() {
	logrus.Info("Starting Employee Management Service....")

	//loading the config of env file
	config.LoadConfig(config.GetAppEnvLocation())

	//make sure all necessary variables are present
	if err := config.CheckRequiredVariables(); err != nil {
		logrus.Fatalf("Failed to fulfill CheckRequiredVariables Condition :%v", err)
	}

	//setup to Db
	dbConfig := database.DbConfig{
		Host:   config.GetPostgresHost(),
		Port:   config.GetPostgresPort(),
		User:   config.GetPostgresUser(),
		Pass:   config.GetPostgresPass(),
		DbName: config.GetPostgresDb(),
	}

	dbCon, err := database.InitDatabase(dbConfig)

	if err != nil {
		logrus.Fatalf("Error while establishing connection with databse %s ", err.Error())
	}

	err = dbCon.PingDB()
	if err != nil {
		log.Fatalf("Error pinging portals database: %s", err.Error())
	}
	router := mux.NewRouter()

	fs := http.FileServer(http.Dir("./swagger-ui"))
	router.PathPrefix("/swagger-ui").
		Handler(http.StripPrefix("/swagger-ui", fs))

	//set log level
	log := logrus.New()
	log.SetFormatter(&logrus.TextFormatter{DisableQuote: true}) //this will allow for multiline stack traces

	studentDao := dao.NewStudentDao(dbCon, log)
	studentService := services.NewStudentService(studentDao, log)
	studentEndpoint := endpoints.MakeStudentEndpoints(studentService)
	transport.CreateStudentHttpHandler(studentEndpoint, router)

	startServer(router)
}

// starting the server function
func startServer(router *mux.Router) {

	serverPort := fmt.Sprintf(":%s", config.GetServerPort())
	logrus.Infof("Starting server on %s.........", serverPort)
	readTimeout := time.Duration(config.GetReadTimeout())
	writeTimeout := time.Duration(config.GetWriteTimeout())

	server := &http.Server{
		Addr:         serverPort,
		ReadTimeout:  readTimeout * time.Second,
		WriteTimeout: writeTimeout * time.Second,
		Handler:      router,
	}

	if err := server.ListenAndServe(); err != nil {
		logrus.Fatalf("failed to start the server on %s %v", config.GetServerPort(), err)
	}

}
