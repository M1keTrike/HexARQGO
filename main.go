package main

import (
	"demo/dependencies"
)

func main() {
	deps := dependencies.NewDependencies()
	r, db := deps.Execute()

	
	defer db.Close()

	r.Run(":8080")
}
