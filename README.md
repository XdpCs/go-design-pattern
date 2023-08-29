# Go-Design-Pattern

## use chatgpt3.5 to generate

当涉及到在 Go 语言中使用常见的 23 种设计模式时，以下是每种模式的简要解析以及示例代码。请注意，每种模式都有多种用例场景和实现方式，以下示例代码只是其中一种方式。

1. **单例模式（Singleton Pattern）**：
    - **解析**：确保一个类只有一个实例，并提供一个全局访问点。
    - **示例代码**：使用 `sync.Once` 来实现线程安全的单例模式。
   ```go
   package singleton

   import (
       "sync"
   )

   type Singleton struct {
       // fields here
   }

   var instance *Singleton
   var once sync.Once

   func GetInstance() *Singleton {
       once.Do(func() {
           instance = &Singleton{}
       })
       return instance
   }
   ```

2. **工厂模式（Factory Pattern）**：
    - **解析**：定义一个创建对象的接口，由子类决定实例化的类是哪一个。
    - **示例代码**：使用工厂模式创建不同类型的产品。
   ```go
   package factory

   type Product interface {
       DoSomething()
   }

   type ConcreteProductA struct{}

   func (p *ConcreteProductA) DoSomething() {
       // implementation for Product A
   }

   type ConcreteProductB struct{}

   func (p *ConcreteProductB) DoSomething() {
       // implementation for Product B
   }

   func CreateProduct(productType string) Product {
       switch productType {
       case "A":
           return &ConcreteProductA{}
       case "B":
           return &ConcreteProductB{}
       default:
           return nil
       }
   }
   ```

3. **抽象工厂模式（Abstract Factory Pattern）**：
    - **解析**：提供一个创建一系列相关或相互依赖对象的接口，而无需指定其具体类。
    - **示例代码**：实现一个抽象工厂和多个具体工厂，每个工厂创建一组相关的产品。
   ```go
   package abstractfactory

   type AbstractFactory interface {
       CreateProductA() ProductA
       CreateProductB() ProductB
   }

   type ConcreteFactory1 struct{}

   func (cf *ConcreteFactory1) CreateProductA() ProductA {
       return &ConcreteProductA1{}
   }

   func (cf *ConcreteFactory1) CreateProductB() ProductB {
       return &ConcreteProductB1{}
   }

   type ConcreteFactory2 struct{}

   func (cf *ConcreteFactory2) CreateProductA() ProductA {
       return &ConcreteProductA2{}
   }

   func (cf *ConcreteFactory2) CreateProductB() ProductB {
       return &ConcreteProductB2{}
   }
   ```

4. **建造者模式（Builder Pattern）**：
    - **解析**：将一个复杂对象的构建过程与其表示分离，使得同样的构建过程可以创建不同的表示。
    - **示例代码**：创建一个产品结构体和一个构造器，构造器用于逐步构建产品。
   ```go
   package builder

   type Product struct {
       Part1 string
       Part2 string
   }

   type Builder interface {
       BuildPart1()
       BuildPart2()
       GetResult() Product
   }

   type ConcreteBuilder struct {
       product Product
   }

   func (cb *ConcreteBuilder) BuildPart1() {
       cb.product.Part1 = "Part 1"
   }

   func (cb *ConcreteBuilder) BuildPart2() {
       cb.product.Part2 = "Part 2"
   }

   func (cb *ConcreteBuilder) GetResult() Product {
       return cb.product
   }

   type Director struct {
       builder Builder
   }

   func (d *Director) Construct() Product {
       d.builder.BuildPart1()
       d.builder.BuildPart2()
       return d.builder.GetResult()
   }
   ```

5. **原型模式（Prototype Pattern）**：
    - **解析**：通过复制现有对象来创建新对象，而不是从头开始创建。
    - **示例代码**：实现一个原型接口，克隆方法根据现有对象创建新对象。
   ```go
   package prototype

   type Prototype interface {
       Clone() Prototype
   }

   type ConcretePrototype struct {
       Field string
   }

   func (cp *ConcretePrototype) Clone() Prototype {
       return &ConcretePrototype{
           Field: cp.Field,
       }
   }
   ```

