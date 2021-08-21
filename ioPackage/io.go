package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func main()  {
	//Copying data
	r := strings.NewReader("Hello world\n")

	_, err := io.Copy(os.Stdout, r)
	CheckError(err)
	//fmt.Println(r)

	// Reading data using the io package
	s := strings.NewReader("Read this string\n")

	buf := make([]byte,20)
	_, err1 := io.ReadAtLeast(s,buf,4)
	CheckError(err1)
	fmt.Println(string(buf))

	bufTwo := make([]byte, 2)
	_,errTwo := io.ReadAtLeast(s,bufTwo,4)
	CheckError(errTwo)
	fmt.Println(bufTwo)
	fmt.Println(string(bufTwo))

	bufThree := make([]byte,20)
	_, errThree := io.ReadAtLeast(s,bufThree,20)
	fmt.Println(errThree)
}

func CheckError(e error)  {
	if e != nil{
		fmt.Println(e)
	}
}

