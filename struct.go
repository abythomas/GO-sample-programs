package main

import "fmt"
import "encoding/json"

type address struct{
	House_no int
	Street string
	Place string
}

type person struct{
	Name string  
	Age int
	address
}

func main(){
	pr := &person{
		Name : "Jason Bourne",
		Age : 42,
		address: address{
		House_no : 151,
		Street : "Holly Lane",
		Place : "Virginia",
		},
    }
    fmt.Println(pr)
    res, _ := json.Marshal(pr)
    fmt.Println(string(res))

}
