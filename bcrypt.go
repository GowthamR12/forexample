package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main(){
	s := 'password123'
	bs,err := bcrypt.GenerateFromPassword([]byte(s),bcrypt.MinCost)
	if err!=nil{
		fmt.Println("Error")
	}
	log := 'password123'
	err=bcrypt.CompareHashAndPassword(bs,[]byte(log))
	if err!=nil{
		fmt.Println("cant")
		return
	}
	fmt.Println("Login")
}