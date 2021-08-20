package main
import (
	"encoding/json"
	"fmt"
	"log"
)

type UserTwo struct {
	Id int
	Name string
	Occupation string
}
/*
var u1 = UserTwo{
	Id:         1,
	Name:       "Emmanuel",
	Occupation: "decadev",
}

var u2 = []UserTwo{
	{Id: 1,Name: "Emmanuel",Occupation: "student"},
	{Id: 2,Name: "tambarie",Occupation: "staff"},
	{Id: 3,Name: "clinton",Occupation: "senior dev"},
}
*/
//UNMARSHALLING
var uOne UserTwo
var uTwo []UserTwo


//MARSHAL INDENT



func main()  {

	//Marshalling
	//jsonData, err := json.Marshal(u1)
	//if err != nil{
	//	log.Fatal(err)
	//}
	//fmt.Println(string(jsonData))
	//
	//jsonDataTwo, err := json.Marshal(u2)
	//if err != nil{
	//	log.Fatal(err)
	//}

	//fmt.Println(string(jsonDataTwo))

	///////////////////
	//Unmarshalling
	data := []byte(`{"Id":1,
					"Name":"Emmanuel",
					"Occupation":"decadev"}`)
	errOne := json.Unmarshal(data,&uOne)
	if errOne != nil{
		log.Fatal(errOne)
	}
	//fmt.Printf("%+v",uOne)

	dataTwo := []byte(`
				[
				{"Id" : 2, "Name": "Roger Roe","Occupation": "driver"},
				{"Id": 3, "Name": "Lucy Smith","Occupation": "teacher"},
				{"Id" :4, "Name": "David Brown","Occupation": "Programmer"}
				]
			`)
	errTwo := json.Unmarshal(dataTwo,&uTwo)
	if errTwo != nil{
		log.Fatal(errTwo)
	}

	for _,v := range uTwo{
		fmt.Printf("%+v",v)
	}

	//Pretty print using Marshall Indent

	birds := map[string]interface{}{
		"sounds": map[string]string{
			"Pigeon": "coo",
			"eagle": "squak",
			"owl": "hoot",
			"duck": "quack",
			"cuckoo": "ku-ku",
			"raven": "cruck-cruck",
			"chicken": "cluck",
			"rooster": "cock-a-doodle-do",
		},
	}

	//dataThree,errThree := json.MarshalIndent(birds, "", "   ")
	dataThree,errThree := json.Marshal(birds)
	if errThree != nil{
		log.Fatal(errThree)
	}
	fmt.Println(string(dataThree))

}