package server

import "github.com/naughtydevelopment/todo-go/config"

func Init() {
	c := config.GetConfig()
	port := c.GetString("server.port")
	r := NewRouter()
	r.Run(port)
}
