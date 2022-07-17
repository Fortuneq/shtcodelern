package main

import (
	"encoding/json"
	"fmt"
)

func main(){
mygreatJson := `{"birds":{"pigeon":"nice but not fly","eagle":700},"animal":{"lion":"roaring","gorilla":"glass"}}`

var result map[string]interface{}

json.Unmarshal([]byte(mygreatJson),&result)

for key,val := range result{
	fmt.Println(key, val)
}
}