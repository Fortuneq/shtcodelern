package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "537j04222"
	dbname   = "postgres"
)

func main() {
	psqlconn := fmt.Sprintf("host = %s port = %d user = %s password = %s dbname = %s", host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlconn)
	tryerr(err)
	inserstJackhm := `insert into "clients"("id","first_name","last_name","role") values($1,$2,$3,$4)`
	_,g := db.Exec(inserstJackhm ,4,"Jack","London","Programmer")
	tryerr(g)
	defer db.Close()
}

func tryerr(err error) {
	if err != nil {
		panic(err)
	}
}
