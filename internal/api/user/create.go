package user

import (
	"github.com/gin-gonic/gin"
	"messageBox/internal/err"
	"messageBox/internal/model"
	"messageBox/internal/repository/database"
	"net/http"
	"time"
)

type Request struct {
	Username string `json:"username" binding:"required"`
}

func Create(ctx *gin.Context) {
	var req Request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := model.User{
		Username:     req.Username,
		CreatedDate:  time.Now(),
		LastLoggedIn: time.Now(),
	}
	result := database.DB.Create(&user)
	if result.Error != nil {
		if result.Error.Error() == "UNIQUE constraint failed: users.username" {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.DuplicateUserError.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": user})
	return
}
