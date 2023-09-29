package main

import (
	"fmt"

	"github.com/devaneando/tags/pkg/model"
)

func main() {
	// p := model.Person{
	// 	Name: "Señor Lolladas",
	// }

	// err := p.Save("myfile.json")
	// if err != nil {
	// 	panic(err)
	// }

	var p2 model.Person
	p2.Load("myfile.json")
	fmt.Printf("at main decoded %v\n", p2)

	// fmt.Printf("%v\n", p)
}
