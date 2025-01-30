package main

import (
	"demo/src/database"
	orderDeps "demo/src/orders/dependencies"
	productDeps "demo/src/products/dependencies"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal("Could not establish database connection:", err)
	}
	defer db.Close()

	r := gin.Default()

	
	productsDeps := productDeps.NewProductsDependencies(db)
	productsDeps.Execute(r)
	
	ordersDeps := orderDeps.NewOrdersDependencies(db)
	ordersDeps.Execute(r)
	r.Run(":8080")
}
