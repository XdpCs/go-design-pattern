package main

import "fmt"

// @Title        main.go
// @Description
// @Create       XdpCs 2024-05-13 15:03
// @Update       XdpCs 2024-05-13 15:03

type Employee struct {
	name string
}

func (e *Employee) getName() string {
	return e.name
}

func (e *Employee) printTimeSheetReport() {
	fmt.Println("TimeSheet Report: Name: ", e.name)
}

type ModifyEmployee struct {
	name string
}

func (m *ModifyEmployee) getName() string {
	return m.name
}

type TimeSheetReport struct{}

func (t *TimeSheetReport) print(employee *Employee) {
	fmt.Println("TimeSheet Report: Name: ", employee.getName())
}
