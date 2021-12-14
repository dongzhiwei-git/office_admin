package main

import (
	"fmt"
	"inherited/internal"
	"inherited/internal/dao"
	"inherited/internal/models"
	"inherited/internal/services"
	"log"
	"testing"
)

func TestCreateSysUser(t *testing.T) {

	if err := internal.Init(); err != nil {
		log.Println("Init failed." + err.Error())
		return
	}

	SysUser := new(services.SysUser)
	SysUser.CreateSysUser("rot", "2355")
	fmt.Println("er")

}

func TestUser(t *testing.T) {
	fmt.Println("er")
}

func TestAttendanceInfo(t *testing.T) {
	if err := internal.Init(); err != nil {
		log.Println("Init failed." + err.Error())
		return
	}
	AttendanceInfo := new([]models.AttendanceInfo)
	err := dao.Orm.Find(AttendanceInfo).Error

	if err != nil {
		fmt.Println("errrrrr")
		return
	}

	fmt.Println(AttendanceInfo)
}

func TestFileInfo(t *testing.T) {
	if err := internal.Init(); err != nil {
		log.Println("Init failed." + err.Error())
		return
	}
	FileInfo := new([]models.FileInfo)
	err := dao.Orm.Find(FileInfo).Error

	if err != nil {
		fmt.Println("errrrrr")
		return
	}

	fmt.Println(FileInfo)
}
