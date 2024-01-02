package controllers

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"

	"fmt"
)

var DB *sql.DB

func ConnectDatabase() {
	err := godotenv.Load("./local.env")
	if err != nil {
		fmt.Println("Error loading.env file")
	}

	dbuser := os.Getenv("DBUSER")
	dbname := os.Getenv("DBNAME")
	dbhost := os.Getenv("DBHOST")
	dbpass := os.Getenv("DBPASS")

	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp("+dbhost+":3306)/"+dbname)
	if err != nil {
		fmt.Println("Err", err.Error())
	}

	if err != nil {
		panic("Failed to connect to database")
	}

	DB = db

}
