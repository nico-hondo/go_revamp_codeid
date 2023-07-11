package server

import (
	"database/sql"
	"log"

	"codeid.revamptwo/controllers"
	"codeid.revamptwo/repositories"
	"codeid.revamptwo/services"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config               *viper.Viper
	router               *gin.Engine
	departmentController *controllers.DepartmentController
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
	departmentRepository := repositories.NewDepartmentRepository(dbHandler)
	departmentServices := services.NewDepartmentService(departmentRepository)
	departmentController := controllers.NewDepartmentController(departmentServices)

	router := gin.Default()
	// router endpoint/url http category
	router.GET("/department", departmentController.GetListDepartment)
	router.GET("/department/:id", departmentController.GetDepartment)
	router.POST("/department", departmentController.CreateDepartment)
	router.PUT("/department/:id", departmentController.UpdateDepartment)
	router.DELETE("/department/:id", departmentController.DeleteDepartment)

	// router

	return HttpServer{
		config:               config,
		router:               router,
		departmentController: departmentController,
	}
}

// method running gin httpserver
func (hs HttpServer) Start() {
	err := hs.router.Run(hs.config.GetString("http.server_address"))
	if err != nil {
		log.Fatalf("Error while starting HTTP Server : %v ", err)
	}
}
