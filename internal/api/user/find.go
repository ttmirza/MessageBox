package user

import (
	"github.com/gin-gonic/gin"
	"messageBox/internal/model"
	"messageBox/internal/repository/database"
	"net/http"
)

func Find(context *gin.Context) {
	var Users []model.User
	database.DB.Find(&Users)
	context.JSON(http.StatusOK, gin.H{"data": Users})
}
