package main

import (
	"encoding/json"
	"fmt"
)

//JSON:convert từ Go object sang JSON strings
/*
type User struct {
	Id       string `json:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Social   Social `json:"social"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

*/

//Unmarshalling JSON ( convert tu JSON sang Go object)
/*
type Book struct {
	Title  string `json:"title"`
	Author string `json:"author"`
}

*/

func main() {
	/*social := Social{Facebook: "https://facebook.com", Twitter: "https://twitter.com"}
	user := User{Id: "ID001", Name: "LanKa", Password: "123465", Social: social}
	byteArray, err := json.MarshalIndent(user, "", "")
	if err != nil {
		fmt.Println(err)S
	}

	fmt.Println(string(byteArray))


	*/

	jsonString := `{"title":"Learning JSON in Golang","author":"LanKa"}`
	//var book Book
	//Unstructured data
	var book map[string]interface{}
	//--------------
	err := json.Unmarshal([]byte(jsonString), &book)

	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", book)

}
