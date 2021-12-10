package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Every seventh day in a month has the same name as the previous
func FindDay(day int) int {
	dList := map[int][]int{
		0: []int{0, 7, 14, 21, 28},
		1: []int{1, 8, 15, 22, 29},
		2: []int{2, 9, 16, 23, 30},
		3: []int{3, 10, 17, 24, 31},
		4: []int{4, 11, 18, 25},
		5: []int{5, 12, 19, 26},
		6: []int{6, 13, 20, 27},
	}
	rD := 0

	for key, element := range dList {
		for _, value := range element {
			if day == value {
				rD = key
			}
		}
	}

	return rD
}

// "Corresponding months" are those months within the calendar year that start on the same day of the week
func FindMonth(month string) int {
	mList := map[int][]string{
		0: []string{"Jan", "Oct"},
		1: []string{"May"},
		2: []string{"Aug"},
		3: []string{"Feb", "Mar", "Nov"},
		4: []string{"Jun"},
		5: []string{"Sept", "Dec"},
		6: []string{"Apr", "July"},
	}
	rM := 0

	for key, element := range mList {
		for _, value := range element {
			if month == value {
				rM = key
			}
		}
	}

	return rM
}

// Year of the century mod 28
func FindYear(year int) (int, int) {
	yList := map[int][]int{
		0: []int{0, 6, 12, 17, 23},
		1: []int{1, 7, 12, 18, 24},
		2: []int{2, 8, 13, 19, 24},
		3: []int{3, 8, 14, 20, 25},
		4: []int{4, 9, 15, 20, 26},
		5: []int{4, 10, 16, 21, 27},
		6: []int{5, 11, 16, 22, 0},
	}

	// Year leap
	yLeapList := map[int][]int{
		0: []int{0, 12},
		1: []int{12, 24},
		2: []int{8, 24},
		3: []int{8, 20},
		4: []int{4, 20},
		5: []int{4, 16},
		6: []int{0, 16},
	}
	rY := 0
	isLeapYear := 0

	for key, element := range yList {
		for _, value := range element {
			if year%28 == value {
				rY = key
			}
		}
	}
	// Find leave year from y_leap_list
	for _, element := range yLeapList {
		for _, value := range element {
			if year%28 == value {
				isLeapYear = 1
			}
		}
	}

	return rY, isLeapYear
}

/*
Determination of the day of the week
The determination of the day of the week for any date may be performed with a variety of algorithms.
In addition, perpetual calendars require no calculation by the user, and are essentially lookup tables.
A typical application is to calculate the day of the week on which someone was born or a specific event occurred.
refer: https://en.wikipedia.org/wiki/Determination_of_the_day_of_the_week
*/
func main() {
	/*
		day := FindDay(18)
		fmt.Println(day)
		month := FindMonth("Dec")
		fmt.Println(month)
		year, leapYear := FindYear(2021)
		fmt.Println(year, leapYear)
	*/

	// Weekdays list
	weekdays := []string{"Saturday", "Sunday", "Monday", "Tuesday", "Wednesday", "Thusday", "Friday"}
	fmt.Println("Eenter date (day, month, year) example: 7-Dec-2021: ")
	var dateInput string
	fmt.Scanln(&dateInput)

	// Split date from format d, m, yyyy to each variables
	dateInputSplit := strings.Split(dateInput, "-")
	// fmt.Println(dateInputSplit)
	day, errDay := strconv.Atoi(strings.Trim(dateInputSplit[0], ""))
	if errDay != nil {
		fmt.Println(errDay)
	}
	month := strings.Trim(dateInputSplit[1], "")
	y, errYear := strconv.Atoi(strings.Trim(dateInputSplit[2], ""))
	if errYear != nil {
		fmt.Println(errYear)
	}
	// fmt.Println(err)

	year, leapYear := FindYear(y)
	weekNum := (FindDay(day) + FindMonth(month) + year) % 7

	fmt.Println("i. ", weekdays[weekNum])

	if leapYear == 1 {
		fmt.Println("ii. Leap year")
	}
}
