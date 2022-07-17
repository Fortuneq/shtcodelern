package main

import "fmt"

type library []string

func (l library) getover(){
	for _,val:=range(l){
		fmt.Println(val)
	}
}
func main(){
	var lib library=library{"Mark","Reamk","Maria","Erich"}
	lib.getover()
}