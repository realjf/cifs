package main

import (
	"cifs/dashboard/utils"
	"cifs/service/config"
	"cifs/service/db"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
)

var (
	Config *config.Config
)


func init() {
	Config = config.NewConfig().LoadConfig("../service/config/config.json")
	db.NewMysql(Config).Init()
}


func main() {
	r := gin.New()

	new(utils.Router).Register(r, []func(engine *gin.Engine){
		utils.Route,
	})

	serve := &http.Server{
		Addr:         ":8001",
		Handler:      r,
		ReadTimeout:  20 * time.Second,
		WriteTimeout: 20 * time.Second,
	}

	log.Println("Listen On: " + serve.Addr)
	log.Fatal(serve.ListenAndServe())
}


