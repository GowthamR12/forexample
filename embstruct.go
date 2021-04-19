package main

import "fmt"

type person struct {
	first string
	last  string
	age   int
}
type secretAgent struct {
	person
	ltk bool
}

func main() {
	sa := secretAgent{
		person: person{
			first: "James",
			last:  "Bond",
			age:   32,
		},
		ltk: true,
	}
	p1 := person{
		first: "miss",
		last:  "money",
		age:   27,
	}
	fmt.Println(sa.person, sa.ltk)
	fmt.Println(p1.first, p1.last, p1.age)
}
