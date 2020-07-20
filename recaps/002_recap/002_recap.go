package main

import (
	"fmt"
)

// create a struct that holds person fields //
// create a struct that holds secret agent fields and embeds person type //
// attach a method to person: pSpeak //
// attach a method to secret agent: saSpeak //
// create a variable of type person //
// create a variable of type secret agent //
// print a field from person //
// run pSpeak attached to the variable of type person //
// print a field from secret agent //
// run saSpeak attached to the variable of type secret agent //
// run pSpeak attached to the variable of type secret agent //

type person struct {
	fname string
	lname string
}

type secretAgent struct{
	person
	weapon string
	vehicle string
}
func (p person)pSpeak(){
	fmt.Println("Hello,", p.fname, p.lname)

}
func (sa secretAgent)saSpeak(){
	fmt.Printf("Welcome, %v, your %v is waiting with your %v in the glove compartment.", sa.person,sa.vehicle, sa.weapon)

}

func main(){
	p := person{
		"Bootsy",
		"Kitty",
	}
	s := secretAgent {
		person {
			"Beta",
			"Pupper",
		},
		"Bone",
		"Subie",
	}
	fmt.Println(p.fname)
	fmt.Println(s.vehicle)
	p.pSpeak()
	s.saSpeak()



}



