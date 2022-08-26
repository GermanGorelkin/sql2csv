package main

import (
	"context"
	"database/sql"
	"io/ioutil"
	"log"
	"os"

	_ "github.com/denisenkom/go-mssqldb"

	"github.com/germangorelkin/sql2csv"
)

func main() {
	queryFile := os.Getenv("QUERY_FILE")
	if queryFile == "" {
		log.Fatal("QUERY_FILE is not set.")
	}
	outFile := os.Getenv("OUT_FILE")
	if outFile == "" {
		log.Fatal("OUT_FILE is not set.")
	}
	connDB := os.Getenv("DATABASE_URL")
	if connDB == "" {
		log.Fatal("DATABASE_URL is not set.")
	}

	b, err := ioutil.ReadFile(queryFile)
	if err != nil {
		log.Fatal(err)
	}
	query := string(b)

	db, err := sql.Open("sqlserver", connDB)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rd := sql2csv.SQLReader{DB: db}

	fd, err := os.Create(outFile)
	if err != nil {
		log.Fatal(err)
	}
	csvWriter := sql2csv.NewCSVWriter([]byte("\t"), []byte("\r\n"), fd)

	ctx, stop := context.WithCancel(context.Background())
	defer stop()

	err = rd.Read(ctx, query, csvWriter)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("export completed")
}
