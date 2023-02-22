快速开始[(中文)](./README-CN.md)

--- 

# Introduction

Welcome to Zenlayer Cloud API Software Developer Kit (SDK). SDK is a supporting tool for Zenlayer Cloud API. It
currently supports Bare Metal Instance, Elastic IP, DDoS Protected IP and other products. More cloud services will be
supported for SDK.

# Requirements

1. You must use Go 1.9.x or later.
2. A Zenlayer Cloud account is created and an Access Key ID and an Access Key Password are created.
   See [Generate an API Access Key](https://docs.console.zenlayer.com/welcome/platform/team-management/generate-an-api-access-key)
   for more details.

# Installation

If you use Maven to manage Java projects, you can add Maven dependencies to the pom.xml file to install Zenlayer Cloud
SDK for Java. In the Maven repository, you can view the Maven dependencies of Zenlayer Cloud services. Add the following
Maven dependency to install the core library of Zenlayer Cloud SDK for Java.

Use go get to install SDK：

```shell
$ go get -u github.com/zenlayer/zenlayercloud-sdk-go
```

# Quick Examples

Take DescribeInstances as an example.

```go
package main

import (
	"fmt"
	bmc "github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/bmc20221120"
	"github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"
	"os"
)

func main() {
	client, _ := bmc.NewClientWithSecretKey(os.Getenv("ZENLAYERCLOUD_SECRET_KEY_ID"), os.Getenv("ZENLAYERCLOUD_SECRET_KEY_PASSWORD"))
	request := bmc.NewDescribeInstancesRequest()
	request.PageSize = 1
	request.PageNum = 100
	response, err := client.DescribeInstances(request)
	if _, ok := err.(*common.ZenlayerCloudSdkError); ok {
		fmt.Printf("An API error has returned: %s", err)
		return
	}
	fmt.Printf("%v\n", response)
}
```

# Quick Examples

Before creating client, if needed, you can specify some other configuration by setting `common.NewConfig`

```go
conf := common.NewConfig()
```

The configurations are as follows:

## Request Timeout

SDK有默认的超时时间，如非必要请不要修改默认设置。 如有需要请在代码中查阅以获取最新的默认值。 Unit：second

```go
    conf.Timeout = 30
```

你可以设置环境变量 `DEBUG=on`开启调试模式，调试模式会打印更详细的日志，当您需要进行详细的排查错误时可以开启。默认调试模式为关闭。 你也可以设置配置 config.Debug = Bool(true)  来进行开启，如下所示：

默认为 `false`

```go
    conf.Debug = Bool(true)
```

## 请求重试

当调用发生网络错误时，可以配置重试发起API的重试，默认重试为关闭的。  
您可以全局开启重试，也可以只对某一个接口请求设置重试

```go
package main

import (
	bmc "github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/bmc20221120"
	"github.com/zenlayer/zenlayercloud-sdk-go/zenlayercloud/common"
	"os"
)

func main() {
	// 全局开启重试，并设置重试次数为3次
	config := common.NewConfig()
	config.AutoRetry = true
	config.MaxRetryTime = 3

	// 配置接口重试
	client, _ := bmc.NewClient(config, os.Getenv("ZENLAYERCLOUD_SECRET_KEY_ID"), os.Getenv("ZENLAYERCLOUD_SECRET_KEY_PASSWORD"))
	request := bmc.NewDescribeInstanceTypesRequest()
	request.SetAutoRetries(true) // 如果不设置true，则依旧使用全局的配置
	request.SetMaxAttempts(2)    // 覆盖全局的配置3次
	response, err = client.DescribeInstanceTypes(request)
}

```

---
快速开始[(中文)](./README-CN.md)
