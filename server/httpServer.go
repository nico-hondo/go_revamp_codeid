package server

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

type HttpServer struct {
	config               *viper.Viper
	router               *gin.Engine
	departmentController *controllers.DepartmentController
}

func InitHttpServer(config *viper.Viper, dbHandler *sql.DB) HttpServer {
	categoryRepository := repositories.NewCategoryRepository(dbHandler)
	categoryServices := services.NewCategoryService(categoryRepository)
	categoryController := controllers.NewCategoryController(categoryServices)

	router := gin.Default()
	// router endpoint/url http category
	router.GET("/employee", departmentController.GetListEmployee)
	router.GET("/employee/:id", departmentController.GetEmployee)
	router.POST("/employee", departmentController.CreateEmployee)
	router.PUT("/employee/:id", departmentController.UpdateEmployee)
	router.DELETE("/employee/:id", departmentController.DeleteEmployee)

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
