package main

import (
	"fmt"
	"log"
	"service/handler"
	"service/user"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	dbUrl := viper.GetString("database.url")
	dbUsername := viper.GetString("database.username")
	dbPassword := viper.GetString("database.password")
	port := viper.GetString("server.port")

	dsn := fmt.Sprintf("%s:%s@tcp%s?charset=utf8mb4&parseTime=True&loc=Local", dbUsername, dbPassword, dbUrl)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Connect to database successfully")

	var userRepository user.Repository = user.NewRepository(db)
	var userService user.Service = user.NewService(userRepository)
	userHandler := handler.NewUserHandler(userService)

	user, err := userRepository.FindByEmail("aaaaa")
	if err != nil {
		fmt.Println("user record not found!")
		return
	}
	fmt.Println(user)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)

	router.Run(port)
}
