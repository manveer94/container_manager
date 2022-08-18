package host_service

import "com/manveer/manager/model"

var BasePath = "/host"

var Routes = []model.Route{
	{Path: "/all", Method: "GET", Callback: cGetAllHosts},
	{Path: "/:id", Method: "GET", Callback: cGetHostById},
}
