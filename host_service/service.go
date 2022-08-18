package host_service

import (
	"com/manveer/manager/database"
	"com/manveer/manager/model"
	"errors"
	"fmt"
)

func GetAllHosts() (*[]model.Host, error) {
	result := make([]model.Host, 0)
	rows, err := database.Get().Query("SELECT id, uuid, name, ip_address from hosts;")
	if err == nil {
		for rows.Next() {
			host := model.Host{}
			scanErr := rows.Scan(&host.Id, &host.Uuid, &host.Name, &host.Ip)
			if scanErr != nil {
				rows.Close()
				return nil, fmt.Errorf("error occured while fetching hosts\n%s", err.Error())
			}
			result = append(result, host)
		}
		rows.Close()
	} else {
		return nil, fmt.Errorf("error occured while fetching hosts\n%s", err.Error())
	}
	return &result, nil
}

func GetHostById(id int) (*model.Host, error) {
	result := &model.Host{}
	stmt, _ := database.Get().Prepare("SELECT id, uuid, name, ip_address from hosts where id = ?")
	row, err := stmt.Query(id)
	stmt.Close()
	if err == nil {
		if row.Next() {
			scanErr := row.Scan(&result.Id, &result.Uuid, &result.Name, &result.Ip)
			row.Close()
			if scanErr != nil {
				return nil, fmt.Errorf("error occured while fetching host\n%s", err.Error())
			}
		} else {
			return nil, errors.New("host not found")
		}
	}
	return result, nil
}
