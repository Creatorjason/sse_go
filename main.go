package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Creatorjason/sse_go/api"
)

func main(){
	eng := gin.Default()
	apiServer :=api.NewApiServer(eng)
	apiServer.RunServer()

}