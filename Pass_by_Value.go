package main

import "fmt"

type Person struct {
	firstName string
	lastName string
}
// Changing the Value to Another Name : Raj //
 func changeName (p Person){
 	p.firstName ="Raj"
 }

func main(){
	person := Person{
		firstName:"Maal",
		lastName : "Dai",
	}

	changeName(person)
	fmt.Println(person)


}

// After you run the code, you shall get the Following output : {Maal Dai}
// Note that even though function changeName changes firstName to "Raj",
// the change doesn't affect variable person in function main. This happens because,
// function changeName modifies a copy of variable person, not person itself.