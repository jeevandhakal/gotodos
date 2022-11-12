package model

import (
	"log"
)

type Result struct {
	ID          string
	Title       string
	Description string
	Completed   bool
	UserID      string
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

func GetTodoUserID(id string) string {
	todo := Result{}
	err := GetDB().Table("todos").Where("id = ?", id).Find(&todo).Error
	if err != nil {
		log.Fatal(err)
	}

	return todo.UserID
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

func UpdateTodo(input CompletedstatusInput) Result {
	var result Result
	err := GetDB().Table("todos").Where("id = ?", input.ID).Find(&result).Error
	if err != nil {
		log.Fatal(err)
	}
	result.Completed = input.Completed
	err = GetDB().Table("todos").Save(&result).Error
	if err != nil {
		log.Fatal(err)
	}
	return result
}

func DeleteTodo(id string, username string) []*Todo {
	result := Result{ID: id}
	err := GetDB().Table("todos").Delete(&result).Error
	if err != nil {
		log.Fatal(err)
	}

	var todos []*Todo
	var results []Result
	user_id, _ := GetUserIdByUsername(username)
	err = GetDB().Table("todos").Where("user_id = ?", user_id).Find(&results).Error
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
