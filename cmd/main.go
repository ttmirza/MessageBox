package main

import (
	"github.com/gin-gonic/gin"
	"messageBox/internal/api/group"
	"messageBox/internal/api/message"
	"messageBox/internal/api/message/reply"
	"messageBox/internal/api/user"
	"messageBox/internal/repository/database"
)

func main() {
	r := gin.Default()
	database.ConnectDatabase()

	r.POST("/users", user.Create)
	r.GET("/users", user.Find)
	r.GET("/users/:username/mailbox", message.Find)

	r.POST("/groups", group.Create)

	r.POST("/messages", message.Create)
	r.GET("/messages/:messageId", message.FindByID)

	r.POST("/messages/:messageId/replies", reply.Create)
	r.GET("/messages/:messageId/replies", reply.FindByID)

	r.Run(":3001")
}
