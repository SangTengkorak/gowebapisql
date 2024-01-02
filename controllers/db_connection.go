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
	//Load .env file using godotenv module, and use it as varibles value
	err := godotenv.Load("./local.env")
	if err != nil {
		fmt.Println("Error loading.env file")
	}

	dbuser := os.Getenv("DBUSER")
	dbname := os.Getenv("DBNAME")
	dbhost := os.Getenv("DBHOST")
	dbpass := os.Getenv("DBPASS")

	//Initate database connecton and store the session
	db, err := sql.Open("mysql", dbuser+":"+dbpass+"@tcp("+dbhost+":3306)/"+dbname)
	if err != nil {
		fmt.Println("Err", err.Error())
	}

	if err != nil {
		panic("Failed to connect to database")
	}

	DB = db

}
