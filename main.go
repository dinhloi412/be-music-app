package main

import (
	"be-music/model"
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=admin password=qwe123 dbname=postgres port=5438 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	User := &model.User{
		UserName:   "dinhloi",
		Password:   "123456",
		FirstName:  "FRANK",
		LastName:   "NGUYEN",
		Gender:     true,
		Disabled:   false,
		Updated_at: time.Now().Unix(),
		Created_at: time.Now().Unix(),
	}

	if err != nil {
		panic("failed to connect database")
	}
	log.Println(db, err)
	result := db.Create(User)
	fmt.Println(result)
}
