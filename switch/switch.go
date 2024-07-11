package main

import "fmt"

func dayOfWeek(day int) string {
	switch day {
	case 1:
		return "Monday"
	case 2:
		return "Tuesday"
	case 3:
		return "Wednesday"
	case 4:
		return "Thursday"
	case 5:
		return "Friday"
	case 6:
		return "Saturday"
	case 7:
		return "Sunday"
	default:
		return "Invalid day"
	}
}

func dayOfWeek1(day int) string {
	var dayOfWeek2 string 

	switch {
	case day == 1:
		dayOfWeek2 = "Monday"
		fallthrough
	case day == 2:
		dayOfWeek2 = "Tuesday"
	case day == 3:
		dayOfWeek2 = "Wednesday"
	case day == 4:
		dayOfWeek2 = "Thursday"
	case day == 5:
		dayOfWeek2 = "Friday"
	case day == 6:
		dayOfWeek2 = "Saturday"
	case day == 7:
		dayOfWeek2 = "Sunday"
	default:
		dayOfWeek2 = "Invalid day"
	}

	return dayOfWeek2
}

func main() {
	fmt.Println("Switch")
	
	day := dayOfWeek(1)
	fmt.Println(day)

	day1 := dayOfWeek1(2)
	fmt.Println(day1)

	dayInvalid := dayOfWeek(8)
	fmt.Println(dayInvalid)

}