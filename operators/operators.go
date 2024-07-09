package main

import "fmt"

func main() {

   // atithmetic operators
   sum := 1 + 2
   subtraction := 1 - 2
   multiplication := 1 * 2
   division := 1 / 2
   rest := 1 % 2
   
   fmt.Println("Sum:", sum)
   fmt.Println("Subtraction:", subtraction)
   fmt.Println("Multiplication:", multiplication)
   fmt.Println("Division:", division)
   fmt.Println("Rest:", rest)

   var num1 int = 10
   var num2 int = 20
   sum2 := num1 + num2
   fmt.Println("Sum:", sum2)

   // atribuition operators
   var variable string = "Hello"
   fmt.Println("Variable:", variable)
   variable2 := "World"
   fmt.Println("Variable2:", variable2)

   // relational operators
   fmt.Println("1 > 2:", 1 > 2)
   fmt.Println("1 < 2:", 1 < 2)
   fmt.Println("1 >= 2:", 1 >= 2)
   fmt.Println("1 <= 2:", 1 <= 2)
   fmt.Println("1 == 2:", 1 == 2)
   fmt.Println("1 != 2:", 1 != 2)

   // logical operators
   fmt.Println("true && false:", true && false)
   fmt.Println("true || false:", true || false)
   fmt.Println("!true:", !true)

   // bitwise operators
   fmt.Println("1 & 2:", 1 & 2)
   fmt.Println("1 | 2:", 1 | 2)
   fmt.Println("1 ^ 2:", 1 ^ 2)
   fmt.Println("1 << 2:", 1 << 2)
   fmt.Println("1 >> 2:", 1 >> 2)

   // assignment operators
   var a int = 10
   a += 10
   fmt.Println("a += 10:", a)
   a -= 10
   fmt.Println("a -= 10:", a)
   a *= 10
   fmt.Println("a *= 10:", a)
   a /= 10
   fmt.Println("a /= 10:", a)
   a %= 10
   fmt.Println("a %= 10:", a)
   a &= 10
   fmt.Println("a &= 10:", a)
   a |= 10
   fmt.Println("a |= 10:", a)
   a ^= 10
   fmt.Println("a ^= 10:", a)
   a <<= 10
   fmt.Println("a <<= 10:", a)
   a >>= 10
   fmt.Println("a >>= 10:", a)

   // unary operators
   var x int = 10
   fmt.Println("x:", x)
   x = -x
   fmt.Println("-x:", x)
   x = +x
   fmt.Println("+x:", x)
   x = ^x
   fmt.Println("^x:", x)

   d := 10
   fmt.Println("d:", d)
   d++
   fmt.Println("d++:", d)
   d += 10
   fmt.Println("d += 10:", d)

   d--
   fmt.Println("d--:", d)
   d -= 10
   fmt.Println("d -= 10:", d)

   d *= 10
   fmt.Println("d *= 10:", d)
   d /= 10
   fmt.Println("d /= 10:", d)
   d %= 10
   fmt.Println("d %= 10:", d)

   // ternary operator
   // text := d > 10 ? "Yes" : "No"
   var text string
   if d < 10 {
	  text = "Yes"
   } else {
	  text = "No"
   }

   fmt.Println("text:", text)
   

}