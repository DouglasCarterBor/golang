package main

import "fmt"

func toWrite() {
	fmt.Println("To write...")
}

type user struct {
	name string
	age uint8
}

func (u user) save() {
	fmt.Printf("Saving user %v with age %v\n", u.name, u.age)
}

func (u user) ofLegalAge() bool {
	return u.age >= 18
}

func haveABirthday(u *user) {
	u.age++
}

func main() {
	toWrite()
	user1 := user{"John", 25}
	fmt.Println(user1)
	user1.save()
	
	user2 := user{"Jane", 30}
	fmt.Println(user2)
	user2.save()

	user3 := user{"Jack", 15}
	ofLegalAge := user3.ofLegalAge()
	fmt.Printf("Is %v of legal age? %v\n", user3.name, ofLegalAge)

	haveABirthday(&user3)
	fmt.Printf("Happy birthday %v! You are now %v years old\n", user3.name, user3.age)

}