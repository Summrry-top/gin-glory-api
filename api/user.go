package api

import (
	"github.com/Summrry-top/gin-glory-api/models/orm"
	"github.com/Summrry-top/gin-glory-api/services"
	"github.com/Summrry-top/gin-glory-api/utils"
	"github.com/gin-gonic/gin"
)

func GetUser(c *gin.Context) {
	g := services.GetParamDefault()
	_ = c.ShouldBind(&g)
	if !utils.VerifyParam(g) {
		Err400(c)
		return
	}
	isAdmin := c.GetBool("isAdmin")
	if utils.Zero(g.Id) {
		Success(c, services.GetUserAll(isAdmin, g))
		return
	}
	Success(c, services.GetUserById(isAdmin, g.Id))
}

func PostUser(c *gin.Context) {
	mode := c.GetString("mode")
	if mode == "delete" {
		id := c.PostForm("id")
		Success(c, services.PostUserDelete(id))
		return
	}
	var o orm.User
	if err := c.ShouldBind(&o); err != nil || o.Id < 0 {
		Err400(c)
		return
	}
	switch mode {
	case "login":
		Success(c, services.Login(&o))
	case "register":
		code := c.PostForm("code")
		Success(c, services.Register(&o, code))
	case "edit":
		isAdmin := c.GetBool("isAdmin")
		if !isAdmin {
			Success(c, "无权限")
		}
		Success(c, services.PostUserEdit(&o))
	default:
		Err400(c)
	}
	//var p request.PostUserParam
	//if err:=c.ShouldBind(&p);err!=nil{
	//	Err400(c)
	//	return
	//}
	//session,_:=global.Store.Get(c.Request,"session-code")
	//data:=services.PostUser(&p,session)
	//Success(c,data)
}
