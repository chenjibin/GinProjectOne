package main

import (
	"GinProjectOne/common"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"os"
)

func main() {
	initConfig()
	common.InitDB()
	r := gin.Default()
	CollectRoute(r)
	err := r.Run(":8088")
	if err != nil {
		return
	}
}

func initConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}