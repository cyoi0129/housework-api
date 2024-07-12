package services

import (
	"fmt"
	"workout-note/models"
)

func FetchMasterList(user_id int) ([]models.Master, error) {
	var masters []models.Master
	rows, err := models.DB.Query("SELECT id, userID, name, category, point FROM \"masters\" WHERE userID = $1", user_id)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var master models.Master
		rows.Scan(&master.Id, &master.UserID, &master.Name, &master.Category, &master.Point)
		masters = append(masters, master)
	}
	fmt.Printf("%v", masters)
	return masters, nil
}

func CreateMaster(input models.Master) (models.Master, error) {
	master := models.Master{
		UserID:   input.UserID,
		Name:     input.Name,
		Category: input.Category,
		Point:    input.Point,
	}
	err := models.DB.QueryRow("INSERT INTO masters(userID, name, category, point) VALUES($1, $2, $3, $4) RETURNING id", master.UserID, master.Name, master.Category, master.Point).Scan(&master.Id)
	if err != nil {
		fmt.Println(err)
		return master, err
	}
	return master, nil
}

func UpdateMaster(master_id int, input models.Master) (models.Master, error) {
	master := models.Master{
		Id:       input.Id,
		UserID:   input.UserID,
		Name:     input.Name,
		Category: input.Category,
		Point:    input.Point,
	}

	_, err := models.DB.Query("UPDATE \"masters\" SET (userID, name, category, point) = ($1, $2, $3, $4) WHERE id = $5", master.UserID, master.Name, master.Category, master.Point, master_id)

	if err != nil {
		return master, err
	}
	return master, nil
}

func DeleteMaster(master_id int) (bool, error) {
	_, err := models.DB.Query("DELETE FROM \"masters\" WHERE id = $1", master_id)
	if err != nil {
		return false, err
	}
	return true, nil
}
