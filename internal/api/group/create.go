package group

import (
	"github.com/gin-gonic/gin"
	"messageBox/internal/err"
	"messageBox/internal/model"
	"messageBox/internal/repository/database"
	"net/http"
	"time"
)

type Request struct {
	Name      string   `json:"groupname" binding:"required"`
	Usernames []string `json:"usernames" binding:"required"`
}

func Create(ctx *gin.Context) {
	var req Request
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	group := model.Group{
		Name:        req.Name,
		CreatedDate: time.Now(),
	}

	result := database.DB.Create(&group)
	if result.Error != nil {
		if result.Error.Error() == "UNIQUE constraint failed: groups.name" {
			ctx.JSON(http.StatusConflict, gin.H{"error": err.DuplicateGroupError.Error()})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// add user to the UserGroup table to show relationship between group & user
	for _, username := range req.Usernames {
		groupUser := model.GroupUser{
			Name:     req.Name,
			Username: username,
		}
		result = database.DB.Create(&groupUser)
		if result.Error != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
	}

	// passed above validations, add to lookup table
	ctx.JSON(http.StatusCreated, gin.H{"data": group})
	return
}
