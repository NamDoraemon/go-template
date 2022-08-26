package services

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"urbox-viettel-go/entities"
)

func (service Services) GetCodeController(ctx *gin.Context) {
}

func (service Services) HoldCodeController(ctx *gin.Context) {
	requestHoldCode := &entities.RequestHoldCode{}
	if err := ctx.BindJSON(&requestHoldCode); err != nil {
		log.Println(err)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"error":   "VALIDATEERR-1",
			"message": "Invalid inputs. Please check your inputs",
		})
		return
	}
	log.Println(requestHoldCode)
}
