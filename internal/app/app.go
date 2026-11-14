package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/eyobmdev/femProject/internal/api"
	"github.com/eyobmdev/femProject/internal/store"
	"github.com/eyobmdev/femProject/migrations"
)

type Application struct {
	Logger         *log.Logger
	WorkoutHandler *api.WorkoutHandler
	DB             *sql.DB
}

func NewApplication() (*Application, error) {
	pgDB, err := store.Open()
	if err != nil {
		return nil, fmt.Errorf("app: open database: %w", err)
	}

	err = store.MigrateFS(pgDB, migrations.FS, ".")
	if err != nil {
		pgDB.Close()
		return nil, fmt.Errorf("app: migrate database: %w", err)
	}

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lshortfile)

	workoutStore := store.NewPostgresWorkoutStore(pgDB)
	workoutHandler := api.NewWorkoutHandler(workoutStore, logger)

	app := &Application{
		Logger:         logger,
		WorkoutHandler: workoutHandler,
		DB:             pgDB,
	}

	return app, nil
}

func (app *Application) HealthChecker(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available\n")

}
