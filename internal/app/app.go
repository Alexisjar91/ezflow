package app

import (
	"context"
	"database/sql"
	"log"
	"time"

	"ezflow/internal/config"
	"ezflow/internal/projects"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Run() {
	cfg := config.Load()

	connectDB(cfg.DatabaseURL)

	for {
		time.Sleep(time.Hour)
	}
}

func connectDB(url string) *sql.DB {
	db, err := sql.Open("pgx", url)
	if err != nil {
		log.Printf("error abriendo la bd: %v", err)
		return nil
	}

	for i := 0; i < 5; i++ {
		if err := db.Ping(); err == nil {
			log.Println("conectado a la base de datos")
			updateProject(db, "01KZTN4478YXVVGSJE8TY1W2R2", "proyecto actualizado", "descripcion actualizada")
			return db
		}
		time.Sleep(2 * time.Second)
	}

	log.Println("no se pudo conectar a la bd, la app sigue corriendo")
	return db
}

func updateProject(db *sql.DB, id, name, description string) {
	repo := projects.NewRepository(db)
	p, err := repo.Update(context.Background(), id, name, description)
	if err != nil {
		log.Printf("error actualizando proyecto: %v", err)
		return
	}
	log.Printf("proyecto actualizado: %v", p)
}
