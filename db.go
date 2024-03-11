package main

import (
	"fmt"
	"os"

	"github.com/igs-treasure-island/utils/db"
	_ "github.com/microsoft/go-mssqldb"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var (
	/* 資料庫 */
	backendDb *gorm.DB // 讀資料專用通道
)

func dbInit() {
	backendDbInfo := db.Get("BackendDb")
	var err error
	backendDb, err = gorm.Open(sqlserver.Open("server="+os.Getenv("backend")+";user id="+backendDbInfo.Username+";password="+backendDbInfo.Password+";database=backend"), &gorm.Config{QueryFields: true, SkipDefaultTransaction: true})
	if err != nil {
		panic(fmt.Sprintf("cannot create engine. ERR=%s", err.Error()))
	}
}
