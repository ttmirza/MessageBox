package reply

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"messageBox/internal/err"
	"messageBox/internal/model"
	"messageBox/internal/repository/database"
	"net/http"
	"time"
)

type Request struct {
	Sender  string `json:"sender" binding:"required"`
	Subject string `json:"subject" binding:"required""`
	Body    string `json:"body" binding:"required"`
}

func Create(ctx *gin.Context) {
	var req Request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	messageId := ctx.Param("messageId")
	var message model.Message
	result := database.DB.Where(&model.Message{ID: messageId}).Find(&message)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.MessageNotFoundError.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	var reply model.Message
	var replyRelationship model.Reply
	newUuid := uuid.NewString()
	reply = model.Message{
		ID:            newUuid,
		Sender:        req.Sender,
		RecipientType: message.RecipientType,
		RecipientName: message.Sender,
		Body:          req.Body,
		Subject:       req.Subject,
		CreatedDate:   time.Now(),
	}
	replyRelationship = model.Reply{
		FromMessageID: message.ID,
		ToMessageID:   newUuid,
	}
	if message.RecipientType == "group" {
		reply.RecipientName = message.RecipientName
	}
	result = database.DB.Create(&reply)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	result = database.DB.Create(&replyRelationship)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// passed above validations, add to lookup table
	ctx.JSON(http.StatusCreated, gin.H{"data": reply})
	return

}
