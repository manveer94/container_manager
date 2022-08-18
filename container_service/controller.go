package container_service

import (
	"com/manveer/manager/model"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func cGetAll(context *gin.Context) {
	hosts, err := GetAllContainers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
	context.IndentedJSON(http.StatusOK, hosts)
}

func cGetById(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid id provided",
		})
	} else {
		host, err2 := GetContainerById(id)

		if err2 == nil {
			context.IndentedJSON(http.StatusOK, *host)
		} else {
			context.JSON(http.StatusNotFound, gin.H{
				"message": err2.Error(),
			})
		}
	}
}

func cGetAllByHostId(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid id provided",
		})
	} else {
		hosts, err2 := GetContainersByHostId(id)

		if err2 == nil {
			context.IndentedJSON(http.StatusOK, hosts)
		} else {
			context.JSON(http.StatusNotFound, gin.H{
				"message": err2.Error(),
			})
		}
	}
}

func cCreateNewContainer(context *gin.Context) {
	var newContainer model.Container

	if err := context.BindJSON(&newContainer); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Error while parsing the body",
		})
		log.Panic(err.Error())
		return
	}
	if len(strings.Trim(newContainer.ImageName, " ")) == 0 {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Image name cannot be empty",
		})
		return
	}
	err2 := CreateNewContainer(&newContainer)
	if err2 != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": err2.Error(),
		})
		return
	}
	context.IndentedJSON(http.StatusOK, newContainer)
}
