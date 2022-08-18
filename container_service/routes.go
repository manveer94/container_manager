package container_service

import "com/manveer/manager/model"

var BasePath = "/container"
var Routes = []model.Route{
	{Path: "/all", Method: "GET", Callback: cGetAll},
	{Path: "/:id", Method: "GET", Callback: cGetById},
	{Path: "/all/host/:id", Method: "GET", Callback: cGetAllByHostId},
	{Path: "", Method: "PUT", Callback: cCreateNewContainer},
}
