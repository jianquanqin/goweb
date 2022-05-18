package models

import "gorm/program/dao"

//创建结构体，其实就相当于是创建数据表

type ToDo struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

func CreateToDo(todo *ToDo) (err error) {
	if err = dao.DB.Create(&todo).Error; err != nil {
		return err
	} else {
		return nil
	}
}

func GetAllToDo() (todoList []*ToDo, err error) {
	if err = dao.DB.Find(&todoList).Error; err != nil {
		return nil, err
	} else {
		return todoList, nil
	}
}

func GetToDo(id string) (todo *ToDo, err error) {
	todo = new(ToDo)
	if err = dao.DB.Where("id = ?", id).First(&todo).Error; err != nil {
		return nil, err
	} else {
		return todo, nil
	}
}

func UpdateTodo(todo *ToDo) (err error) {
	err = dao.DB.Save(&todo).Error
	return
}

func DeleteTodo(id string) (err error) {
	err = dao.DB.Where("id = ?", id).Delete(&ToDo{}).Error
	return
}
