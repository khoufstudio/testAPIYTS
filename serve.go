package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/adapter/mysql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/labstack/echo"

	_httpdelivery "testAPIYTS/app/jadwalkajian/http"
	_repo "testAPIYTS/app/jadwalkajian/repository/mysql"
	_ucase "testAPIYTS/app/jadwalkajian/usecase"
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

	adapter, err := mysql.Open(dsn)
	if err != nil {
		panic(err)
	}
	defer adapter.Close()

	repo := rel.New(adapter)

	e := echo.New()

	timeoutContext := time.Duration(2) * time.Second
	jkrepo := _repo.NewMySQLJKRepo(repo)
	jkucase := _ucase.NewJKUsecase(jkrepo, timeoutContext)
	_httpdelivery.NewJKHandler(e, jkucase)

	log.Fatal(e.Start(":1323"))
}
