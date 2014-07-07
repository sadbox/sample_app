package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	db     *sql.DB
	config Config
)

type Config struct {
	DBUsername string
	DBPassword string
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, world!<br>Just testing!")
}

func init() {
	configfile, err := os.Open("/etc/sample_app_config.json")
	if err != nil {
		log.Fatal(err)
	}
	err = json.NewDecoder(configfile).Decode(&config)
	if err != nil {
		log.Fatal(err)
	}
	mysqlLogin := fmt.Sprintf("%s:%s@db.civis.sadbox.org/greeting", config.DBUsername, config.DBPassword)
	db, err = sql.Open("mysql", mysqlLogin)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
