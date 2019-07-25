package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := `password123`
	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	CheckError(err)
	fmt.Println(s)
	fmt.Println(string(bs))

	loginPswd := `password123`

	err = bcrypt.CompareHashAndPassword(bs, []byte(loginPswd))
	CheckError(err)

	fmt.Println("OK")
}
