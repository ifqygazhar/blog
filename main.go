package main

import (
	"blog/handler"
	"blog/user"
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {

	//refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:torabikA1#@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}
	userRepository := user.NewRepositoryUser(db)

	userService := user.NewServiceUser(userRepository)

	userHandler := handler.NewHandlerUser(userService)

	router := gin.Default()
	router.Use(cors.Default())
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.UserLogin)
	router.Run()

}
