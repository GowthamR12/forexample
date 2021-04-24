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
func (s secretAgent) speak(){
	fmt.Println("I'm",s.first,s.last," secret agent speaks")
}
func (s person) speak(){
	fmt.Println("I'm",s.first,s.last,"Person speaks")
}
type human interface{
	speak()
}
func bar(h human){
	switch h.(type){
	case person: fmt.Println("Hello person",h.(person).first)
	case secretAgent:fmt.Println("Hello SA",h.(secretAgent).first)
	}
	fmt.Println("hello",h)
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

	sa.speak()
	p1.speak()

	bar(sa)
	bar(p1)
}
