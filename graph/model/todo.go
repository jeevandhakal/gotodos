package model

import (
	"fmt"
	"log"
)

func (todo Todo) Create() string {
	result := GetDB().Exec("Insert into todos(title,description,completed,user_id) values(?,?,?,?) returning id;", todo.Title, todo.Description, todo.Completed, todo.User.ID)

	if result.Error != nil {
		log.Fatal("Error:", result.Error)
	}
	var id string
	result.Scan(&id)
	fmt.Println(result)
	log.Print("Row inserted!")
	return id
}

func GetTodo(id uint) Todo {

	todo := Todo{}
	err := GetDB().Table("contacts").Where("id = ?", id).First(todo).Error
	if err != nil {
		log.Fatal(err)
	}
	return todo
}

func GetTodos(username string) []Todo {

	todos := make([]Todo, 0)
	user_id, _ := GetUserIdByUsername(username)
	err := GetDB().Table("todos").Where("user_id = ?", user_id).Find(&todos).Error
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(todos)
	for i, todo := range todos {
		todos[i].User = &User{
			ID:       user_id,
			Username: username,
		}

		fmt.Println(todo.User)
	}
	return todos
}
