package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/adapter/mysql"
	"github.com/go-rel/rel/migrator"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"go.uber.org/zap"

	"testAPIYTS/db/migrations"
)

var (
	logger, _ = zap.NewProduction(zap.Fields(zap.String("type", "main")))
	shutdowns []func() error
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	username := os.Getenv("dbuname")
	dbname := os.Getenv("dbname")
	dbhost := os.Getenv("dbhost")
	dbport := os.Getenv("dbport")

	dsn := fmt.Sprintf("%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, dbhost, dbport, dbname)

	fmt.Println(dsn)

	adapter, err := mysql.Open(dsn)
	if err != nil {
		logger.Fatal(err.Error(), zap.Error(err))
	}

	var op string
	repo := rel.New(adapter)
	m := migrator.New(repo)
	ctx := context.Background()

	m.Register(20202806225100, migrations.MigrateCreateJadwalKajian, migrations.RollbackCreateJadwalKajian)

	if len(os.Args) > 1 {
		op = os.Args[1]
	}

	switch op {
	case "migrate", "up":
		m.Migrate(ctx)
	case "rollback", "down":
		m.Rollback(ctx)
	default:
		logger.Fatal("command not recognized", zap.String("command", op))
	}

}
