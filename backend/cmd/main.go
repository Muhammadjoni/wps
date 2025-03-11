package main

import (
	"backend/config"
	"backend/internal/controller"
	"backend/internal/repository"
	"backend/internal/routes"
	"backend/internal/service"
	"net/http"

	"go.uber.org/zap"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// LoadConfig("config")
	cfg, err := config.LoadConfig(".")
	if err != nil {
		panic(err)
	}

	//init logger
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	//db connection

	// pgPool, err := pgxpool.New(context.Background(),
	// 	"postgres://"+cfg.DB.User+":"+cfg.DB.Password+"@"+cfg.DB.Host+":"+cfg.DB.Port+"/"+cfg.DB.DBName)
	// if err != nil {
	// 	logger.Fatal("Failed to connect to PostgreSQL", zap.Error(err))
	// }
	// defer pgPool.Close()

	//redis connection
	/////// redisCache := redis.NewCache(&cfg.Redis)

	//init repository

	//create router

	//middleware

	// routes

	// start server

	// Initialize dependencies
	cfg := config.Load()
	db := database.ConnectPostgres(cfg)
	redis := cache.NewRedis(cfg)

	// Initialize layers
	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo, redis)
	productHandler := controller.NewProductHandler(productService)

	// Create router
	r := chi.NewRouter()

	// Global middleware
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	// Register routes
	routes.ProductRoutes(r, productHandler, middleware.TokenAuth)

	// Start server
	http.ListenAndServe(":8080", r)

}
