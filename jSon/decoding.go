package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type UserTwo struct {
	Id	int
	Name	string
	Occupation	string
}

func main()  {
	var ul UserTwo
	data := []byte(`{"Id":1,"Name":"John Doe","Occupation":"student"}`)
	err := json.Unmarshal(data, &ul)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println("Struct is:",ul)
	fmt.Printf("%s is a %s. \n", ul.Name, ul.Occupation)


}


//{"Id":1,"Name":"John Doe","Occupation":"student"}
//[{"Id":2,"Name":"Roger Roe","Occupation":"driver"},{"Id":3,"Name":"Lucy Smith","Occupation":"teacher"},{"Id":4,"Name":"David Brown","Occupation":"programmer"}]

