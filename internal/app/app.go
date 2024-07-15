package app

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"recruiter/internal/config"
	"recruiter/internal/handlers"
	"recruiter/internal/repository"
	"recruiter/internal/routes"

	httpSwagger "github.com/swaggo/http-swagger"
)

type App struct {
	Router     *mux.Router
	DB         *mongo.Database
	Config     *config.Config
	HTTPServer *http.Server
}

func New(cfg *config.Config) *App {
	app := &App{
		Router: mux.NewRouter(),
		Config: cfg,
	}

	app.initializeDB()
	app.initializeRoutes()

	return app
}

func (a *App) initializeDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI(a.Config.MongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	a.DB = client.Database(a.Config.DBName)
}

func (a *App) initializeRoutes() {
	repo := repository.NewMongoRepository(a.DB)

	routes.SetupApplicantRoutes(a.Router, handlers.NewApplicantHandler(repo))
	routes.SetupUserRoutes(a.Router, handlers.NewUserHandler(repo))

	a.Router.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler)
}

func (a *App) Run() {
	a.HTTPServer = &http.Server{
		Addr:    ":" + a.Config.Port,
		Handler: a.Router,
	}

	log.Printf("Server is running on port %s", a.Config.Port)
	if err := a.HTTPServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Could not listen on %s: %v\n", a.Config.Port, err)
	}
}

func (a *App) Shutdown(ctx context.Context) error {
	return a.HTTPServer.Shutdown(ctx)
}
