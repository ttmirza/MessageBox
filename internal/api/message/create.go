package message

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"messageBox/internal/model"
	"messageBox/internal/repository/database"

	"net/http"
	"time"
)

type Request struct {
	Sender    string    `json:"sender" binding:"required"`
	Recipient Recipient `json:"recipient" binding:"required"`
	Subject   string    `json:"subject" binding:"required"`
	Body      string    `json:"body" binding:"required"`
}

type Recipient struct {
	Groupname string `json:"groupname""`
	Username  string `json:"username""`
}

func Create(context *gin.Context) {
	var req Request
	if err := context.ShouldBindJSON(&req); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	message := model.Message{
		ID:          uuid.NewString(),
		Sender:      req.Sender,
		Body:        req.Body,
		Subject:     req.Subject,
		CreatedDate: time.Now(),
	}
	if req.Recipient.Groupname != "" {
		message.RecipientType = "group"
		message.RecipientName = req.Recipient.Groupname
	} else if req.Recipient.Username != "" {
		message.RecipientType = "user"
		message.RecipientName = req.Recipient.Username
	}

	result := database.DB.Create(&message)
	if result.Error != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// passed above validations, add to lookup table
	context.JSON(http.StatusCreated, gin.H{"data": message})
	return
}
