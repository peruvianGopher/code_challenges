package main

import "fmt"

var monthsOrder = []string{
    "January",
    "February",
    "March",
    "April",
    "May",
    "June",
    "July",
    "August",
    "September",
    "October",
    "November",
    "December",
}

var monthsNumber = map[string]int{
    "January": 1,
    "February": 2,
    "March": 3,
    "April": 4,
    "May": 5,
    "June": 6,
    "July": 7,
    "August": 8,
    "September": 9,
    "October": 10,
    "November": 11,
    "December": 12,
}

var monthsDays = map[string]int{
    "January": 31,
    "February": 28,
    "March": 31,
    "April": 30,
    "May": 31,
    "June": 30,
    "July": 31,
    "August": 31,
    "September": 30,
    "October": 31,
    "November": 30,
    "December": 31,
}

var weeks = map[string][]string{
    "January": {},
    "February": {},
    "March": {},
    "April": {},
    "May": {},
    "June": {},
    "July": {},
    "August": {},
    "September": {},
    "October": {},
    "November": {},
    "December": {},
}

var weekDays = []string{
    "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday",
}

var weekDaysNumber = map[string]int{
    "Monday": 1, "Tuesday": 2, "Wednesday": 3, "Thursday": 4,
    "Friday":5 , "Saturday": 6, "Sunday": 7,
}

func main()  {
    fmt.Println(solution(2014, "April", "May", "Wednesday"))
}

func solution(Y int, A string, B string, W string) int {
    if Y % 4 == 0 {
        monthsDays["February"] = 29
    }

    weekDaysCounter := weekDaysNumber[W]
    for _, name := range monthsOrder {
        for i := 1; i <= monthsDays[name]; i++ {
            weeks[name] = append(weeks[name], weekDays[weekDaysCounter-1])
            if weekDaysCounter == 7 {
                weekDaysCounter = 1
            } else {
                weekDaysCounter = weekDaysCounter + 1
            }
        }
    }

    checkCounter := 0
    vacationWeeksCounter := 0
    for i := monthsNumber[A]; i <= monthsNumber[B]; i++ {
        for _, day := range weeks[monthsOrder[i - 1]] {
            if day == "Monday" {
                checkCounter = checkCounter + 1
            }

            if checkCounter > 0 && day == "Sunday" {
                vacationWeeksCounter = vacationWeeksCounter + 1
                checkCounter = 0
            }
        }
    }

    return vacationWeeksCounter
}