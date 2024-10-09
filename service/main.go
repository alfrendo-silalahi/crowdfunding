package main

import (
	"fmt"
	"log"
	"service/user"
	"strings"

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
	// port := viper.GetString("server.port")

	dsn := fmt.Sprintf("%s:%s@tcp%s?charset=utf8mb4&parseTime=True&loc=Local", dbUsername, dbPassword, dbUrl)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Connect to database successfully")

	var userRepository user.Repository = user.NewRepository(db)
	var userService user.Service = user.NewService(userRepository)
	registerUserRequest := user.RegisterUserRequest{
		Name:       "Alan Walker",
		Occupation: "Music Creator",
		Email:      "alanwalker@gmail.com",
		Password:   "alanwalker",
	}

	userService.RegisterUser(registerUserRequest)
}
