package main

import "fmt"

type polygon struct{
	name string  
	sides int
}

func main(){
	sq := polygon{"Square",4}
	fmt.Println(sq.name, sq.sides)

}
