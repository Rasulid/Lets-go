package main

import (
	"database/sql"
	"flag"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"snippetbox.rasulabduvaitov.net/internal/models"
)

type applicatiion struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	snippet  *models.SnippetModel
}

func main() {
	var addr string
	loginfo := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	logerr := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	flag.StringVar(&addr, "addr", ":4000", "IP address")

	flag.Parse()

	//database

	dsn := "host=localhost dbname=snippetbox sslmode=disable password=123rasulQq"
	db, er := sql.Open("postgres", dsn)
	if er != nil {
		log.Fatalf(er.Error())
	}
	defer db.Close()

	// Write logs on file
	f, r := os.OpenFile("/tmp/info.log", os.O_RDWR|os.O_CREATE, 0666)
	if r != nil {
		log.Fatal(r)
	}
	defer f.Close()

	// Config Logs

	app := &applicatiion{
		errorLog: logerr,
		infoLog:  loginfo,
		snippet:  &models.SnippetModel{DB: db},
	}
	//

	// ROUTER

	// SERVER AND LOGS

	// SERVER Configuration

	srv := &http.Server{
		Addr:     addr,
		ErrorLog: logerr,
		Handler:  app.routes(),
	}

	loginfo.Printf("Starting Server on %s", addr)
	err := srv.ListenAndServe()
	logerr.Fatal(err)
}