6. **适配器模式（Adapter Pattern）**：
    - **解析**：将一个类的接口转换成客户希望的另一个接口。
    - **示例代码**：通过实现目标接口并在内部包含适配的对象来实现适配器模式。
   ```go
   package adapter

   type Target interface {
       Request() string
   }

   type Adaptee struct{}

   func (a *Adaptee) SpecificRequest() string {
       return "Adaptee's request"
   }

   type Adapter struct {
       adaptee *Adaptee
   }

   func (ad *Adapter) Request() string {
       return ad.adaptee.SpecificRequest()
   }
   ```

7. **桥接模式（Bridge Pattern）**：
    - **解析**：将抽象部分与其实现部分分离，使它们可以独立地变化。
    - **示例代码**：定义抽象部分和实现部分的接口，并实现它们的具体类。
   ```go
   package bridge

   type Implementor interface {
       OperationImpl() string
   }

   type ConcreteImplementorA struct{}

   func (ia *ConcreteImplementorA) OperationImpl() string {
       return "ConcreteImplementorA"
   }

   type ConcreteImplementorB struct{}

   func (ib *ConcreteImplementorB) OperationImpl() string {
       return "ConcreteImplementorB"
   }

   type Abstraction struct {
       implementor Implementor
   }

   func (a *Abstraction) Operation() string {
       return a.implementor.OperationImpl()
   }
   ```

8. **组合模式（Composite Pattern）**：
    - **解析**：将对象组合成树状结构以表示“部分-整体”的层次结构。使得用户对单个对象和组合对象的使用具有一致性。
    - **示例代码**：创建抽象组件和叶子节点，通过组合形成树状结构。

    ```go
    package composite

    type Component interface {
        Operation() string
    }
    
    type Leaf struct{}
    
    func (l *Leaf) Operation() string {
        return "Leaf"
    }
    
    type Composite struct {
        children []Component
    }
    
    func (c *Composite) Operation() string {
        result := "Composite("
        for _, child := range c.children {
            result += child.Operation() + " "
        }
        result += ")"
        return result
    }
   ```

9. **装饰器模式（Decorator Pattern）**：
    - **解析**：动态地给一个对象添加一些额外的职责，而不会影响其类。
    - **示例代码**：通过组合来实现装饰器模式，每个具体装饰器增加特定的功能。
   ```go
   package decorator

   type Component interface {
       Operation() string
   }

   type ConcreteComponent struct{}

   func (c *ConcreteComponent) Operation() string {
       return "ConcreteComponent"
   }

   type Decorator interface {
       Operation() string
   }

   type ConcreteDecoratorA struct {
       component Component
   }

   func (d *ConcreteDecoratorA) Operation() string {
       return "ConcreteDecoratorA(" + d.component.Operation() + ")"
   }

   type ConcreteDecoratorB struct {
       component Component
   }

   func (d *ConcreteDecoratorB) Operation() string {
       return "ConcreteDecoratorB(" + d.component.Operation() + ")"
   }
   ```

10. **外观模式（Facade Pattern）**：
     - **解析**：提供一个统一的接口，用来访问子系统中的一群接口。
     - **示例代码**：创建一个外观，它封装了多个子系统的调用，对外提供简化的接口。
    ```go
    package facade

    type SubsystemA struct{}

       func (sa *SubsystemA) OperationA() string {
           return "SubsystemA OperationA"
       }
    
       type SubsystemB struct{}
    
       func (sb *SubsystemB) OperationB() string {
           return "SubsystemB OperationB"
       }
    
       type Facade struct {
           subsystemA *SubsystemA
           subsystemB *SubsystemB
       }
    
       func NewFacade() *Facade {
           return &Facade{
               subsystemA: &SubsystemA{},
               subsystemB: &SubsystemB{},
           }
       }
    
       func (f *Facade) Operation() string {
           result := f.subsystemA.OperationA() + "\n"
           result += f.subsystemB.OperationB()
           return result
       }
       ```

