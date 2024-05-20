# 里氏替换原则(Liskov Substitution Principle)

## 原理

当你扩展一个类时，你应该要能在不修改客户端代码的情况下将子类的对象作为父类对象进行传递。

子类方法的参数类型必须与超类的参数类型相匹配或者更加抽象。

子类方法的返回值类型必须与超类方法的返回值类型或是其子类别相匹配。

子类中的方法不应该抛出比基础方法预期之外的异常类型。

子类不应该加强其前置条件。

子类不能削弱其后置条件。

超类的不变量必须保留。

子类不能修改超类中私有成员变量的值。


## 例子

[Go](https://github.com/XdpCs/go-design-pattern/blob/master/SOLID/03.Liskov-Substitution-Principle/go/main.go)

假设有两种类型的文件，一种是只读文件，一种是可写文件。

### 问题1

```go
type Document struct {
	data     []byte
	fileName string
}

type IDocument interface {
	open() ([]byte, error)
	save([]byte) error
}

type ReadOnlyDocument struct {
	Document
}

func (r *ReadOnlyDocument) open() ([]byte, error) {
	return r.data, nil
}

func (r *ReadOnlyDocument) save(data []byte) error {
	return errors.New("cannot save read-only document")
}
```

这里我们定义了一个只读文件，但是我们的`save`方法是抛出了一个错误，这违反了里氏替换原则，且对于只读文件来说，`save`方法是没有意义的。

### 解决方案

```go
type ModifyIDocument interface {
	open() ([]byte, error)
}

type ModifyIWritableDocument interface {
	ModifyIDocument
	save([]byte) error
}
```

我们将`IDocument`接口拆分成两个接口，一个是只读文件的接口，然后基于只读文件的接口，我们定义了一个可写文件的接口。



