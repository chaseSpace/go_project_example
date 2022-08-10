package utils

import (
	"github.com/gin-gonic/gin"
	"go_project_template/application/appA/proto"
	"go_project_template/config"
	"gorm.io/gorm"
)

type getConf func(fpath string) (*config.AppConf, error)
type getDB func(*config.AppConf) (*gorm.DB, error)
type initApi func(*gorm.DB) *gin.Engine

func MustBuild(fpath string, getConf getConf, getDB getDB, initApi initApi) *gin.Engine {
	c, e := getConf(fpath)
	if e != nil {
		panic(e)
	}
	db, e := getDB(c)
	if e != nil {
		panic(e)
	}
	return initApi(db)
}

func EJSON(ctx *gin.Context, emsg string) {
	ctx.AbortWithStatusJSON(200, &proto.HttpRsp{
		Code: 400,
		Msg:  emsg,
		Data: nil,
	})
}

func JSON(ctx *gin.Context, data interface{}) {
	ctx.JSON(200, &proto.HttpRsp{
		Code: 200,
		Msg:  "操作成功",
		Data: data,
	})
}
