package main

import (
	"go_api/db"
	"go_api/server"
)

func main() {
	db.Init()
	server.Init()
	db.Close()
}