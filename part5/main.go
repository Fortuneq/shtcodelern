package main

import "fmt"

const (
	host = "localhost"
	port = 5432
	user = "postgres"
	//password = "537j04222"
	dbname = "postgres"
)
func main(){
	fmt.Println(host,port,user,dbname)
}