11. **享元模式（Flyweight Pattern）**：
    - **解析**：共享对象以减小内存占用，特别适用于大量相似对象的场景。
    - **示例代码**：创建享元工厂，通过缓存共享的对象来节省资源。
   ```go
   package flyweight

   type Flyweight interface {
       Operation() string
   }

   type ConcreteFlyweight struct {
       sharedState string
   }

   func (cf *ConcreteFlyweight) Operation() string {
       return cf.sharedState
   }

   type FlyweightFactory struct {
       flyweights map[string]Flyweight
   }

   func NewFlyweightFactory() *FlyweightFactory {
       return &FlyweightFactory{
           flyweights: make(map[string]Flyweight),
       }
   }

   func (ff *FlyweightFactory) GetFlyweight(key string) Flyweight {
       if flyweight, exists := ff.flyweights[key]; exists {
           return flyweight
       }
       newFlyweight := &ConcreteFlyweight{sharedState: key}
       ff.flyweights[key] = newFlyweight
       return newFlyweight
   }
   ```

12. **代理模式（Proxy Pattern）**：
    - **解析**：为其他对象提供一种代理以控制对这个对象的访问。
    - **示例代码**：创建一个代理，实现与实际对象相同的接口，但可以控制对实际对象的访问。
   ```go
   package proxy

   type Subject interface {
       Request() string
   }

   type RealSubject struct{}

   func (rs *RealSubject) Request() string {
       return "RealSubject request"
   }

   type Proxy struct {
       realSubject *RealSubject
   }

   func (p *Proxy) Request() string {
       if p.realSubject == nil {
           p.realSubject = &RealSubject{}
       }
       result := "Proxy: preparing request\n"
       result += p.realSubject.Request()
       return result
   }
   ```

13. **策略模式（Strategy Pattern）**：
    - **解析**：定义一系列算法，将它们封装起来，并且可以相互替换。
    - **示例代码**：使用接口实现策略模式，允许在运行时切换算法。
   ```go
   package strategy

   type Strategy interface {
       Execute()
   }

   type ConcreteStrategyA struct{}

   func (s *ConcreteStrategyA) Execute() {
       // implementation for strategy A
   }

   type ConcreteStrategyB struct{}

   func (s *ConcreteStrategyB) Execute() {
       // implementation for strategy B
   }

   type Context struct {
       strategy Strategy
   }

   func (c *Context) SetStrategy(strategy Strategy) {
       c.strategy = strategy
   }

   func (c *Context) ExecuteStrategy() {
       c.strategy.Execute()
   }
   ```

14. **模板方法模式（Template Method Pattern）**：
    - **解析**：定义一个算法的骨架，将一些步骤延迟到子类中实现。
    - **示例代码**：使用函数参数或接口来实现模板方法模式。
   ```go
   package templatemethod

   type AbstractClass interface {
       TemplateMethod()
       PrimitiveOperation1()
       PrimitiveOperation2()
   }

   type ConcreteClass struct{}

   func (c *ConcreteClass) TemplateMethod() {
       c.PrimitiveOperation1()
       c.PrimitiveOperation2()
   }

   func (c *ConcreteClass) PrimitiveOperation1() {
       // implementation for step 1
   }

   func (c *ConcreteClass) PrimitiveOperation2() {
       // implementation for step 2
   }
   ```

15. **状态模式（State Pattern）**：

    - **解析**：允许对象在其内部状态改变时改变其行为。
    - **示例代码**：创建状态接口和多个具体状态，将状态切换交给上下文来管理。
    ```go
    package state
    
        type State interface {
            Handle() string
        }
    
        type ConcreteStateA struct{}
    
        func (sa *ConcreteStateA) Handle() string {
            return "ConcreteStateA"
        }
    
        type ConcreteStateB struct{}
    
        func (sb *ConcreteStateB) Handle() string {
            return "ConcreteStateB"
        }
    
        type Context struct {
            state State
        }
    
        func (c *Context) SetState(state State) {
            c.state = state
        }
    
        func (c *Context) Request() string {
            return c.state.Handle()
        }
    ```

