package app

import (
	"context"
	"log"
	"time"

	"ezflow/internal/config"
	"ezflow/internal/projects"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

func Run() {
	cfg := config.Load()

	connectDB(cfg.DatabaseURL)

	for {
		time.Sleep(time.Hour)
	}
}

func connectDB(url string) *sqlx.DB {
	db, err := sqlx.Open("pgx", url)
	if err != nil {
		log.Printf("error abriendo la bd: %v", err)
		return nil
	}

	for i := 0; i < 5; i++ {
		if err := db.Ping(); err == nil {
			log.Println("conectado a la base de datos")
			seveProject(db)
			return db
		}
		time.Sleep(2 * time.Second)
	}

	log.Println("no se pudo conectar a la bd, la app sigue corriendo")
	return db
}

func seveProject(db *sqlx.DB) projects.Project {
	p := &projects.Project{
		Name:        "prueba service 2",
		Description: "descripcion de prueba service",
		Color:       "#FF0000",
	}
	repo := projects.NewProjectRepository(db)
	service := projects.NewProjectService(*repo)
	createdProject, err := service.CreateProject(context.Background(), p)
	if err != nil {
		log.Printf("error guardando columna: %v", err)
		return projects.Project{}
	}
	return createdProject
}
