package main

import (
	"fmt"
	"os"
	"time"
)

//Dealing directly with the os.Args
// func main() {
// 	if len(os.Args) < 2 {
// 		fmt.Printf("usage: %s <your_name>\n", os.Args[0])
// 		return
// 	}

// 	name := os.Args[1]
// 	fmt.Printf("Hello %s!\n", name)

// 	fmt.Println("\nAll extra arguments provided:")
// 	for i, v := range os.Args[1:] {
// 		fmt.Printf("number %d is %s\n", i, v)
// 	}
// }

// Now Converting values of os.Args(strings) to intergers
// func main() {
// 	if len(os.Args) < 2 {
// 		fmt.Println("Error: Please provide an age.")
// 		return
// 	}

// 	ageString := os.Args[1]
// 	fmt.Printf("I am currently %s\n", ageString)

// 	age, err := strconv.Atoi(ageString)
// 	if err != nil {
// 		fmt.Println("Please enter a valid whole number!")
// 		return
// 	}

// 	fmt.Printf("In 5 years time, I will be %d and I'll be a super amazing programmer by then\n", age+5)
// }

//Now let's see if we can use everything we've learnt to solve the exercise

func parseDOB(input string) (time.Time, error) {

	layout := "2006-01-02"
	dob, err := time.Parse(layout, input)
	if err != nil {
		return time.Time{}, fmt.Errorf("Invalid input. Enter a valid date of birth: YYYY-MM-DD")
	}

	now := time.Now()

	if dob.After(now) {
		return time.Time{}, fmt.Errorf("Error: date of birth cannot be in the future")
	}

	return dob, nil
}

func calculateAge(dob time.Time) (years, months, days int) {
	now := time.Now()
	years = now.Year() - dob.Year()
	month := now.Month()
	months = int(month) - int(dob.Month())
	year := now.Year()

	previousMonthDays := time.Date(year, month, 0, 0, 0, 0, 0, time.UTC)

	days = now.Day() - dob.Day()

	//if birthday hasn't occurred yet this year, subtract one from years and add 12 to months

	if days < 0 {
		months--
		days += previousMonthDays.Day()
	}

	if months < 0 {
		years--
		months += 12
	}

	return years, months, days
}

func main() {

	if len(os.Args) != 2 {
		fmt.Println("Invalid input. Enter a valid date of birth: YYYY-MM-DD")
		return
	}

	dob, err := parseDOB(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	ageYear, ageMonth, ageDay := calculateAge(dob)

	fmt.Printf("You are %d years, %d months and %d days old\n", ageYear, ageMonth, ageDay)
}
