package utils

import (
	"cifs/dashboard/control"
	"github.com/gin-gonic/gin"
)

// Router 路由配置
type Router struct {
}

// Register 路由注册
func (r *Router) Register(g *gin.Engine, modules []func(*gin.Engine)) {
	for _, module := range modules {
		module(g)
	}
}

func Route(g *gin.Engine) {
	groupAdmin := g.Group("/")
	{

		groupIndex := groupAdmin.Group("/index")
		{
			ctlIndex := control.Index{}
			groupIndex.GET("/index", ctlIndex.Index)
		}
	}
}
