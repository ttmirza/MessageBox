package message

import (
	"github.com/gin-gonic/gin"
	"messageBox/internal/err"
	"messageBox/internal/model"
	"messageBox/internal/repository/database"
	"net/http"
)

func Find(ctx *gin.Context) {
	var userMessages []model.Message
	var groupMessages []model.Message
	var allMessages []model.Message
	var user model.User
	username := ctx.Param("username")

	result := database.DB.Where(&model.User{Username: username}).Find(&user)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.UserNotFoundError.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// TODO: update LastLoggedIn Field in User Data Model indicating a login, since at this stage we have validated the user exists

	result = database.DB.Where(&model.Message{RecipientName: username}).Find(&userMessages)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
	}

	// find if any groups related, so that we can pull group messages
	var myGroups []model.GroupUser
	result = database.DB.Where(&model.GroupUser{Username: username}).Find(&myGroups)
	if result.Error != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
	}
	if len(myGroups) > 0 {
		var groupNames []string
		for _, group := range myGroups {
			groupNames = append(groupNames, group.Name)
		}

		// find all group messages
		result = database.DB.Where("recipient_name IN (?)", groupNames).Find(&groupMessages)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
	}

	// do we return a 404 when no message in mailbox?
	if len(userMessages) == 0 && len(groupMessages) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.MessageNotFoundError})
		return
	}

	allMessages = append(allMessages, userMessages...)
	allMessages = append(allMessages, groupMessages...)

	ctx.JSON(http.StatusOK, gin.H{"data": allMessages})
}

func FindByID(ctx *gin.Context) {
	var message model.Message
	messageId := ctx.Param("messageId")

	result := database.DB.Where(&model.Message{ID: messageId}).Find(&message)
	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.MessageNotFoundError.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": message})
}
