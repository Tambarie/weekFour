package main
import (
	"encoding/json"
	"fmt"
	"log"
)

type User struct {
	Id	int
	Name	string
	Occupation string
}




func main()  {
	u1 := User{
		Id:         1,
		Name:       "John Doe",
		Occupation: "student",
	}

	json_data,err :=json.Marshal(u1)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(string(json_data))

	users := []User{
		{Id: 2,Name: "Roger Roe",Occupation: "driver"},
		{Id: 3, Name: "Lucy Smith",Occupation: "teacher"},
		{Id: 4,Name: "David Brown",Occupation: "programmer"},
	}

	jsonDataTwo, err := json.Marshal(users)
	if err != nil{
		log.Fatal(err)
	}

	fmt.Println(string(jsonDataTwo))
}