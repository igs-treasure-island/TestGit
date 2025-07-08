package main

import (
	"fmt"

	"github.com/igs-treasure-island/utils/db"
	_ "github.com/microsoft/go-mssqldb"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

var (
	/* 資料庫 */
	backendDb *gorm.DB // 讀資料專用通道
	logDb *gorm.DB // 讀資料專用通道
)

func dbInit() {
	var err error
	backendDb, err = gorm.Open(sqlserver.Open(db.Get("BackendDb").SQLServerDSN("backend")), &gorm.Config{QueryFields: true, SkipDefaultTransaction: true})
	if err != nil {
		panic(fmt.Sprintf("cannot create engine. ERR=%s", err.Error()))
	}
	logDb, err = gorm.Open(sqlserver.Open(db.Get("EventDb").SQLServerDSN("EventLog")), &gorm.Config{QueryFields: true, SkipDefaultTransaction: true})
	if err != nil {
		panic(fmt.Sprintf("cannot create engine. ERR=%s", err.Error()))
	}
}
