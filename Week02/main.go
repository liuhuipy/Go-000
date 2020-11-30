package main

import (
	"Go-000/Week02/config"
	"Go-000/Week02/controller"
	"Go-000/Week02/initialize"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	configFilePath := flag.String("C", "./config.yaml", "config file path")
	flag.Parse()
	if err := config.LoadConfiguration(*configFilePath); err != nil {
		fmt.Println("failed parsing config file, err=", err)
		return
	}
	db, err := initialize.InitDB()
	if err != nil {
		fmt.Println("failed open database, err=", err)
		return
	}
	config.LoadDB(db)

	router := gin.Default()
	router.GET("/user/:id", controller.GetUser)
	router.Run(config.GetSystemConfiguration().Port)
}
