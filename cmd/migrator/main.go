package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/pressly/goose/v3"
	"github.com/pujidjayanto/goginboilerplate/pkg/log"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	log.Init()
	defer log.SyncLogger()

	var (
		doMigration   bool
		doSeed        bool
		migrationName string
		seedName      string
	)

	flag.BoolVar(&doMigration, "migrate", false, "Run database migrations")
	flag.BoolVar(&doSeed, "seed", false, "Run database seeds")
	flag.StringVar(&migrationName, "create-migration", "", "Create a new migration file")
	flag.StringVar(&seedName, "create-seed", "", "Create a new seed file")

	// use flag.Usage to add help description on how to run the program
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		fmt.Println("Options:")
		flag.PrintDefaults()
		fmt.Println("\nExamples:")
		fmt.Println("  Create a new migration:")
		fmt.Println("    go run ./cmd/migrator -create-migration init_tables")
		fmt.Println("  Create a new seed:")
		fmt.Println("    go run ./cmd/migrator -create-seed seed_init")
		fmt.Println("  Run migrations up:")
		fmt.Println("    go run ./cmd/migrator -migrate up")
		fmt.Println("  Run migrations down:")
		fmt.Println("    go run ./cmd/migrator -migrate down")
		fmt.Println("  Run seeds:")
		fmt.Println("    go run ./cmd/migrator -seed")
	}

	flag.Parse()

	if migrationName != "" {
		createNewMigration(migrationName)
		return
	}

	if seedName != "" {
		createNewSeed(seedName)
		return
	}

	if doMigration {
		args := flag.Args()
		if len(args) < 1 {
			log.Fatal("Please specify 'up' or 'down' for migration")
		}
		migrationDirection := args[0]

		env, err := loadConfiguration()
		if err != nil {
			log.Fatal(err.Error())
		}

		dsn := env.DatabaseDSN()
		db, err := prepareDatabase(dsn)
		if err != nil {
			log.Fatal(err.Error())
		}

		log.Info("migration running", zap.String("env", env.Migrator.Env))
		runMigration(db, migrationDirection)

		if env.Migrator.Env == "development" {
			log.Info("migration running", zap.String("env", "test"))
			dsnTestDatabase := env.TestDatabaseDSN()
			dbTest, err := prepareDatabase(dsnTestDatabase)
			if err != nil {
				log.Fatal(err.Error())
			}
			runMigration(dbTest, migrationDirection)
		}
		return
	}

	if doSeed {
		env, err := loadConfiguration()
		if err != nil {
			log.Fatal(err.Error())
		}

		dsn := env.DatabaseDSN()
		log.Info(dsn)
		db, err := prepareDatabase(dsn)
		if err != nil {
			log.Fatal(err.Error())
		}

		runSeeds(db)
	}
}

func prepareDatabase(dsn string) (*sql.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NowFunc: func() time.Time { return time.Now().UTC() },
	})
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %w", err)
	}

	sqlDb, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("error opening database pool: %w", err)
	}

	err = sqlDb.Ping()
	if err != nil {
		return nil, fmt.Errorf("can't ping the db: %v", err)
	}

	return sqlDb, nil
}

func runMigration(db *sql.DB, direction string) {
	migrationDir := filepath.Join("migrations")
	goose.SetTableName("migration_history")

	switch direction {
	case "up":
		if err := goose.Up(db, migrationDir); err != nil {
			log.Fatal("failed to run migrations", zap.String("error", err.Error()))
		}
		log.Info("Migrations applied successfully.")
	case "down":
		if err := goose.Down(db, migrationDir); err != nil {
			log.Fatal("failed to run migrations")
		}
		log.Info("Migrations rolled back successfully.")
	default:
		log.Fatal("Invalid migration direction. Use 'up' or 'down'.")
	}
}

func runSeeds(db *sql.DB) {
	seedDir := filepath.Join("migrations", "seeds")
	goose.SetTableName("seed_history")
	if err := goose.Up(db, seedDir); err != nil {
		log.Fatal("failed to run seeds", zap.String("error", err.Error()))
	}

	log.Info("seeds applied successfully.")
}

func createNewMigration(name string) {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to get working directory: %s", err))
	}

	migrationDir := filepath.Join(workingDir, "migrations")

	if _, err := os.Stat(migrationDir); os.IsNotExist(err) {
		if err := os.MkdirAll(migrationDir, os.ModePerm); err != nil {
			log.Fatal(fmt.Sprintf("failed to create migrations directory: %s", err))
		}
	}

	if err := goose.Create(nil, migrationDir, name, "sql"); err != nil {
		log.Fatal(fmt.Sprintf("failed to create migration: %s", err))
	}
	log.Info(fmt.Sprintf("Migration %s created successfully in %s.\n", name, migrationDir))
}

func createNewSeed(name string) {
	workingDir, err := os.Getwd()
	if err != nil {
		log.Fatal(fmt.Sprintf("failed to get working directory: %s", err))
	}

	seedDir := filepath.Join(workingDir, "migrations", "seeds")

	if _, err := os.Stat(seedDir); os.IsNotExist(err) {
		if err := os.MkdirAll(seedDir, os.ModePerm); err != nil {
			log.Fatal(fmt.Sprintf("failed to create seeds directory: %s", err))
		}
	}

	if err := goose.Create(nil, seedDir, name, "sql"); err != nil {
		log.Fatal(fmt.Sprintf("failed to create seed: %s", err))
	}
	log.Info(fmt.Sprintf("Seed %s created successfully in %s.\n", name, seedDir))
}
