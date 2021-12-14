package router

import (
	"fmt"
	"inherited/internal/api"
	"inherited/internal/pkg"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter() {
	var r *gin.Engine
	r = gin.Default()
	r.Use(pkg.Cors())

	// to solve the cross domain
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, nil)
	})

	var admin = r.Group("/admin")
	{
		admin.POST("/reg", api.CreateAdminUser)
		admin.GET("/user", api.GetAdminUserInfo)
		admin.GET("/attendance-info", api.GetAttendanceInfo)
		admin.GET("/file-info", api.GetFileInfo)
		admin.GET("/meeting-info", api.GetMeetingInfo)
		admin.GET("/office-routine-info", api.GetRoutineInfo)
	}

	// setup listen
	err := r.Run(":8001")
	if err != nil {
		fmt.Printf("run failed: %v\n", err)
		return

	}

}
