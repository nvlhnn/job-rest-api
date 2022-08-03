package controller

import (
	"dansmultipro/recruitment/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type JobController interface {
	GetList(ctx *gin.Context)
	GetById(ctx *gin.Context)
}

type jobController struct {
	jobService service.JobService
}

func JobHandler(jobService service.JobService) JobController {
	return &jobController{jobService}
}


func (c *jobController) GetList(ctx *gin.Context) {
	query := ctx.Request.URL.Query()

	res := c.jobService.GetList(query)
	ctx.JSON(http.StatusOK, res)	
}


func (c *jobController) GetById(ctx *gin.Context) {
	id := ctx.Param("id") 
	res := c.jobService.GetById(id)
	ctx.JSON(http.StatusOK, res)	
}

