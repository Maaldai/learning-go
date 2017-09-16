package main

import "fmt"

type Person struct {
	firstName string
	lastName string
}
// Changing the Value to Another Name : Raj //
func changeName (p *Person){
	p.firstName ="Raj"
}

func main(){
	person := Person{
		firstName:"Maal",
		lastName : "Dai",
	}

	changeName(&person)
	fmt.Println(person)


}

// After you run the code, you shall get the Following output : {Raj Dai}
// In this Case, Variable person in function main is modified inside function
// changeName. This happens because &person and p are two different pointers to the same struct
//which is stored at the same memory location.