16. **观察者模式（Observer Pattern）**：
    - **解析**：定义一种依赖关系，当一个对象的状态发生改变时，所有依赖于它的对象都会得到通知并自动更新。
    - **示例代码**：使用通道实现观察者模式，让订阅者可以接收到发布者的更新。
    ```go
    package observer

    type Observer interface {
        Update(message string)
    }

    type Publisher struct {
        observers []Observer
    }

    func (p *Publisher) Attach(observer Observer) {
        p.observers = append(p.observers, observer)
    }

    func (p *Publisher) Notify(message string) {
        for _, observer := range p.observers {
            observer.Update(message)
        }
    }

    type ConcreteObserver struct {
        id int
    }

    func (co *ConcreteObserver) Update(message string) {
        // handle update
    }
    ```

17. **备忘录模式（Memento Pattern）**：
    - **解析**：在不违反封装的情况下，捕获一个对象的内部状态，并在该对象之外保存这个状态。
    - **示例代码**：创建备忘录类来存储对象的状态，以便后续恢复。
    ```go
    package memento

    type Memento struct {
        state string
    }

    type Originator struct {
        state string
    }

    func (o *Originator) SetState(state string) {
        o.state = state
    }

    func (o *Originator) SaveToMemento() Memento {
        return Memento{state: o.state}
    }

    func (o *Originator) RestoreFromMemento(m Memento) {
        o.state = m.state
    }

    type Caretaker struct {
        mementos []Memento
    }

    func (c *Caretaker) AddMemento(m Memento) {
        c.mementos = append(c.mementos, m)
    }

    func (c *Caretaker) GetMemento(index int) Memento {
        return c.mementos[index]
    }
    ```

18. **解释器模式（Interpreter Pattern）**：
    - **解析**：提供了评估语言的语法或表达式的方式。
    - **示例代码**：创建抽象表达式和终结符表达式，定义文法并解释表达式。
    ```go
    package interpreter

    type Expression interface {
        Interpret() int
    }

    type TerminalExpression struct {
        data int
    }

    func (te *TerminalExpression) Interpret() int {
        return te.data
    }

    type NonTerminalExpression struct {
        left  Expression
        right Expression
    }

    func (nte *NonTerminalExpression) Interpret() int {
        // interpret and combine expressions
    }
    ```

19. **命令模式（Command Pattern）**：
    - **解析**：将请求封装成对象，以便在不同的请求和参数之间进行参数化。
    - **示例代码**：创建命令接口和具体命令，将请求封装为命令对象。
    ```go
    package command

    type Command interface {
        Execute()
    }

    type Receiver struct{}

    func (r *Receiver) Action() {
        // action implementation
    }

    type ConcreteCommand struct {
        receiver *Receiver
    }

    func (cc *ConcreteCommand) Execute() {
        cc.receiver.Action()
    }

    type Invoker struct {
        command Command
    }

    func (i *Invoker) SetCommand(command Command) {
        i.command = command
    }

    func (i *Invoker) ExecuteCommand() {
        i.command.Execute()
    }
    ```

20. **迭代器模式（Iterator Pattern）**：
    - **解析**：提供一种方法顺序访问一个聚合对象中的各个元素，而又不需暴露其内部表示。
    - **示例代码**：定义迭代器接口和具体迭代器，实现在聚合对象中遍历元素。
    ```go
    package iterator

    type Iterator interface {
        Next() interface{}
        HasNext() bool
    }

    type ConcreteIterator struct {
        collection []interface{}
        index      int
    }

    func (ci *ConcreteIterator) Next() interface{} {
        item := ci.collection[ci.index]
        ci.index++
        return item
    }

    func (ci *ConcreteIterator) HasNext() bool {
        return ci.index < len(ci.collection)
    }

    type Aggregate interface {
        CreateIterator() Iterator
    }

    type ConcreteAggregate struct {
        collection []interface{}
    }

    func (ca *ConcreteAggregate) CreateIterator() Iterator {
        return &ConcreteIterator{
            collection: ca.collection,
            index:      0,
        }
    }
    ```

