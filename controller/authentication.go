package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hazaloolu/GoRestApi/model"
)

func Register(context *gin.Context) {
	var input model.AuthenticationInput

	if err := context.ShouldBindBodyWithJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := model.User{
		Username: input.Username,
		Password: input.Password,
	}

	savedUser, err := Save()

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"user": savedUser})

}
