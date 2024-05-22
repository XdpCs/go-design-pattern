# 依赖倒转原则(Dependency Inversion Principle)

## 原理

高层次的类不应该依赖于低层次的类，两者都应该依赖于抽象接口。抽象接口不应该依赖具体实现，具体实现应该依赖于抽象接口。

低层次的类实现基础操作，高层次的类实现复杂业务逻辑。

## 例子

[Go](https://github.com/XdpCs/go-design-pattern/blob/master/SOLID/05.Dependency-Inversion-Principle/go/main.go)

假如有一个邮件通知系统

### 问题1

```go
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
```

这里`User`类依赖于`EmailNotification`类，如果我们需要更换通知方式，那么我们必须修改`User`类。

### 解决方案

```go
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
```

定义了一个`Notification`接口，`User`类依赖于`Notification`接口，这样当我们需要更换通知方式时，只需要实现`Notification`接口即可。
