package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First  string
	Last   string
	Age    int
	Gender string
}

func main() {
	p1 := person{
		First:  "Tako",
		Last:   "Yaki",
		Age:    5,
		Gender: "male",
	}

	p2 := person{
		First:  "Mimi",
		Last:   "Nini",
		Age:    16,
		Gender: "female",
	}

	people := []person{p1, p2}
	fmt.Println(people)

    // Marshal converts a Go struct into a JSON object.
	bs, err := json.Marshal(people)
	CheckError(err)
	fmt.Println(string(bs)) // now it's represented in JSON!

	s := string(bs)
	bs = []byte(s)
	fmt.Println(bs)
	people = []person{}
	//var people []person

    // The unmarshal function will take a JSON object in a Golang slice of
    // bytes and convert it into a Golang struct by filling in JSON fields
    // into a Golang struct's member variables. 
	err = json.Unmarshal(bs, &people)
	CheckError(err)

	for i, v := range people {
		fmt.Print("PERSON N ", i, " - ")
		fmt.Println(v.First, v.Last, v.Age, v.Gender)
	}
}
