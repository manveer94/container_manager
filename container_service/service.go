package container_service

import (
	"com/manveer/manager/database"
	"com/manveer/manager/host_service"
	"com/manveer/manager/model"
	"com/manveer/manager/utils"
	"errors"
	"fmt"
)

func GetAllContainers() (*[]model.Container, error) {
	result := make([]model.Container, 0)
	rows, err := database.Get().Query("SELECT c.id , c.host_id , c.name , c.image_name , h.name as host_name  from containers c , hosts h WHERE h.id == c.host_id;")
	if err == nil {
		for rows.Next() {
			container := model.Container{}
			scanErr := rows.Scan(&container.Id, &container.HostId, &container.Name, &container.ImageName, &container.HostName)
			if scanErr != nil {
				rows.Close()
				return nil, fmt.Errorf("error occured while fetching containers\n%s", err.Error())
			}
			result = append(result, container)
		}
		rows.Close()
	} else {
		return nil, fmt.Errorf("error occured while fetching containers\n%s", err.Error())
	}
	return &result, nil
}

func GetContainerById(id int) (*model.Container, error) {
	result := &model.Container{}
	stmt, _ := database.Get().Prepare("SELECT c.id , c.host_id , c.name , c.image_name , h.name as host_name  from containers c , hosts h WHERE h.id == c.host_id  and c.id = ?")
	row, err := stmt.Query(id)
	stmt.Close()
	if err == nil {
		if row.Next() {
			scanErr := row.Scan(&result.Id, &result.HostId, &result.Name, &result.ImageName, &result.HostName)
			row.Close()
			if scanErr != nil {
				return nil, fmt.Errorf("error occured while fetching container\n%s", err.Error())
			}
		} else {
			return nil, errors.New("container not found")
		}
	}
	return result, nil
}

func GetContainersByHostId(id int) (*[]model.Container, error) {
	result := make([]model.Container, 0)
	stmt, _ := database.Get().Prepare("SELECT c.id , c.host_id , c.name , c.image_name , h.name as host_name  from containers c , hosts h WHERE h.id == c.host_id  and h.id = ?")
	rows, err := stmt.Query(id)
	stmt.Close()
	if err == nil {
		for rows.Next() {
			container := model.Container{}
			scanErr := rows.Scan(&container.Id, &container.HostId, &container.Name, &container.ImageName, &container.HostName)
			if scanErr != nil {
				rows.Close()
				return nil, fmt.Errorf("error occured while fetching containers\n%s", err.Error())
			}
			result = append(result, container)
		}
		rows.Close()
	} else {
		return nil, fmt.Errorf("error occured while fetching containers\n%s", err.Error())
	}
	return &result, nil
}

func CreateNewContainer(container *model.Container) error {
	host, err := host_service.GetHostById(container.HostId)
	if err != nil {
		return err
	}
	container.Name = utils.GenerateUuid()
	stmt, _ := database.Get().Prepare("INSERT INTO containers (host_id, name, image_name) VALUES(?, ?, ?);")
	result, err := stmt.Exec(container.HostId, container.Name, container.ImageName)
	stmt.Close()
	if err != nil {
		return fmt.Errorf("error occured while creating image\n%s", err.Error())
	}
	id, _ := result.LastInsertId()
	container.Id = int(id)
	container.HostName = host.Name
	return nil
}
