package main

import (
	"dansmultipro/recruitment/config"
	"dansmultipro/recruitment/controller"
	"dansmultipro/recruitment/middleware"
	"dansmultipro/recruitment/service"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {

	var db  *gorm.DB   = config.OpenConnection()

	var loginService service.LoginService = service.NewAuthService(db)
	var jwtService service.JWTService = service.NewJWTService()
	var loginController controller.LoginController = controller.LoginHandler(loginService, jwtService)

	var jobService service.JobService = service.NewJobService()
	var jobController controller.JobController = controller.JobHandler(jobService)

	
	defer config.CloseConnection(db)
	
	r := gin.Default()

	// cors
	r.Use(middleware.CORSMiddleware())

	r.POST("/login", loginController.Login)	

	// jwt
	r.Use(middleware.AuthorizeJWT(jwtService))	
	{
		r.GET("/jobs", jobController.GetList)
		r.GET("/jobs/:id", jobController.GetById)
	}



	r.Run() 
}