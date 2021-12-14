package api

import (
	"fmt"
	"log"
	"net/http"

	"inherited/internal/models"
	"inherited/internal/services"

	"github.com/gin-gonic/gin"
)

func CreateAdminUser(ctx *gin.Context) {
	//Parameter parsing
	adminUser := models.SysUser{}
	err := ctx.BindJSON(&adminUser)
	if err != nil {
		fmt.Printf("[api.CreateAdminUser], Parameter parsing error")
	}

	userName := adminUser.UserName
	password := adminUser.Password
	if userName == "" || password == "" {
		log.Printf("userName or password is null")

		return
	}

	sysUser := new(services.SysUser)
	err = sysUser.CreateSysUser(adminUser.UserName, adminUser.Password)
	if err != nil {
		fmt.Printf("[api.CreateAdminUser], err: %v", err)

		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"date": "success",
	})

	return
}

func GetAttendanceInfo(ctx *gin.Context) {
	//Parameter parsing
	info := models.AttendanceInfo{}
	err := ctx.ShouldBind(&info)
	if err != nil {
		fmt.Printf("[api.GetAttendanceInfo], Parameter parsing error")
	}

	attendance := new(services.Attendance)
	attendanceInfo, err := attendance.GetAttendanceInfo()
	if err != nil {
		fmt.Printf("[api.GetAttendanceInfo], err: %v", err)

		return
	}
	fmt.Println(attendanceInfo)
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"date":   attendanceInfo,
	})

	return
}

func GetFileInfo(ctx *gin.Context) {
	fileInfo := new(services.File)
	info, err := fileInfo.GetFileInfo()
	if err != nil {
		fmt.Printf("[api.GetFileInfo], err: %v", err)

		return
	}
	fmt.Println(info)
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"date":   info,
	})

	return
}

func GetMeetingInfo(ctx *gin.Context) {
	MeetingRecord := new(services.MeetingRecord)
	info, err := MeetingRecord.GetFileInfo()
	if err != nil {
		fmt.Printf("[api.GetMeetingInfo], err: %v", err)

		return
	}
	fmt.Println(info)
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"date":   info,
	})

	return
}

func GetRoutineInfo(ctx *gin.Context) {
	RoutineInfo := new(services.Routine)
	info, err := RoutineInfo.GetRoutineInfo()
	if err != nil {
		fmt.Printf("[api.GetRoutineInfo], err: %v", err)

		return
	}
	fmt.Println(info)
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"date":   info,
	})

	return
}

func GetAdminUserInfo(ctx *gin.Context) {
	adminUser := new(services.SysUser)
	info, err := adminUser.GetUserInfo()
	if err != nil {
		fmt.Printf("[api.GetAdminUserInfo], err: %v", err)

		return
	}
	fmt.Println(info)
	ctx.JSON(http.StatusOK, gin.H{
		"status": "success",
		"date":   info,
	})

	return
}