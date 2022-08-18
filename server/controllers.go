package server

import (
	"com/manveer/manager/container_service"
	"com/manveer/manager/host_service"
	"com/manveer/manager/model"
)

// Registry for all the controllers in the application
var Controllers = []model.Controller{
	{Path: container_service.BasePath, Routes: container_service.Routes},
	{Path: host_service.BasePath, Routes: host_service.Routes},
}
