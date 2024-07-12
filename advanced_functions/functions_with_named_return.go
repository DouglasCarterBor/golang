package main

func calculate (n1, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}

func main () {
   sum, sub := calculate(10, 5)
   println(sum, sub)

}