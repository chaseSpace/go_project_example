package api

import (
	"github.com/gin-gonic/gin"
	"go_project_template/application/appA/proto"
	"go_project_template/utils"
)

type handler struct {
}

// api层调用不同service层接口做数据聚合
func (handler) GetUser(ctx *gin.Context) {
	req := ctx.Value("req").(*proto.GetUserReq)
	core, err := serv.GetUser(ctx, req.Uid)
	if err != nil {
		utils.EJSON(ctx, err.Error())
		return
	}
	other, err := serv.GetUserOther(ctx, req.Uid)
	if err != nil {
		utils.EJSON(ctx, err.Error())
		return
	}
	utils.JSON(ctx, &proto.GetUserRsp{Core: core, Other: other})
}
