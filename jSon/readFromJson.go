package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type UserFour struct {
	Name string
	Occupation string
	Born string
}

func main()  {
	filename,errFour := os.Open("data.json")
	if errFour != nil{
		log.Fatal(errFour)
	}
	defer filename.Close()

	data, errFive := ioutil.ReadAll(filename)
	if errFive != nil{
		log.Fatal(errFive)
	}

	var res []UserFour

	jsonError := json.Unmarshal(data, &res)
	if jsonError != nil{
		log.Fatal(jsonError)
	}

	fmt.Println(res)
}