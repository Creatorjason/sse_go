package controllers

import(
	"github.com/gin-gonic/gin"
	"github.com/Creatorjason/sse_go/models"
	"net/http"
	"fmt"
)

var (
	messages []models.Message
	clients chan []models.Message
)


func HandleReceiveMessage(c *gin.Context){

	
}

func HandleSendMessage(c *gin.Context){
	var msg models.Message
	err := c.ShouldBind(&msg)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("invalid data format provided :%v", err.Error())
		})
	}

	msg.ID = len(message) + 1
	msg.CreatedAt = time.Now()
	messages = append(messages, msg)

	for _, sseChan := range client{
		sseChan <- msg
	}

	c.JSON(http.StatusOk, msg)

}