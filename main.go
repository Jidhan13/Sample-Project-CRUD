package main

import (
	"crud/routes"

	"crud/db"
)

func main() {

	db.Init()

	e := routes.Init()

	e.Logger.Fatal(e.Start(":8000"))
}
