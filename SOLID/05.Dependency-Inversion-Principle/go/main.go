package main

import "fmt"

// @Title        main.go
// @Description
// @Create       XdpCs 2024-05-22 13:14
// @Update       XdpCs 2024-05-22 13:14

// Before Dependency Inversion Principle

type EmailNotification struct{}

func (e *EmailNotification) SendEmail(message string) {
	fmt.Println("Send email: ", message)
}

type User struct {
	emailNotification *EmailNotification
}

func (u *User) Notify(email string) {
	u.emailNotification.SendEmail(email)
}

// After Dependency Inversion Principle

type Notification interface {
	Send(message string)
}

type ModifyEmailNotification struct{}

func (m *ModifyEmailNotification) Send(message string) {
	fmt.Println("Send email: ", message)
}

type ModifyUser struct {
	notification Notification
}

func (m *ModifyUser) Notify(email string) {
	m.notification.Send(email)
}

func main() {
	fmt.Println("Before Dependency Inversion Principle")
	emailNotification := &EmailNotification{}
	user := &User{emailNotification: emailNotification}
	user.Notify("Hello, XdpCs")

	fmt.Println("After Dependency Inversion Principle")
	modifyEmailNotification := &ModifyEmailNotification{}
	modifyUser := &ModifyUser{notification: modifyEmailNotification}
	modifyUser.Notify("Hello, ModifyXdpCs")
}
