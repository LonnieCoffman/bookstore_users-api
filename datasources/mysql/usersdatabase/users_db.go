package usersdatabase

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql" // mysql driver
	"github.com/joho/godotenv"
)

var (
	Client *sql.DB
)

func init() {
	godotenv.Load()
	fmt.Printf(os.Getenv("DB_USER"))
	dataSourceName := fmt.Sprintf(
		"%s:%s@tcp(%s)/%s?charset=utf8",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	var err error

	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}

	if err = Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("database successfully configured")
}
