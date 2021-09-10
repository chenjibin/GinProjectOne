package main

import (
	"GinProjectOne/common"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"os"
)

func main() {
	initConfig()
	common.InitDB()
	r := gin.Default()
	r.StaticFS("/static", http.Dir("./static"))
	CollectRoute(r)
	//err := r.Run(":8088")
	err := r.RunTLS(":8088", "server.cer", "server.key")
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
