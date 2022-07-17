package main
import "fmt"
 
type contact struct{
    email string
    phone string
}
 
type person struct{
    name string
    age int
    contactInfo contact
}
 
func main(){
	var tom person
	tom.age = 18
	tom.name = "TOM"
	tom.contactInfo.email = "vladvolkov03@mail.ru"
	tom.contactInfo.phone = "89821135322"
	fmt.Println(tom)
}