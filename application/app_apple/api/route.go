package api

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go_project_template/application/appA/proto"
	"go_project_template/application/appA/repo"
	"go_project_template/application/appA/service"
	"go_project_template/utils"
	"gorm.io/gorm"
	"io"
	"log"
)

var serv service.UserService

func Init(userDB *gorm.DB) *gin.Engine {
	userRepo := repo.NewUserRepository(userDB)
	serv = service.NewUserService(userRepo)

	// TODO: add init funcs belong to appA

	hdr := new(handler)
	r := gin.Default()
	register(r, "/", hdr.GetUser, &proto.GetUserReq{})
	// register other handler
	return r
}

func register(r *gin.Engine, path string, handlerFunc gin.HandlerFunc, req interface{}) {
	hdr := func(ctx *gin.Context) {
		b, _ := io.ReadAll(ctx.Request.Body)
		err := json.Unmarshal(b, req)
		if err != nil {
			log.Printf("register parse req<%T> fail, err=%s", req, err.Error())
			utils.EJSON(ctx, "非法参数")
			return
		}
		ctx.Set("req", req) // 提前设置req对象到handler，后者无需再解析req
		handlerFunc(ctx)
	}
	r.POST(path, hdr)
}
