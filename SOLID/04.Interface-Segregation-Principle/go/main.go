package main

import "fmt"

// @Title        main.go
// @Description
// @Create       XdpCs 2024-05-21 13:14
// @Update       XdpCs 2024-05-21 13:14

// Before Interface Segregation Principle

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

// After Interface Segregation Principle

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

func main() {
	fmt.Println("Before Interface Segregation Principle")
	alibabaCloud := &AlibabaCloud{}
	amazonCloud := &AmazonCloud{}
	cloudProvider := []CloudProvider{alibabaCloud, amazonCloud}
	for _, provider := range cloudProvider {
		if err := provider.storeFile("test.txt", []byte("test data")); err != nil {
			return
		}
		if err := provider.getFile("test.txt"); err != nil {
			return
		}
		fmt.Println("CDN address: ", provider.getCDNAddress())
	}
	fmt.Println("After Interface Segregation Principle")
	modifyAlibabaCloud := &ModifyAlibabaCloud{}
	modifyAmazonCloud := &ModifyAmazonCloud{}
	cloudStorageProvider := []CloudStorageProvider{modifyAlibabaCloud, modifyAmazonCloud}
	cdnProvider := []CDNProvider{modifyAlibabaCloud}
	for _, provider := range cloudStorageProvider {
		if err := provider.storeFile("test.txt", []byte("test data")); err != nil {
			return
		}
		if err := provider.getFile("test.txt"); err != nil {
			return
		}
	}

	for _, provider := range cdnProvider {
		fmt.Println("CDN address: ", provider.getCDNAddress())
	}
}
