package model

import (
	"log"
)

type DbTotos struct {
	ID          string
	Title       string
	Description string
	Completed   bool
	UserID      string
}

func (todo Todo) Create() string {
	db_todo := DbTotos{
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   todo.Completed,
		UserID:      todo.User.ID,
	}
	err := GetDB().Table("todos").Create(&db_todo).Error

	if err != nil {
		log.Fatal("Error: ", err)
	}

	log.Print("Row inserted!")
	return db_todo.ID
}

func GetTodoUserID(id string) string {
	todo := DbTotos{}
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
	var results []DbTotos
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

func UpdateTodo(input CompletedstatusInput) DbTotos {
	var result DbTotos
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
	result := DbTotos{ID: id}
	err := GetDB().Table("todos").Delete(&result).Error
	if err != nil {
		log.Fatal(err)
	}

	var todos []*Todo
	var results []DbTotos
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
