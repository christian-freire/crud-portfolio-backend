package main

import (
	"projetos-pessoais/crud-porfolio-backend/database"
	"projetos-pessoais/crud-porfolio-backend/routes"
)

func main() {
	database.ConectToDataBase()
	routes.HandleRequests()
}
