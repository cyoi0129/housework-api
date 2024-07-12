package services

import (
	"fmt"
	"workout-note/models"
)

func FetchAllTaskList(user_id int) ([]models.Task, error) {
	var tasks []models.Task
	rows, err := models.DB.Query("SELECT id, userID, masterID, person, date FROM \"tasks\" WHERE userID = $1", user_id)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var task models.Task
		rows.Scan(&task.Id, &task.UserID, &task.MasterID, &task.Person, &task.Date)
		tasks = append(tasks, task)
	}
	fmt.Printf("%v", tasks)
	return tasks, nil
}

func FetchTaskList(user_id int, start string, end string) ([]models.Task, error) {
	var tasks []models.Task
	rows, err := models.DB.Query("SELECT id, userID, masterID, person, date FROM \"tasks\" WHERE userID = $1 AND date >= $2 AND date <= $3", user_id, start, end)
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		var task models.Task
		rows.Scan(&task.Id, &task.UserID, &task.MasterID, &task.Person, &task.Date)
		tasks = append(tasks, task)
	}
	fmt.Printf("%v", tasks)
	return tasks, nil
}

func CreateTask(input models.Task) (models.Task, error) {
	task := models.Task{
		UserID:   input.UserID,
		MasterID: input.MasterID,
		Person:   input.Person,
		Date:     input.Date,
	}
	err := models.DB.QueryRow("INSERT INTO tasks(userID, masterID, person, date) VALUES($1, $2, $3, $4) RETURNING id", task.UserID, task.MasterID, task.Person, task.Date).Scan(&task.Id)
	if err != nil {
		fmt.Println(err)
		return task, err
	}
	return task, nil
}

func UpdateTask(task_id int, input models.Task) (models.Task, error) {
	task := models.Task{
		Id:       input.Id,
		UserID:   input.UserID,
		MasterID: input.MasterID,
		Person:   input.Person,
		Date:     input.Date,
	}
	_, err := models.DB.Query("UPDATE \"tasks\" SET (userID, masterID, person, date) = ($1, $2, $3, $4) WHERE id = $5", task.UserID, task.MasterID, task.Person, task.Date, task_id)
	if err != nil {
		return task, err
	}
	return task, nil
}

func DeleteTask(task_id int) (bool, error) {
	_, err := models.DB.Query("DELETE FROM \"tasks\" WHERE id = $1", task_id)
	if err != nil {
		return false, err
	}
	return true, nil
}
