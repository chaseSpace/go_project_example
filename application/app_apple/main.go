package main

import (
	"github.com/gin-gonic/gin"
	"go_project_template/application/appA/api"
	"go_project_template/config"
	"go_project_template/global/infra"
	"go_project_template/utils"
	"log"
)

func main() {
	gin.SetMode(gin.ReleaseMode)

	engine := utils.MustBuild("config/conf.json", config.NewConf, infra.NewDB, api.Init)
	err := engine.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	if err != nil {
		log.Println("server down, err=", err.Error())
	}

}
