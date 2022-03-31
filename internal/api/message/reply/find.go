package reply

import (
	"github.com/gin-gonic/gin"
	"messageBox/internal/err"
	"messageBox/internal/model"
	"messageBox/internal/repository/database"
	"net/http"
)

func FindByID(ctx *gin.Context) {
	var messages []model.Message
	var replies []model.Reply
	messageId := ctx.Param("messageId")

	result := database.DB.Where(&model.Reply{FromMessageID: messageId}).Find(&replies)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	if result.RowsAffected == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.MessageNotFoundError.Error()})
		return
	}

	var msgIds []string
	for _, reply := range replies {
		msgIds = append(msgIds, reply.ToMessageID)
	}

	result = database.DB.Find(&messages, msgIds)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": messages})
}
