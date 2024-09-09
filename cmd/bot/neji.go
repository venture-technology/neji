package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/venture-technology/neji/config"
	"github.com/venture-technology/neji/internal/bot"
	"github.com/venture-technology/neji/internal/repository"
	"github.com/venture-technology/neji/internal/usecase"

	_ "github.com/lib/pq"
)

func main() {

	config, err := config.Load("../../config/config.yaml")
	if err != nil {
		log.Fatalf("failed to load config: %v", err)
	}

	db, err := sql.Open("postgres", postgres(config.Database))
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	err = migrate(db, config.Database.Schema)
	if err != nil {
		log.Fatalf("failed to execute migrations: %v", err)
	}

	nejiRepository := repository.NewNejiRepository(db)
	uc := usecase.NewNejiUseCase(nejiRepository)

	bot := bot.NewDiscordBot(config.Discord.Token, uc)

	bot.Setup()

}

func postgres(dbconfig config.Database) string {
	return "user=" + dbconfig.User +
		" password=" + dbconfig.Password +
		" dbname=" + dbconfig.Name +
		" host=" + dbconfig.Host +
		" port=" + dbconfig.Port +
		" sslmode=disable"
}

func migrate(db *sql.DB, filepath string) error {
	schema, err := os.ReadFile(filepath)
	if err != nil {
		return err
	}

	_, err = db.Exec(string(schema))
	if err != nil {
		return err
	}

	return nil
}
