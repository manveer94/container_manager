package host_service

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func cGetAllHosts(context *gin.Context) {
	hosts, err := GetAllHosts()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
	}
	context.IndentedJSON(http.StatusOK, hosts)
}

func cGetHostById(context *gin.Context) {
	id, err := strconv.Atoi(context.Param("id"))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid id provided",
		})
	} else {
		host, err2 := GetHostById(id)

		if err2 == nil {
			context.IndentedJSON(http.StatusOK, *host)
		} else {
			context.JSON(http.StatusNotFound, gin.H{
				"message": err2.Error(),
			})
		}
	}
}
