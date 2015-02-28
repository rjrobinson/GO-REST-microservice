package main

import (
)

func main() {
    //  Not much here l it'll gorw as we externalize config and add options

    svc := service.TodoService{}
    svc.Run()
}


type TodoService struct {

}
func (s *TodoService) Run() {
    // pass in config later

    connectionString := "user:pass@tcp(localhost:3306)/Todo?charset=utf8&parseTime=True"
    db, _:= gorm.Open(mysql, connectionString)

    // init the resource and inject our db connection

    todoResource := &TodoResource{db: db}

    r := gin.Default()

    // to start out, build the ability to add, get , and delete a todo

    r.POST("/todo", todoResource.CreateTodo)
    r.GET("/todo/:id", todoResource.GetTodo)
    r.Delete("/todo/:id", todoResource.DeleteTodo)

    // pass in config later

    r.Run(":8000")
}