21. **中介者模式（Mediator Pattern）**：
    - **解析**：用一个中介对象来封装一系列对象之间的交互。
    - **示例代码**：创建中介者接口和具体中介者，协调多个同事类之间的交互。
    ```go
    package mediator

    type Mediator interface {
        Send(message string, colleague Colleague)
    }

    type ConcreteMediator struct {
        colleague1 Colleague
        colleague2 Colleague
    }

    func (cm *ConcreteMediator) Send(message string, colleague Colleague) {
        if colleague == cm.colleague1 {
            cm.colleague2.Receive(message)
        } else {
            cm.colleague1.Receive(message)
        }
    }

    type Colleague interface {
        Send(message string)
        Receive(message string)
    }

    type ConcreteColleague1 struct {
        mediator Mediator


    }

    func (cc1 *ConcreteColleague1) Send(message string) {
        cc1.mediator.Send(message, cc1)
    }

    func (cc1 *ConcreteColleague1) Receive(message string) {
        // handle received message
    }

    type ConcreteColleague2 struct {
        mediator Mediator
    }

    func (cc2 *ConcreteColleague2) Send(message string) {
        cc2.mediator.Send(message, cc2)
    }

    func (cc2 *ConcreteColleague2) Receive(message string) {
        // handle received message
    }
    ```

22. **访问者模式（Visitor Pattern）**：
    - **解析**：表示一个作用于某对象结构中的各元素的操作。
    - **示例代码**：定义访问者和元素接口，实现访问者来执行不同操作。
    ```go
    package visitor

    type Visitor interface {
        VisitElementA(element ElementA)
        VisitElementB(element ElementB)
    }

    type ConcreteVisitor struct{}

    func (cv *ConcreteVisitor) VisitElementA(element ElementA) {
        // perform operation on ElementA
    }

    func (cv *ConcreteVisitor) VisitElementB(element ElementB) {
        // perform operation on ElementB
    }

    type Element interface {
        Accept(visitor Visitor)
    }

    type ElementA struct{}

    func (ea *ElementA) Accept(visitor Visitor) {
        visitor.VisitElementA(*ea)
    }

    type ElementB struct{}

    func (eb *ElementB) Accept(visitor Visitor) {
        visitor.VisitElementB(*eb)
    }
    ```

23. **责任链模式（Chain of Responsibility Pattern）**：
    - **解析**：为解除请求的发送者和接收者之间耦合，而使多个对象都有机会处理请求。
    - **示例代码**：创建处理器接口和多个具体处理器，通过链表连接处理器形成责任链。
    ```go
    package chainofresponsibility

    type Handler interface {
        SetSuccessor(handler Handler)
        HandleRequest(request int) string
    }

    type ConcreteHandlerA struct {
        successor Handler
    }

    func (cha *ConcreteHandlerA) SetSuccessor(handler Handler) {
        cha.successor = handler
    }

    func (cha *ConcreteHandlerA) HandleRequest(request int) string {
        if request < 10 {
            return "Handled by ConcreteHandlerA"
        } else if cha.successor != nil {
            return cha.successor.HandleRequest(request)
        }
        return "Request cannot be handled"
    }

    type ConcreteHandlerB struct {
        successor Handler
    }

    func (chb *ConcreteHandlerB) SetSuccessor(handler Handler) {
        chb.successor = handler
    }

    func (chb *ConcreteHandlerB) HandleRequest(request int) string {
        if request >= 10 && request < 20 {
            return "Handled by ConcreteHandlerB"
        } else if chb.successor != nil {
            return chb.successor.HandleRequest(request)
        }
        return "Request cannot be handled"
    }
    ```

这些示例代码提供了 Go 语言中各种设计模式的基本实现，以及如何在不同场景中应用它们。请根据具体的项目需求和代码架构选择适合的设计模式，以提高代码的可维护性和可扩展性。
