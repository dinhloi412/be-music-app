package main

import (
	"be-music/appctx"
	"be-music/middleware"
	"be-music/module/user/transport/ginuser"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=admin password=qwe123 dbname=postgres port=5438 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	appContext := appctx.NewAppContext(db)
	r := gin.Default()
	r.Use(middleware.Recover(appContext))
	r.POST("/test", ginuser.CreateUser(appContext))
	r.Run()
}
