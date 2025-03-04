package Controllers

import (
	"2024_akutansi_project/Services"

	"github.com/gin-gonic/gin"
)

type (
	ILogActivityController interface {
		FindAll(ctx *gin.Context)
	}

	LogActivityController struct {
		logActivityService Services.ILogActivityService
	}
)

func LogActivityControllerProvider(logActivityService Services.ILogActivityService) *LogActivityController {
	return &LogActivityController{logActivityService: logActivityService}
}

func (l *LogActivityController) FindAll(ctx *gin.Context) {
	userID := ctx.GetString("id")

	logActivities, err := l.logActivityService.FindAll(userID)

	if err != nil {
		ctx.JSON(500, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"data": logActivities,
	})
}
