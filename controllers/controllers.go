package controllers

import(
	"github.com/gin-gonic/gin"
	"github.com/Creatorjason/sse_go/models"
	"net/http"
	"fmt"
	"time"
)

var (
	messages []models.Message
	clients []chan models.Message
)


func HandleReceiveMessage(c *gin.Context){
	// Set up SSE
	c.Header("Content-type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")


	for _, message := range messages{
		c.Writer.Write([]byte(fmt.Sprintf("ID: %s\n", message.ID)))
		c.Writer.Write([]byte(fmt.Sprintf("Sender: %s\n", message.Sender)))
		c.Writer.Write([]byte(fmt.Sprintf("Message: %s\n", message.Text)))
		c.Writer.Write([]byte(fmt.Sprintf("CreatedAt: %s\n", message.CreatedAt)))

		c.Writer.Write([]byte("-----------------------\n"))
	}
	c.Writer.Flush()

	// create new channel for user
	messageChan := make(chan models.Message)
	clients = append(clients, messageChan)
	for{
		select{
		case sseMsg :=  <- messageChan:
			c.Writer.Write([]byte(fmt.Sprintf("ID: %s\n", sseMsg.ID)))
            c.Writer.Write([]byte(fmt.Sprintf("Sender: %s\n", sseMsg.Sender)))
            c.Writer.Write([]byte(fmt.Sprintf("Message: %s\n", sseMsg.Text)))
            c.Writer.Write([]byte(fmt.Sprintf("Created At: %s\n", sseMsg.CreatedAt)))
            c.Writer.Write([]byte("-----------------------\n"))
            c.Writer.Flush()
		case <-c.Writer.CloseNotify():
			close(messageChan)
			clients = removeClient(clients, messageChan)
			return
		}
	}
}

func HandleSendMessage(c *gin.Context){
	var msg models.Message
	err := c.ShouldBind(&msg)
	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{
			"error": fmt.Sprintf("invalid data format provided :%v", err.Error()),
		})
	}

	msg.ID = fmt.Sprintf("%d",len(messages) + 1)
	msg.CreatedAt = time.Now()
	messages = append(messages, msg)

	for _, sseChan := range clients{
		sseChan <- msg
	}

	c.JSON(http.StatusOK, msg)

}
func removeClient(clients []chan models.Message, client chan models.Message) []chan models.Message {
    for i, c := range clients {
        if c == client {
            clients[i] = clients[len(clients)-1]
            return clients[:len(clients)-1]
        }
    }
    return clients
}