package main

import "fmt"

// @Title        main.go
// @Description
// @Create       XdpCs 2024-05-13 15:03
// @Update       XdpCs 2024-05-13 15:03

// Before Single Responsibility Principle

type Employee struct {
	name string
}

func (e *Employee) getName() string {
	return e.name
}

func (e *Employee) printTimeSheetReport() {
	fmt.Println("TimeSheet Report: Name: ", e.name)
}

// After Single Responsibility Principle

type ModifyEmployee struct {
	name string
}

func (m *ModifyEmployee) getName() string {
	return m.name
}

type TimeSheetReport struct{}

func (t *TimeSheetReport) print(employee *ModifyEmployee) {
	fmt.Println("TimeSheet Report: Name: ", employee.getName())
}

func main() {
	fmt.Println("Before Single Responsibility Principle")
	e := Employee{name: "XdpCs"}
	e.printTimeSheetReport()
	fmt.Println("After Single Responsibility Principle")
	employee := ModifyEmployee{name: "ModifyXdpCs"}
	timeSheetReport := TimeSheetReport{}
	timeSheetReport.print(&employee)
}
