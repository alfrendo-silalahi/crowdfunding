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
	// read configuration
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err.Error())
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// database connection
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

	// instantiation dependencies
	var userRepository user.Repository = user.NewRepository(db)
	var userService user.Service = user.NewService(userRepository)

	user, err := userService.Login(user.LoginRequest{
		Email:    "alfrendo.silalahi@email.com",
		Password: "password",
	})

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("Berhasil menemukan user")
	fmt.Println(user)

	userHandler := handler.NewUserHandler(userService)

	// routing
	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)

	router.Run(port)
}
