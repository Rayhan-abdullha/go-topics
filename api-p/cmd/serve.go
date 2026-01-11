package cmd

import (
	"fmt"
	"server/config"
	"server/infra/db"
	productSvc "server/product"
	"server/repo"
	"server/rest"
	"server/rest/handler/product"
	"server/rest/middleware"
)

func Serve() {
	cnf := config.LoadConfig()

	// database connection setup can be added here if needed
	dbCon, err := db.NewConnection()
	if err != nil {
		fmt.Println("Error connecting to the database:")
		panic(err)
	}
	err = db.MigrateDB(dbCon, "./migrations")
	if err != nil {
		panic(err)
	}

	productRepo := repo.NewProductRepo(dbCon)
	middlewares := middleware.NewMiddleware(cnf)

	productSvc := productSvc.NewService(productRepo)

	productHandler := product.NewProductHandler(
		middlewares,
		productSvc,
	)
	rest := rest.NewServer(
		cnf,
		productHandler,
	)

	rest.Start()
}
