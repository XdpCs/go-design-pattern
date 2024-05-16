# 开闭原则(Pen/Closed Principle)

## 原理

类对于扩展是开放的，对于修改是封闭的。

该原则是在实现新功能时候，能保持原有代码的不变。

如果类中存在BUG，应该直接修改类，而不是让子类去对父类负责

## 例子

[Go](https://github.com/XdpCs/go-design-pattern/blob/master/SOLID/02.Pen-Closed-Principle/go/main.go)

假如有一个订单支付系统，有三种支付方式：信用卡、paypal、bitcoin。

### 问题1

```go
func (o *Order) pay() {
	switch o.payType {
	case "creditCard":
		o.payByCreditCard(o.totalCost)
	case "bitcoin":
		o.payByBitcoin(o.totalCost)
	case "paypal":
		o.payByPaypal(o.totalCost)
	}
}
```

如果需要添加新的支付方式，那么你必须修改`Order`类。

### 解决方案

使用策略模式，将支付方式抽象出来。这样当需要添加新的支付方式时，只需要添加新的支付方式实现接口即可。

```go
type ModifyOrder struct {
	totalCost float32
	payType   PaymentStrategy
}

type PaymentStrategy interface {
	Pay(amount float32)
}

func (m *ModifyOrder) pay() {
	m.payType.Pay(m.totalCost)
}
```