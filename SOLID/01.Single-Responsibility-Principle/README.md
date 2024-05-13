# 单一职责原则(Single Responsibility Principle)

尽量让每个类只负责软件中的一个功能，并将该功能完全封装在该类中。

如果类负责的东西太多，那么当其中任何一件事发生改变，你对必须对类进行修改。

这样的类是脆弱的，而且很难被重用。

## 例子

假设有一个`Employee`类，它负责管理员工数据，它有两个函数`getName()`和`printTimeSheetReport()`。

### 问题1

该类的工作是负责员工数据，类中包含了多个不同的行为，这违反了单一职责原则。

### 问题2

如果`printTimeSheetReport()`函数需要改变，那么你必须修改`Employee`类。

### 解决方案

将`printTimeSheetReport()`函数移动到一个新的类`TimeSheetReport`中。