package model

import (
	"log"
)

type Result struct {
	ID          string
	Title       string
	Description string
	Completed   bool
}

func (todo Todo) Create() string {
	err := GetDB().Exec("Insert into todos(title,description,completed,user_id) values(?,?,?,?) returning id;", todo.Title, todo.Description, todo.Completed, todo.User.ID).Error

	if err != nil {
		log.Fatal("Error: ", err)
	}
	log.Print("Row inserted!")

	var toDo Todo
	err = GetDB().Table("todos").Last(&toDo).Error
	if err != nil {
		log.Fatal("Error: ", err)
	}
	return toDo.ID
}

func GetTodo(id uint) Todo {

	todo := Todo{}
	err := GetDB().Table("contacts").Where("id = ?", id).First(todo).Error
	if err != nil {
		log.Fatal(err)
	}
	return todo
}

func GetTodos(username string) []*Todo {

	todos := make([]*Todo, 0)
	user_id, _ := GetUserIdByUsername(username)
	log.Printf("Getting todos of user of id:%v and username:%v", user_id, username)
	var results []Result
	err := GetDB().Table("todos").Where("user_id = ?", user_id).Find(&results).Error
	if err != nil {
		log.Fatal(err)
	}

	for _, toDo := range results {
		todo := &Todo{
			ID:          toDo.ID,
			Title:       toDo.Title,
			Description: toDo.Description,
			Completed:   toDo.Completed,
			User: &User{
				ID:       user_id,
				Username: username,
			},
		}

		todos = append(todos, todo)
	}
	return todos
}
