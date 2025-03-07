package main

import (
	"database/sql"
	"flag"
	"fmt"
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
		doMigration bool
		doSeed      bool
		doUp        bool
		doDown      bool
		doCount     int
	)

	flag.BoolVar(&doMigration, "migrate", false, "Run database migrations")
	flag.BoolVar(&doSeed, "seed", false, "Run database seeds")
	flag.BoolVar(&doUp, "up", false, "Do up migration")
	flag.BoolVar(&doDown, "down", false, "Do down migration by 1 version (no need count params)")
	flag.IntVar(&doCount, "count", 0, "Do N up migration, default to do up all migration")

	flag.Parse()

	env, err := loadEnvironment()
	if err != nil {
		log.Fatal(err.Error())
	}

	dsn := env.DatabaseDSN()
	db, err := prepareDatabase(dsn)
	if err != nil {
		log.Fatal(err.Error())
	}

	log.Info("migration running", zap.String("env", env.Migrator.Env))
	runMigrator(db, doMigration, doSeed, doUp, doDown, doCount)

	if env.Migrator.Env == "development" {
		log.Info("migration running", zap.String("env", "test"))
		dsnTestDatabase := env.TestDatabaseDSN()
		dbTest, err := prepareDatabase(dsnTestDatabase)
		if err != nil {
			log.Fatal(err.Error())
		}
		runMigrator(dbTest, doMigration, doSeed, doUp, doDown, doCount)
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

func runMigrator(
	db *sql.DB,
	doMigration, doSeed, doUp, doDown bool,
	doCount int,
) {
	if doSeed {
		runSeeds(db)
		return
	}

	if doMigration && doUp && doDown {
		log.Fatal("what should i do, up or down?, choose one!")
	}

	if doMigration && doUp {
		if doCount == 0 {
			upAllMigration(db)
		} else {
			for i := 0; i < doCount; i++ {
				upByOneMigration(db)
			}
		}
	}

	if doMigration && doDown {
		downMigration(db)
	}
}

func upAllMigration(db *sql.DB) {
	migrationDir := filepath.Join("..", "..", "migrations")
	goose.SetTableName("migration_history")
	if err := goose.Up(db, migrationDir); err != nil {
		log.Fatal("failed to run migrations")
	}
	fmt.Println("Migrations applied successfully.")
}

func upByOneMigration(db *sql.DB) {
	migrationDir := filepath.Join("..", "..", "migrations")
	goose.SetTableName("migration_history")
	if err := goose.UpByOne(db, migrationDir); err != nil {
		log.Fatal("failed to run one migration")
	}
	fmt.Println("Migrations applied successfully.")
}

func downMigration(db *sql.DB) {
	migrationDir := filepath.Join("..", "..", "migrations")
	goose.SetTableName("migration_history")
	if err := goose.Down(db, migrationDir); err != nil {
		log.Fatal("failed to run migrations")
	}
	fmt.Println("Migrations applied successfully.")
}

func runSeeds(db *sql.DB) {
	seedDir := filepath.Join("..", "..", "seeds")
	goose.SetTableName("seed_history")
	if err := goose.Up(db, seedDir); err != nil {
		log.Fatal("failed to run seeds")
	}

	log.Info("seeds applied successfully.")
}
