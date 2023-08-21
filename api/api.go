package api

import (
	"github.com/gin-gonic/gin"
	"github.com/Creatorjason/sse_go/controllers"
)

type ApiServer struct{
	Router *gin.Engine
}

func NewApiServer(eng *gin.Engine) *ApiServer{
	return &ApiServer{
		Router: eng,
	}
}

func (ap *ApiServer) RunServer(){
	ap.Router.GET("/subscribe", controllers.HandleReceiveMessage)
	ap.Router.POST("/message", controllers.HandleSendMessage)
	ap.Router.Run()
}