package cmd

import (
	"server/config"
	"server/infra/db"
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
		panic(err)
	}
	err = db.MigrateDB(dbCon, "./migrations")
	if err != nil {
		panic(err)
	}

	productRepo := repo.NewProductRepo(dbCon)
	middlewares := middleware.NewMiddleware(cnf)
	productHandler := product.NewProductHandler(
		middlewares,
		productRepo,
	)
	rest := rest.NewServer(
		cnf,
		productHandler,
	)

	rest.Start()
}
