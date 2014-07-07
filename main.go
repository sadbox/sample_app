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
	rows, err := db.Query("SELECT * FROM greetings")
	if err != nil {
		log.Println(err)
		return
	}
	defer rows.Close()
	for rows.Next() {
		var greeting string
		var language string
		err = rows.Scan(&greeting, &language)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Fprintf(w, "%s, world! (%s)\n", greeting, language)
	}
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
	mysqlLogin := fmt.Sprintf("%s:%s@tcp(db.civis.sadbox.org:3306)/test_application", config.DBUsername, config.DBPassword)
	log.Println(mysqlLogin)
	db, err = sql.Open("mysql", mysqlLogin)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
