package server

import (
	"com/manveer/manager/model"
	"fmt"
	"log"
	"strings"

	"github.com/gin-gonic/gin"
)

// Creates all the routes for the application by fetching all the controllers. All the routes are grouped by controller
func buildRoutes(router *gin.Engine) {
	for _, controller := range Controllers {
		group := router.Group(checkAndSlash(controller.Path))
		for _, route := range controller.Routes {
			createRoute(&route, group)
		}
	}

}

// Create the route group for each controller
func createRoute(route *model.Route, group *gin.RouterGroup) {
	path := checkAndSlash(route.Path)

	if route.Method == "GET" {
		group.GET(path, route.Callback)
	} else if route.Method == "PUT" {
		group.PUT(path, route.Callback)
	} else if route.Method == "POST" {
		group.POST(path, route.Callback)
	} else {
		log.Fatalf("Invalid method type provided for path %s : %s", path, route.Method)
	}
}

// Adds slash in the begining of the path if not exists.
func checkAndSlash(str string) string {
	if !strings.HasPrefix(str, "/") {
		str = "/" + str
	}
	return str
}

// Starts the server
func Run(port int) {
	router := gin.Default()
	buildRoutes(router)
	router.Run(fmt.Sprintf("localhost:%d", port))
}
