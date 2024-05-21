# 接口隔离原则(Interface Segregation Principle)

## 原理

客户端不应该被迫依赖于它们不使用的方法。接口应该是小而完整的，而不是大而全的。

## 例子

[Go](https://github.com/XdpCs/go-design-pattern/blob/master/SOLID/04.Interface-Segregation-Principle/go/main.go)

假设有一个云计算提供商整合服务，开始仅支持了阿里云，后来又支持了亚马逊云。

### 问题1

```go
type CloudProvider interface {
	storeFile(name string, date []byte) error
	getFile(name string) error
	getCDNAddress() string
}

type AlibabaCloud struct{}

func (a *AlibabaCloud) storeFile(name string, data []byte) error {
	// Store file to Alibaba Cloud
	fmt.Println("Store file to Alibaba Cloud: ", name)
	return nil
}

func (a *AlibabaCloud) getFile(name string) error {
	// Get file from Alibaba Cloud
	fmt.Println("Get file from Alibaba Cloud: ", name, " data: ", string("data"))
	return nil
}

func (a *AlibabaCloud) getCDNAddress() string {
	// Get CDN address from Alibaba Cloud
	return "https://cdn.alibaba.com"
}

type AmazonCloud struct{}

func (a *AmazonCloud) storeFile(name string, data []byte) error {
	// Store file to Amazon Cloud
	fmt.Println("Store file to Amazon Cloud: ", name, " data: ", string(data))
	return nil
}

func (a *AmazonCloud) getFile(name string) error {
	// Get file from Amazon Cloud
	fmt.Println("Get file from Amazon Cloud: ", name)
	return nil
}

func (a *AmazonCloud) getCDNAddress() string {
	// No CDN address for Amazon Cloud
	return ""
}
```

亚马逊云不支持CDN地址，但是我们的接口中定义了`getCDNAddress`方法，这违反了接口隔离原则。

### 解决方案

```go
type CloudStorageProvider interface {
	storeFile(name string, date []byte) error
	getFile(name string) error
}

type CDNProvider interface {
	getCDNAddress() string
}

type ModifyAlibabaCloud struct{}

func (m *ModifyAlibabaCloud) storeFile(name string, data []byte) error {
	// Store file to Alibaba Cloud
	fmt.Println("Store file to Alibaba Cloud: ", name, " data: ", string(data))
	return nil
}

func (m *ModifyAlibabaCloud) getFile(name string) error {
	// Get file from Alibaba Cloud
	fmt.Println("Get file from Alibaba Cloud: ", name)
	return nil
}

func (m *ModifyAlibabaCloud) getCDNAddress() string {
	// Get CDN address from Alibaba Cloud
	return "https://cdn.alibaba.com"
}

type ModifyAmazonCloud struct{}

func (m *ModifyAmazonCloud) storeFile(name string, data []byte) error {
	// Store file to Amazon Cloud
	fmt.Println("Store file to Amazon Cloud: ", name, " data: ", string(data))
	return nil
}

func (m *ModifyAmazonCloud) getFile(name string) error {
	// Get file from Amazon Cloud
	fmt.Println("Get file from Amazon Cloud: ", name)
	return nil
}
```

我们将接口拆分为两个接口，`CloudStorageProvider`和`CDNProvider`，这样亚马逊云就不需要实现`getCDNAddress`方法了。
创建的接口越多，代码就会变得越复杂，所以我们需要根据实际情况来决定是否拆分接口。
