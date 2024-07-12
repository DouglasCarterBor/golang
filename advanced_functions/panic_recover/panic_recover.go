package main

import "fmt"

func recoveryExec() {
	if r := recover(); r != nil { 
	fmt.Println("Trying to recover")
	}
}

func studentAproval(n1, n2 float32) bool {
	defer recoveryExec()
	
	average := (n1 + n2) / 2 

	if average > 6 {
		return true
	} else if average < 6 {
        return false
	}
	
	panic("The average is exactly 6")

}

func main() {
	fmt.Println(studentAproval(6, 6))
	fmt.Println("End of the program")

}
